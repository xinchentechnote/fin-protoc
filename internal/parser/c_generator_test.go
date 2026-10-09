package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

func newCTestModel() (*CGenerator, *model.BinaryModel) {
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]*model.Packet),
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}
	return NewCGenerator(binModel), binModel
}

func TestNewCGenerator(t *testing.T) {
	generator, binModel := newCTestModel()

	assert.NotNil(t, generator)
	assert.Equal(t, binModel.Config, generator.GetConfig())
}

func TestGenerateC(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "testField", Attr: &model.BasicFieldAttribute{Type: "u32"}},
		},
	}
	generator, binModel := newCTestModel()
	binModel.PacketsMap["TestPacket"] = packet
	binModel.Packets = append(binModel.Packets, packet)

	output, err := generator.Generate(binModel)

	assert.NoError(t, err)
	assert.NotEmpty(t, output)
	assert.Contains(t, output, "test_packet.h")
	assert.Contains(t, output, "test_packet.c")
	assert.Contains(t, output, "root.h")
	assert.Contains(t, output, "tests.c")
}

func TestGenerateCPacketHeader(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "testField", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			{Name: "name", Attr: &model.FixedStringFieldAttribute{Length: 10}},
		},
	}
	generator, _ := newCTestModel()
	code := generator.generatePacketHeader(packet)

	assert.Contains(t, code, "#ifndef FC_TEST_PACKET_H")
	assert.Contains(t, code, "uint32_t test_field;")
	assert.Contains(t, code, "fc_string name;")
	assert.Contains(t, code, "bool test_packet_encode(const test_packet *self, fc_buffer *buf);")
}

func TestGenerateCPacketSource(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "testField", Attr: &model.BasicFieldAttribute{Type: "u32"}},
		},
	}
	generator, _ := newCTestModel()
	code := generator.generatePacketSource(packet)

	assert.Contains(t, code, `#include "test_packet.h"`)
	assert.Contains(t, code, "if (!fc_buffer_put_u32(buf, self->test_field)) return false;")
	assert.Contains(t, code, "if (!fc_reader_get_u32(r, &out->test_field)) return false;")
	assert.Contains(t, code, "void test_packet_free(test_packet *self) {")
}

func TestGenerateCRootHeader(t *testing.T) {
	packet := &model.Packet{Name: "TestPacket"}
	generator, binModel := newCTestModel()
	binModel.PacketsMap["TestPacket"] = packet
	binModel.Packets = append(binModel.Packets, packet)

	code := generator.generateRootHeader()

	assert.Contains(t, code, "#ifndef FC_ROOT_H")
	assert.Contains(t, code, `#include "test_packet.h"`)
}

func TestGenerateCTests(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "testField", Attr: &model.BasicFieldAttribute{Type: "u32"}},
		},
	}
	generator, binModel := newCTestModel()
	binModel.PacketsMap["TestPacket"] = packet
	binModel.Packets = append(binModel.Packets, packet)

	code := generator.generateTests()

	assert.Contains(t, code, "static void test_test_packet_codec(void)")
	assert.Contains(t, code, "msg.test_field = 123456;")
	assert.Contains(t, code, "CHECK(test_packet_encode(&msg, &buf));")
	assert.Contains(t, code, "CHECK(test_packet_eql(&msg, &decoded));")
	assert.Contains(t, code, "int main(void) {")
}

func TestGenerateCMatchKindCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "Body", Attr: &model.MatchFieldAttribute{
				MatchPairs: []model.MatchPair{
					{Key: "1", Value: "FirstValue"},
					{Key: "2", Value: "SecondValue"},
				},
			}},
		},
	}
	generator, _ := newCTestModel()
	code := generator.generateMatchKindCode(packet)

	assert.Contains(t, code, "TEST_PACKET_BODY_NONE = 0,")
	assert.Contains(t, code, "TEST_PACKET_BODY_FIRST_VALUE,")
	assert.Contains(t, code, "} test_packet_body_kind;")
}

