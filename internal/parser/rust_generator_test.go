package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

func TestNewRustGenerator(t *testing.T) {

	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]*model.Packet),
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)

	assert.NotNil(t, generator)
	assert.Equal(t, binModel, generator.binModel)
}

func TestGenerateRust(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{
				Name: "testField",
				Attr: &model.BasicFieldAttribute{Type: "u32"},
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"TestPacket": packet,
		}, Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)
	output, err := generator.Generate(binModel)

	assert.NoError(t, err)
	assert.NotEmpty(t, output)
	assert.Contains(t, output, "test_packet.rs")
	assert.Contains(t, output, "borrowed_reader.rs")
	assert.Contains(t, output, "lib.rs")
	assert.Contains(t, string(output["borrowed_reader.rs"]), "pub struct BorrowedReader<'a>")
}

func TestGenerateRustFileForPacket(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{
				Name: "testField",
				Attr: &model.BasicFieldAttribute{Type: "u32"},
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"TestPacket": packet,
		}, Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)
	code := generator.generateRustFileForPacket(packet)

	assert.Contains(t, code, "pub struct TestPacket")
	assert.Contains(t, code, "impl BinaryCodec for TestPacket")
	assert.Contains(t, code, "mod test_packet_tests")
}

func TestGenerateRustFixedStringFastPath(t *testing.T) {
	field := &model.Field{
		Name: "Code",
		Attr: &model.FixedStringFieldAttribute{Length: 10},
	}
	packet := &model.Packet{Name: "TestPacket", Fields: []*model.Field{field}}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{"TestPacket": packet},
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
			Padding:                &model.Padding{PadChar: "' '", PadLeft: false},
		},
	}

	generator := NewRustGenerator(binModel)
	code := generator.generateRustFileForPacket(packet)

	assert.Contains(t, code, "fn put_fixed_string_field(")
	assert.Contains(t, code, "fn get_fixed_string_field(")
	assert.Contains(t, code, "String::from_utf8(bytes.slice(start..end).to_vec()).ok()")
	assert.Contains(t, code, "put_fixed_string_field(buf, &self.code, 10, b' ', false);")
	assert.Contains(t, code, "let code = get_fixed_string_field(buf, 10, b' ', false)?;")
	assert.NotContains(t, code, "fn put_fixed_string_list_field")
}

func TestGenerateLibCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"TestPacket": packet,
		}, Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)
	code := generator.generateLibCode()

	assert.Equal(t, "pub mod borrowed_reader;\npub mod test_packet;\n", code)
}

func TestGenerateRustBorrowedFixedAndDynamicStrings(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "Code", Attr: &model.FixedStringFieldAttribute{Length: 10}},
			{Name: "Text", Attr: &model.DynamicStringFieldAttribute{}},
			{Name: "Number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{"TestPacket": packet},
		Config: &model.Configuration{
			StringLenPrefixLenType: "u16",
			ListLenPrefixLenType:   "u16",
			LittleEndian:           true,
			Padding:                &model.Padding{PadChar: "' '", PadLeft: false},
		},
	}

	code := NewRustGenerator(binModel).generateRustFileForPacket(packet)

	assert.Contains(t, code, "pub struct TestPacketRef<'a>")
	assert.Contains(t, code, "pub code: &'a str,")
	assert.Contains(t, code, "pub text: &'a str,")
	assert.Contains(t, code, "pub fn from_slice(input: &'a [u8]) -> Option<Self>")
	assert.Contains(t, code, "assert!(TestPacketRef::from_slice(&bytes).is_some());")
	assert.Contains(t, code, "pub fn decode_into(buf: &mut Bytes, out: &mut Self) -> Option<()>")
	assert.Contains(t, code, "If decoding fails, `out` may have been partially updated.")
	assert.Contains(t, code, "get_fixed_string_field_into(buf, &mut out.code, 10, b' ', false)?;")
	assert.Contains(t, code, "get_string_field_into::<u16>(buf, &mut out.text, true)?;")
	assert.Contains(t, code, "TestPacket::decode_into(&mut reuse_bytes, &mut reused).unwrap();")
	assert.Contains(t, code, "assert_eq!(original, reused);")
	assert.Contains(t, code, "let code = reader.get_fixed_str(10, b' ', false)?;")
	assert.Contains(t, code, "let text_wire_len = reader.get_u16_le()? as usize;")
	assert.Contains(t, code, "let text = reader.get_str(text_wire_len)?;")
	assert.Contains(t, code, "let number = reader.get_u32_le()?;")
}

