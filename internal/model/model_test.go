package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationDefaults(t *testing.T) {
	config := NewConfiguration(map[string]string{})

	assert.Equal(t, "u16", config.ListLenPrefixLenType)
	assert.Equal(t, "u16", config.StringLenPrefixLenType)
	assert.False(t, config.LittleEndian)
	assert.Empty(t, config.JavaPackage)
	assert.Empty(t, config.GoPackage)
	assert.Empty(t, config.GoModule)
	assert.NotNil(t, config.Padding)
	assert.True(t, config.Padding.IsDefault())
}

func TestNewConfigurationOverrides(t *testing.T) {
	config := NewConfiguration(map[string]string{
		ArrayPrefixLenType:     "u32",
		StringPrefixLenType:    "u8",
		LittleEndian:           "TRUE",
		JavaPackage:            "com.example.messages",
		GoPackage:              "messages",
		GoModule:               "github.com/example/messages",
		FixedStringPadFromLeft: "true",
		FixedStringPadChar:     "'0'",
	})

	assert.Equal(t, "u32", config.ListLenPrefixLenType)
	assert.Equal(t, "u8", config.StringLenPrefixLenType)
	assert.True(t, config.LittleEndian)
	assert.Equal(t, "com.example.messages", config.JavaPackage)
	assert.Equal(t, "messages", config.GoPackage)
	assert.Equal(t, "github.com/example/messages", config.GoModule)
	assert.Equal(t, &Padding{PadChar: "'0'", PadLeft: true}, config.Padding)
}

func TestAddOption(t *testing.T) {
	tests := []struct {
		name        string
		optionName  string
		optionValue string
		expectSaved bool
		expectError string
	}{
		{
			name:        "valid option",
			optionName:  StringPrefixLenType,
			optionValue: "u32",
			expectSaved: true,
		},
		{
			name:        "free-form option",
			optionName:  JavaPackage,
			optionValue: "com.example.anything",
			expectSaved: true,
		},
		{
			name:        "invalid option value",
			optionName:  StringPrefixLenType,
			optionValue: "i16",
			// current behavior: the error is recorded but the option is
			// still stored (only unknown names and duplicates are rejected)
			expectSaved: true,
			expectError: "Option " + StringPrefixLenType + " is not allowed to be i16",
		},
		{
			name:        "unknown option name",
			optionName:  "NoSuchOption",
			optionValue: "1",
			expectError: "Option NoSuchOption is not allowed in this context",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewBinaryModel()
			m.AddOption(tt.optionName, tt.optionValue, 1, 1)

			if tt.expectError != "" {
				assert.Len(t, m.SyntaxErrors, 1)
				assert.Contains(t, m.SyntaxErrors[0].Msg, tt.expectError)
				assert.Equal(t, 1, m.SyntaxErrors[0].Line)
			} else {
				assert.Empty(t, m.SyntaxErrors)
			}
			if tt.expectSaved {
				assert.Equal(t, tt.optionValue, m.Options[tt.optionName])
			} else {
				assert.NotContains(t, m.Options, tt.optionName)
			}
		})
	}
}

func TestAddOptionDuplicate(t *testing.T) {
	m := NewBinaryModel()
	m.AddOption(StringPrefixLenType, "u16", 1, 1)
	m.AddOption(StringPrefixLenType, "u32", 2, 1)

	assert.Len(t, m.SyntaxErrors, 1)
	assert.Contains(t, m.SyntaxErrors[0].Msg, "already defined")
	assert.Equal(t, 2, m.SyntaxErrors[0].Line)
	assert.Equal(t, "u16", m.Options[StringPrefixLenType])
}

func TestAddPacket(t *testing.T) {
	m := NewBinaryModel()

	root := &Packet{Name: "Root", IsRoot: true}
	child := &Packet{Name: "Child"}
	m.AddPacket(root)
	m.AddPacket(child)

	assert.Equal(t, root, m.PacketsMap["Root"])
	assert.Equal(t, child, m.PacketsMap["Child"])
	assert.Equal(t, []*Packet{root, child}, m.Packets)
	assert.Same(t, root, m.RootPacket)
	assert.Empty(t, m.SyntaxErrors)
}

func TestAddPacketDuplicate(t *testing.T) {
	m := NewBinaryModel()
	m.AddPacket(&Packet{Name: "Root", IsRoot: true})
	m.AddPacket(&Packet{Name: "Root"})

	assert.Len(t, m.SyntaxErrors, 1)
	assert.Contains(t, m.SyntaxErrors[0].Msg, "Duplicate packet definition for Root")
	assert.Len(t, m.Packets, 1)
}

