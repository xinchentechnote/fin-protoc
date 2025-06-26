package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

func TestNewRustGenerator(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]model.Packet),
	}

	generator := NewRustGenerator(config, binModel)

	assert.NotNil(t, generator)
	assert.Equal(t, config, generator.config)
	assert.Equal(t, binModel, generator.binModel)
}

func TestGenerateRust(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
		Fields: []model.Field{
			{
				Name: "testField",
				Type: "u32",
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	generator := NewRustGenerator(config, binModel)
	output, err := generator.Generate(binModel)

	assert.NoError(t, err)
	assert.NotEmpty(t, output)
	assert.Contains(t, output, "test_packet.rs")
	assert.Contains(t, output, "lib.rs")
}

func TestGenerateRustFileForPacket(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
		Fields: []model.Field{
			{
				Name: "testField",
				Type: "u32",
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	generator := NewRustGenerator(config, binModel)
	code := generator.generateRustFileForPacket(&packet)

	assert.Contains(t, code, "pub struct TestPacket")
	assert.Contains(t, code, "impl BinaryCodec for TestPacket")
	assert.Contains(t, code, "mod test_packet_tests")
}

func TestGenerateLibCode(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	generator := NewRustGenerator(config, binModel)
	code := generator.generateLibCode()

	assert.Equal(t, "pub mod test_packet;\n", code)
}

func TestGenerateUseCode(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
		Fields: []model.Field{
			{
				Name: "testField",
				Type: "OtherPacket",
			},
			{
				Name: "matchField",
				Type: "match",
				MatchPairs: []model.MatchPair{
					{
						Key:   "1",
						Value: "MatchedPacket",
					},
				},
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"OtherPacket":   {Name: "OtherPacket"},
			"MatchedPacket": {Name: "MatchedPacket"},
			"TestPacket":    packet,
		},
	}

	generator := NewRustGenerator(config, binModel)
	code := generator.generateUseCode(&packet)

	assert.Contains(t, code, "use crate::other_packet::*;")
	assert.Contains(t, code, "use crate::matched_packet::*;")
}

func TestGenerateStructCode(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
		Fields: []model.Field{
			{
				Name: "testField",
				Type: "u32",
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	generator := NewRustGenerator(config, binModel)
	code := generator.generateStructCode(&packet)

	assert.Contains(t, code, "pub struct TestPacket {")
	assert.Contains(t, code, "pub test_field: u32,")
	assert.Contains(t, code, "impl BinaryCodec for TestPacket")
	assert.Contains(t, code, "mod test_packet_tests")
}

func TestGetFieldType(t *testing.T) {
	tests := []struct {
		name     string
		parent   string
		field    model.Field
		expected string
	}{
		{
			name:     "string type",
			parent:   "Parent",
			field:    model.Field{Type: "string"},
			expected: "String",
		},
		{
			name:     "match type",
			parent:   "Parent",
			field:    model.Field{Type: "match", MatchKey: "Match"},
			expected: "ParentEnum",
		},
		{
			name:     "char array",
			parent:   "Parent",
			field:    model.Field{Type: "char[10]"},
			expected: "String",
		},
		{
			name:     "primitive type",
			parent:   "Parent",
			field:    model.Field{Type: "u32"},
			expected: "u32",
		},
	}
	g := NewRustGenerator(nil, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.GetFieldType(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetFieldName(t *testing.T) {
	tests := []struct {
		name     string
		field    model.Field
		expected string
	}{
		{
			name:     "match field",
			field:    model.Field{Name: "MatchField", Type: "match"},
			expected: "match_field",
		},
		{
			name:     "regular field",
			field:    model.Field{Name: "RegularField", Type: "u32"},
			expected: "regular_field",
		},
	}
	g := NewRustGenerator(nil, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.GetFieldName(tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRustEncodeField(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]model.Packet),
	}

	tests := []struct {
		name     string
		parent   string
		field    model.Field
		expected string
	}{
		{
			name:     "string list",
			parent:   "Parent",
			field:    model.Field{Name: "strings", Type: "string", IsRepeat: true},
			expected: "put_string_list::<u16,u8>(buf, &self.strings);",
		},
		{
			name:     "primitive list",
			parent:   "Parent",
			field:    model.Field{Name: "numbers", Type: "u32", IsRepeat: true},
			expected: "put_list::<u32,u16>(buf, &self.numbers);",
		},
		{
			name:     "string",
			parent:   "Parent",
			field:    model.Field{Name: "text", Type: "string"},
			expected: "put_string::<u8>(buf, &self.text);",
		},
		{
			name:     "primitive",
			parent:   "Parent",
			field:    model.Field{Name: "number", Type: "u32"},
			expected: "buf.put_u32(self.number);",
		},
		{
			name:   "match field",
			parent: "Parent",
			field: model.Field{
				Name:     "matchField",
				Type:     "match",
				MatchKey: "Type",
				MatchPairs: []model.MatchPair{
					{Value: "Value"},
				},
			},
			expected: "match &self.match_field {\n    ParentmatchFieldEnum::Value(msg) => msg.encode(buf),\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewRustGenerator(config, binModel)
			result := generator.EncodeField(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRustDecodeField(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]model.Packet),
	}

	tests := []struct {
		name     string
		parent   string
		field    model.Field
		expected string
	}{
		{
			name:     "string list",
			parent:   "Parent",
			field:    model.Field{Name: "strings", Type: "string", IsRepeat: true},
			expected: "let strings = get_string_list::<u16,u8>(buf)?;",
		},
		{
			name:     "primitive list",
			parent:   "Parent",
			field:    model.Field{Name: "numbers", Type: "u32", IsRepeat: true},
			expected: "let numbers = get_list::<u32,u16>(buf)?;",
		},
		{
			name:     "string",
			parent:   "Parent",
			field:    model.Field{Name: "text", Type: "string"},
			expected: "let text = get_string::<u8>(buf)?;",
		},
		{
			name:     "primitive",
			parent:   "Parent",
			field:    model.Field{Name: "number", Type: "u32"},
			expected: "let number = buf.get_u32();",
		},
		{
			name:   "match field",
			parent: "Parent",
			field: model.Field{
				Name:     "matchField",
				Type:     "match",
				MatchKey: "matchField",
				MatchPairs: []model.MatchPair{
					{Key: "\"1\"", Value: "Value"},
				},
			},
			expected: `let match_field = match match_field.as_str() {
    "1" => ParentmatchFieldEnum::Value(Value::decode(buf)?),
    _ => return None,
};`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewRustGenerator(config, binModel)
			result := generator.DecodeField(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenerateMatchFieldEnumCode(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
		Fields: []model.Field{
			{
				Name:     "Body",
				Type:     "match",
				MatchKey: "Type",
				MatchPairs: []model.MatchPair{
					{
						Key:   "1",
						Value: "FirstValue",
					},
					{
						Key:   "2",
						Value: "SecondValue",
					},
				},
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	generator := NewRustGenerator(config, binModel)
	code := generator.generateMatchFieldEnumCode(&packet)

	assert.Contains(t, code, "pub enum TestPacketBodyEnum")
	assert.Contains(t, code, "FirstValue(FirstValue)")
	assert.Contains(t, code, "SecondValue(SecondValue)")
}

func TestGenerateUnitTestCode(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
		Fields: []model.Field{
			{
				Name: "testField",
				Type: "u32",
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	generator := NewRustGenerator(config, binModel)
	code := generator.generateUnitTestCode(&packet)

	assert.Contains(t, code, "mod test_packet_tests")
	assert.Contains(t, code, "fn test_test_packet_codec()")
	assert.Contains(t, code, "let original = TestPacket")
	assert.Contains(t, code, "assert_eq!(original, decoded)")
}

func TestTestValue(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]model.Packet),
	}

	tests := []struct {
		name     string
		parent   string
		field    model.Field
		expected string
	}{
		{
			name:     "string list",
			parent:   "Parent",
			field:    model.Field{Name: "strings", Type: "string", IsRepeat: true},
			expected: `vec!["example".to_string(), "test".to_string()]`,
		},
		{
			name:     "u32 list",
			parent:   "Parent",
			field:    model.Field{Name: "numbers", Type: "u32", IsRepeat: true},
			expected: "vec![123456,654321]",
		},
		{
			name:     "string",
			parent:   "Parent",
			field:    model.Field{Name: "text", Type: "string"},
			expected: `"example".to_string()`,
		},
		{
			name:     "u32",
			parent:   "Parent",
			field:    model.Field{Name: "number", Type: "u32"},
			expected: "123456",
		},
		{
			name:     "char array",
			parent:   "Parent",
			field:    model.Field{Name: "chars", Type: "char[10]"},
			expected: "vec!['a'; 10].into_iter().collect::<String>()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewRustGenerator(config, binModel)
			result := generator.testValue(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHasQuotes(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`"quoted"`, true},
		{`notquoted`, false},
		{`"partial`, false},
		{`partial"`, false},
		{`""`, true},
		{`"`, false},
		{``, false},
	}
	g := NewRustGenerator(nil, nil)
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := g.HasQuotes(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