func TestGenerateRustBorrowedDynamicStringDoesNotShadowLengthField(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "TextLen", Attr: &model.LengthFieldAttribute{LengthType: "u32"}},
			{Name: "Text", Attr: &model.DynamicStringFieldAttribute{}},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{"TestPacket": packet},
		Config: &model.Configuration{
			StringLenPrefixLenType: "u32",
			ListLenPrefixLenType:   "u32",
		},
	}

	code := NewRustGenerator(binModel).generateRustFileForPacket(packet)

	assert.Contains(t, code, "let text_len = reader.get_u32()?;")
	assert.Contains(t, code, "let text_wire_len = reader.get_u32()? as usize;")
	assert.Contains(t, code, "let text = reader.get_str(text_wire_len)?;")
	assert.Contains(t, code, "text_len,\n            text,")
}

func TestGenerateRustBorrowedMatch(t *testing.T) {
	matchKey := &model.Field{Name: "Kind", Attr: &model.FixedStringFieldAttribute{Length: 3}}
	first := &model.Packet{Name: "FirstValue", Fields: []*model.Field{
		{Name: "Value", Attr: &model.BasicFieldAttribute{Type: "u32"}},
	}}
	second := &model.Packet{Name: "SecondValue"}
	body := &model.Field{
		Name: "Body",
		Attr: &model.MatchFieldAttribute{
			MatchKeyField: matchKey,
			MatchPairs: []model.MatchPair{
				{Key: "\"001\"", Value: "FirstValue"},
				{Key: "\"002\"", Value: "SecondValue"},
			},
		},
	}
	packet := &model.Packet{Name: "Envelope", Fields: []*model.Field{matchKey, body}}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"Envelope":    packet,
			"FirstValue":  first,
			"SecondValue": second,
		},
		Config: &model.Configuration{
			StringLenPrefixLenType: "u16",
			ListLenPrefixLenType:   "u16",
			Padding:                &model.Padding{PadChar: "' '", PadLeft: false},
		},
	}

	code := NewRustGenerator(binModel).generateRustFileForPacket(packet)

	assert.Contains(t, code, "pub enum EnvelopeBodyRefEnum<'a>")
	assert.Contains(t, code, "FirstValue(FirstValueRef<'a>),")
	assert.Contains(t, code, "let body = match kind {")
	assert.Contains(t, code, "\"001\" => EnvelopeBodyRefEnum::FirstValue(FirstValueRef::decode_ref(reader)?),")
}

func TestGenerateRustBorrowedSkipsRepeatPackets(t *testing.T) {
	packet := &model.Packet{Name: "RepeatedPacket", Fields: []*model.Field{
		{Name: "Values", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
	}}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{"RepeatedPacket": packet},
		Config: &model.Configuration{
			StringLenPrefixLenType: "u16",
			ListLenPrefixLenType:   "u16",
		},
	}

	code := NewRustGenerator(binModel).generateRustFileForPacket(packet)

	assert.NotContains(t, code, "pub struct RepeatedPacketRef<'a>")
	assert.NotContains(t, code, "RepeatedPacketRef::from_slice")
	assert.NotContains(t, code, "pub fn decode_into")
}