func TestCGetFieldType(t *testing.T) {
	g, _ := newCTestModel()

	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{
			name:     "string type",
			field:    &model.Field{Attr: &model.DynamicStringFieldAttribute{}},
			expected: "fc_string",
		},
		{
			name:     "char array",
			field:    &model.Field{Attr: &model.FixedStringFieldAttribute{Length: 10}},
			expected: "fc_string",
		},
		{
			name:     "primitive type",
			field:    &model.Field{Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "uint32_t",
		},
		{
			name:     "char scalar",
			field:    &model.Field{Attr: &model.BasicFieldAttribute{Type: "char"}},
			expected: "uint8_t",
		},
		{
			name:     "float",
			field:    &model.Field{Attr: &model.BasicFieldAttribute{Type: "f64"}},
			expected: "double",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.cFieldTypeName(&model.Packet{Name: "P"}, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCGetFieldName(t *testing.T) {
	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{name: "match field", field: &model.Field{Name: "MatchField"}, expected: "match_field"},
		{name: "regular field", field: &model.Field{Name: "RegularField"}, expected: "regular_field"},
		{name: "c keyword escapes", field: &model.Field{Name: "Switch"}, expected: "_switch"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cSnake(tt.field.Name)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCEncodeField(t *testing.T) {
	g, binModel := newCTestModel()
	root := &model.Packet{Name: "TestPacket", LengthField: &model.Field{Name: "BodyLen",
		Attr: &model.LengthFieldAttribute{LengthType: "u32"}}}
	binModel.RootPacket = root

	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{
			name:     "primitive",
			field:    &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "if (!fc_buffer_put_u32(buf, self->number)) return false;",
		},
		{
			name:     "char scalar",
			field:    &model.Field{Name: "side", Attr: &model.BasicFieldAttribute{Type: "char"}},
			expected: "if (!fc_buffer_put_u8(buf, self->side)) return false;",
		},
		{
			name:     "char array",
			field:    &model.Field{Name: "chars", Attr: &model.FixedStringFieldAttribute{Length: 10}},
			expected: "if (!fc_put_char_array(buf, self->chars.data, self->chars.len, 10)) return false;",
		},
		{
			name: "char array nul pad",
			field: &model.Field{Name: "code", Attr: &model.FixedStringFieldAttribute{Length: 8,
				Padding: &model.Padding{PadChar: "'\x00'", PadLeft: false}}},
			expected: `if (!fc_put_char_array_pad(buf, self->code.data, self->code.len, 8, '\0', 0)) return false;`,
		},
		{
			name:     "primitive list",
			field:    &model.Field{Name: "numbers", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
			expected: "if (!fc_buffer_put_u32(buf, self->numbers[i])) return false;",
		},
		{
			name:     "checksum",
			field:    &model.Field{Name: "checksum", Attr: &model.CheckSumFieldAttribute{CheckSumType: `"SSE_BIN"`, Type: "u32"}},
			expected: "const uint32_t checksum_val = fc_sse_bin(buf->data, buf->len);",
		},
		{
			name:     "length field",
			field:    &model.Field{Name: "BodyLen", Attr: &model.LengthFieldAttribute{LengthType: "u32"}},
			expected: "const size_t body_len_pos = buf->len;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.encodeField(root, tt.field)
			assert.Contains(t, result, tt.expected)
		})
	}
}

func TestCDecodeField(t *testing.T) {
	g, _ := newCTestModel()

	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{
			name:     "string list",
			field:    &model.Field{Name: "strings", IsRepeat: true, Attr: &model.DynamicStringFieldAttribute{}},
			expected: "if (!fc_get_string(r, 1, &out->strings[i])) return false;",
		},
		{
			name:     "primitive list",
			field:    &model.Field{Name: "numbers", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
			expected: "if (!fc_reader_get_u32(r, &out->numbers[i])) return false;",
		},
		{
			name:     "string",
			field:    &model.Field{Name: "text", Attr: &model.DynamicStringFieldAttribute{}},
			expected: "if (!fc_get_string(r, 1, &out->text)) return false;",
		},
		{
			name:     "primitive",
			field:    &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "if (!fc_reader_get_u32(r, &out->number)) return false;",
		},
		{
			name: "match field numeric key",
			field: &model.Field{
				Name: "body",
				Attr: &model.MatchFieldAttribute{MatchKeyField: &model.Field{Name: "msgType"},
					MatchPairs: []model.MatchPair{
						{Key: "1", Value: "Logon"},
					},
				}},
			expected: "out->body_tag = TEST_PACKET_BODY_LOGON;",
		},
		{
			name: "match field string key",
			field: &model.Field{
				Name: "sess",
				Attr: &model.MatchFieldAttribute{MatchKeyField: &model.Field{Name: "sessKind",
					Attr: &model.DynamicStringFieldAttribute{}},
					MatchPairs: []model.MatchPair{
						{Key: `"fast"`, Value: "StrA"},
					},
				}},
			expected: `if (fc_streq_lit(out->sess_kind, "fast")) {`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.decodeField(&model.Packet{Name: "TestPacket"}, tt.field)
			assert.Contains(t, result, tt.expected)
		})
	}
}

func TestCChecksumCall(t *testing.T) {
	g, _ := newCTestModel()

	tests := []struct {
		algo     string
		typ      string
		expected string
		ok       bool
	}{
		{algo: "SSE_BIN", typ: "u32", expected: "fc_sse_bin", ok: true},
		{algo: "SZSE_BIN", typ: "i32", expected: "fc_szse_bin", ok: true},
		{algo: "CRC16", typ: "u16", expected: "fc_crc16", ok: true},
		{algo: "CRC32", typ: "u32", expected: "fc_crc32", ok: true},
		{algo: "SSE_BIN", typ: "i32", expected: "", ok: false},
		{algo: "CRC32", typ: "i32", expected: "", ok: false},
		{algo: "NOPE", typ: "u32", expected: "", ok: false},
	}
	for _, tt := range tests {
		t.Run(tt.algo+"/"+tt.typ, func(t *testing.T) {
			fn, ok := g.checksumCall(tt.algo, tt.typ)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.expected, fn)
		})
	}
}

func TestCHasQuotes(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`"quoted"`, true},
		{`notquoted`, false},
		{`"`, false},
		{``, false},
	}
	g, _ := newCTestModel()
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := g.HasQuotes(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