func TestAddPacketMultipleRoots(t *testing.T) {
	m := NewBinaryModel()
	first := &Packet{Name: "First", IsRoot: true}
	second := &Packet{Name: "Second", IsRoot: true}
	m.AddPacket(first)
	m.AddPacket(second)

	assert.Len(t, m.SyntaxErrors, 1)
	assert.Contains(t, m.SyntaxErrors[0].Msg, "Multiple root packets are not allowed")
	assert.Same(t, first, m.RootPacket)
	assert.Len(t, m.Packets, 2)
}

func TestAddMetaData(t *testing.T) {
	m := NewBinaryModel()
	m.AddMetaData(MetaData{Name: "Price", Attr: &BasicFieldAttribute{Type: "u64"}})
	m.AddMetaData(MetaData{Name: "Price", Attr: &BasicFieldAttribute{Type: "u32"}})

	assert.Len(t, m.MetaDataMap, 1)
	assert.Len(t, m.SyntaxErrors, 1)
	assert.Contains(t, m.SyntaxErrors[0].Msg, "Duplicate metadata definition for Price")
}

func TestResolveDependencies(t *testing.T) {
	t.Run("resolves object field reference", func(t *testing.T) {
		child := &Packet{Name: "Child"}
		field := &Field{Name: "Child", Attr: &ObjectFieldAttribute{PacketName: "Child"}}
		parent := &Packet{Name: "Parent", Fields: []*Field{field}}

		m := NewBinaryModel()
		m.AddPacket(child)
		m.AddPacket(parent)
		m.ResolveDependencies()

		assert.Same(t, child, field.Attr.(*ObjectFieldAttribute).RefPacket)
		assert.Empty(t, m.SyntaxErrors)
		assert.NotNil(t, m.Config)
	})

	t.Run("unknown packet reference reports syntax error", func(t *testing.T) {
		field := &Field{Name: "Missing", Line: 3, Column: 1,
			Attr: &ObjectFieldAttribute{PacketName: "Missing"}}
		parent := &Packet{Name: "Parent", Fields: []*Field{field}}

		m := NewBinaryModel()
		m.AddPacket(parent)
		m.ResolveDependencies()

		assert.Nil(t, field.Attr.(*ObjectFieldAttribute).RefPacket)
		assert.Len(t, m.SyntaxErrors, 1)
		assert.Contains(t, m.SyntaxErrors[0].Msg, "Unknown packet type Missing for field Missing")
		assert.Equal(t, 3, m.SyntaxErrors[0].Line)
	})
}

func TestFieldGetType(t *testing.T) {
	child := &Packet{Name: "Child"}
	tests := []struct {
		name     string
		field    *Field
		expected string
	}{
		{name: "alias uint32 normalizes to u32", field: &Field{Attr: &BasicFieldAttribute{Type: "uint32"}}, expected: "u32"},
		{name: "alias int64 normalizes to i64", field: &Field{Attr: &BasicFieldAttribute{Type: "int64"}}, expected: "i64"},
		{name: "alias float64 normalizes to f64", field: &Field{Attr: &BasicFieldAttribute{Type: "float64"}}, expected: "f64"},
		{name: "short name f32 stays f32", field: &Field{Attr: &BasicFieldAttribute{Type: "f32"}}, expected: "f32"},
		{name: "char stays char", field: &Field{Attr: &BasicFieldAttribute{Type: "char"}}, expected: "char"},
		{name: "fixed string is string", field: &Field{Attr: &FixedStringFieldAttribute{Length: 10}}, expected: "string"},
		{name: "dynamic string is string", field: &Field{Attr: &DynamicStringFieldAttribute{}}, expected: "string"},
		{name: "match field is match", field: &Field{Attr: &MatchFieldAttribute{}}, expected: "match"},
		{name: "object field uses packet name", field: &Field{Attr: &ObjectFieldAttribute{RefPacket: child}}, expected: "Child"},
		{name: "length field uses length type", field: &Field{Attr: &LengthFieldAttribute{LengthType: "u16"}}, expected: "u16"},
		{name: "checksum field uses type", field: &Field{Attr: &CheckSumFieldAttribute{Type: "u32"}}, expected: "u32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.field.GetType())
		})
	}
}

func TestPaddingIsDefault(t *testing.T) {
	assert.True(t, Padding{PadChar: "' '", PadLeft: false}.IsDefault())
	assert.False(t, Padding{PadChar: "'0'", PadLeft: false}.IsDefault())
	assert.False(t, Padding{PadChar: "' '", PadLeft: true}.IsDefault())
	assert.False(t, Padding{PadChar: "'\x00'", PadLeft: false}.IsDefault())
}