func TestGenerateRustBorrowedSupportsInlineChildOfRepeatPacket(t *testing.T) {
	child := &model.Packet{Name: "InlineItem", Fields: []*model.Field{
		{Name: "Code", Attr: &model.FixedStringFieldAttribute{Length: 4}},
	}}
	packet := &model.Packet{Name: "RepeatedPacket", Fields: []*model.Field{
		{Name: "Items", Attr: &model.ObjectFieldAttribute{IsIner: true, RefPacket: child}, IsRepeat: true},
	}}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{"RepeatedPacket": packet},
		Config: &model.Configuration{
			StringLenPrefixLenType: "u16",
			ListLenPrefixLenType:   "u16",
			Padding:                &model.Padding{PadChar: "' '", PadLeft: false},
		},
	}

	code := NewRustGenerator(binModel).generateRustFileForPacket(packet)

	assert.Contains(t, code, "use crate::borrowed_reader::BorrowedReader;")
	assert.Contains(t, code, "pub struct InlineItemRef<'a>")
	assert.Contains(t, code, "InlineItemRef::from_slice")
	assert.Contains(t, code, "impl InlineItem {")
	assert.Contains(t, code, "pub fn decode_into")
	assert.NotContains(t, code, "pub struct RepeatedPacketRef<'a>")
}

func TestGenerateUseCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{
				Name: "testField",
				Attr: &model.ObjectFieldAttribute{PacketName: "OtherPacket"},
			},
			{
				Name: "matchField",
				Attr: &model.MatchFieldAttribute{
					MatchPairs: []model.MatchPair{
						{
							Key:   "1",
							Value: "MatchedPacket",
						},
					},
					MatchKeyField: &model.Field{Name: "matchField"},
				},
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"OtherPacket":   {Name: "OtherPacket"},
			"MatchedPacket": {Name: "MatchedPacket"},
			"TestPacket":    packet,
		}, Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)
	code := generator.generateUseCode(packet)

	assert.Contains(t, code, "use crate::other_packet::*;")
	assert.Contains(t, code, "use crate::matched_packet::*;")
}

func TestGenerateStructCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{
				Name: "testField",
				Attr: &model.BasicFieldAttribute{Type: "u32"},
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"TestPacket": packet,
		},
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)
	code := generator.generateStructCode(packet)

	assert.Contains(t, code, "pub struct TestPacket {")
	assert.Contains(t, code, "pub test_field: u32,")
	assert.Contains(t, code, "impl BinaryCodec for TestPacket")
	assert.Contains(t, code, "#[inline]\n    fn encode")
	assert.Contains(t, code, "#[inline]\n    fn decode")
	assert.Contains(t, code, "mod test_packet_tests")
}

func TestGetFieldType(t *testing.T) {
	tests := []struct {
		name     string
		parent   string
		field    *model.Field
		expected string
	}{
		{
			name:     "string type",
			parent:   "Parent",
			field:    &model.Field{Attr: &model.DynamicStringFieldAttribute{}},
			expected: "String",
		},
		{
			name:     "match type",
			parent:   "Parent",
			field:    &model.Field{Attr: &model.MatchFieldAttribute{}},
			expected: "ParentEnum",
		},
		{
			name:     "char array",
			parent:   "Parent",
			field:    &model.Field{Attr: &model.FixedStringFieldAttribute{Length: 10}},
			expected: "String",
		},
		{
			name:     "primitive type",
			parent:   "Parent",
			field:    &model.Field{Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "u32",
		},
	}
	g := NewRustGenerator(&model.BinaryModel{})
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
		field    *model.Field
		expected string
	}{
		{
			name:     "match field",
			field:    &model.Field{Name: "MatchField"},
			expected: "match_field",
		},
		{
			name:     "regular field",
			field:    &model.Field{Name: "RegularField"},
			expected: "regular_field",
		},
	}
	g := NewRustGenerator(&model.BinaryModel{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.GetFieldName(tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRustDecodeField(t *testing.T) {
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]*model.Packet),
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	tests := []struct {
		name     string
		parent   string
		field    *model.Field
		expected string
	}{
		{
			name:   "string list",
			parent: "Parent",
			field: &model.Field{Name: "strings", IsRepeat: true,
				Attr: &model.DynamicStringFieldAttribute{}},
			expected: "let strings = get_string_list::<u16,u8>(buf)?;",
		},
		{
			name:     "primitive list",
			parent:   "Parent",
			field:    &model.Field{Name: "numbers", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
			expected: "let numbers = get_list::<u32,u16>(buf)?;",
		},
		{
			name:     "string",
			parent:   "Parent",
			field:    &model.Field{Name: "text", Attr: &model.DynamicStringFieldAttribute{}},
			expected: "let text = get_string::<u8>(buf)?;",
		},
		{
			name:     "primitive",
			parent:   "Parent",
			field:    &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "let number = buf.get_u32();",
		},
		{
			name:   "match field",
			parent: "Parent",
			field: &model.Field{
				Name: "matchField",
				Attr: &model.MatchFieldAttribute{MatchKeyField: &model.Field{Name: "matchField",
					Attr: &model.DynamicStringFieldAttribute{}},
					MatchPairs: []model.MatchPair{
						{Key: "\"1\"", Value: "Value"},
					},
				}},
			expected: `let match_field = match match_field.as_str() {
    "1" => ParentmatchFieldEnum::Value(Value::decode(buf)?),
    _ => return None,
};`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewRustGenerator(binModel)
			result := generator.DecodeField(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenerateMatchFieldEnumCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{
				Name: "Body",
				Attr: &model.MatchFieldAttribute{
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
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"TestPacket": packet,
		},
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)
	code := generator.generateMatchFieldEnumCode(packet)

	assert.Contains(t, code, "pub enum TestPacketBodyEnum")
	assert.Contains(t, code, "FirstValue(FirstValue)")
	assert.Contains(t, code, "SecondValue(SecondValue)")
}

func TestGenerateUnitTestCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{
				Name: "testField",
				Attr: &model.BasicFieldAttribute{Type: "u32"},
			},
		},
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]*model.Packet{
			"TestPacket": packet,
		},
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	generator := NewRustGenerator(binModel)
	code := generator.generateUnitTestCode(packet)

	assert.Contains(t, code, "mod test_packet_tests")
	assert.Contains(t, code, "fn test_test_packet_codec()")
	// No length/checksum fields to fix up, so `original` needs no `mut`.
	assert.Contains(t, code, "let original = TestPacket")
	assert.NotContains(t, code, "let mut original")
	assert.Contains(t, code, "assert_eq!(original, decoded)")
}

func TestTestValue(t *testing.T) {
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]*model.Packet),
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}

	tests := []struct {
		name     string
		parent   string
		field    *model.Field
		expected string
	}{
		{
			name:     "string list",
			parent:   "Parent",
			field:    &model.Field{Name: "strings", IsRepeat: true, Attr: &model.DynamicStringFieldAttribute{}},
			expected: `vec!["example".to_string(), "test".to_string()]`,
		},
		{
			name:     "u32 list",
			parent:   "Parent",
			field:    &model.Field{Name: "numbers", IsRepeat: true, Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "vec![123456,654321]",
		},
		{
			name:     "string",
			parent:   "Parent",
			field:    &model.Field{Name: "text", Attr: &model.DynamicStringFieldAttribute{}},
			expected: `"example".to_string()`,
		},
		{
			name:     "u32",
			parent:   "Parent",
			field:    &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "123456",
		},
		{
			name:   "char array",
			parent: "Parent",
			field: &model.Field{Name: "chars",
				Attr: &model.FixedStringFieldAttribute{Length: 10,
					Padding: &model.Padding{PadChar: "' '", PadLeft: false},
				}},
			expected: "\"a\".repeat(10)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewRustGenerator(binModel)
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
	g := NewRustGenerator(&model.BinaryModel{})
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := g.HasQuotes(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
