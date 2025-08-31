package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

func TestNewLuaWspGenerator(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]model.Packet),
	}

	generator := NewLuaWspGenerator(config, binModel)

	assert.NotNil(t, generator)
	assert.Equal(t, config, generator.config)
	assert.Equal(t, binModel, generator.binModel)
}

func TestGenerateLua(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	rootPacket := model.Packet{
		Name:   "TestProtocol",
		IsRoot: true,
		Fields: []model.Field{
			{
				Name: "testField",
				Type: "u32",
			},
		},
	}
	binModel := &model.BinaryModel{
		RootPacket: &rootPacket,
		PacketsMap: map[string]model.Packet{
			"TestProtocol": rootPacket,
		},
	}

	generator := NewLuaWspGenerator(config, binModel)
	output, err := generator.Generate(binModel)

	assert.NoError(t, err)
	assert.NotEmpty(t, output)
	assert.Contains(t, output, "test_protocol.lua")
	assert.Contains(t, string(output["test_protocol.lua"]), "local test_protocol_proto")
	assert.Contains(t, string(output["test_protocol.lua"]), "function test_protocol_proto.dissector")
}

func TestGenerateSubDissector(t *testing.T) {
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

	generator := NewLuaWspGenerator(config, binModel)
	code := generator.generateSubDissector("Parent", packet)

	assert.Contains(t, code, "local function dissect_test_packet")
	assert.Contains(t, code, "local subtree = tree:add(parent_proto")
	assert.Contains(t, code, "subtree:add(fields.test_packet_test_field")
}

func TestGenerateMainDissector(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	rootPacket := model.Packet{
		Name:   "TestProtocol",
		IsRoot: true,
		Fields: []model.Field{
			{
				Name: "testField",
				Type: "u32",
			},
		},
	}
	binModel := &model.BinaryModel{
		RootPacket: &rootPacket,
		PacketsMap: map[string]model.Packet{
			"TestProtocol": rootPacket,
		},
	}

	generator := NewLuaWspGenerator(config, binModel)
	code := generator.generateMainDissector(&rootPacket)

	assert.Contains(t, code, "function test_protocol_proto.dissector")
	assert.Contains(t, code, "pinfo.cols.protocol = \"test_protocol\"")
	assert.Contains(t, code, "tree:add(fields.test_protocol_test_field")
}

func TestDecodeList(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
	}
	field := model.Field{
		Name:     "testList",
		Type:     "u32",
		IsRepeat: true,
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	generator := NewLuaWspGenerator(config, binModel)
	code := generator.decodeList("tree", &packet, field)

	assert.Contains(t, code, "local test_packet_test_list_size = buf(offset, 2):uint()")
	assert.Contains(t, code, "tree:add(\"testList Size: \".. test_packet_test_list_size")
	assert.Contains(t, code, "for i=1,test_packet_test_list_size do")
	assert.Contains(t, code, "tree:add(fields.test_packet_test_list")
}

func TestLuaDecodeField(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]model.Packet),
	}

	tests := []struct {
		name     string
		packet   model.Packet
		field    model.Field
		expected string
	}{
		{
			name:   "u32 field",
			packet: model.Packet{Name: "TestPacket"},
			field:  model.Field{Name: "testField", Type: "u32"},
			expected: `tree:add(fields.test_packet_test_field, buf(offset, 4))
offset = offset + 4`,
		},
		{
			name:   "string field",
			packet: model.Packet{Name: "TestPacket"},
			field:  model.Field{Name: "testField", Type: "string"},
			expected: `local test_packet_test_field_len = buf(offset, 1):uint()
tree:add("testField Len: ".. test_packet_test_field_len, buf(offset, 1))
offset = offset + 1
tree:add(fields.test_packet_test_field, buf(offset, test_packet_test_field_len))
offset = offset + test_packet_test_field_len`,
		},
		{
			name:   "match field",
			packet: model.Packet{Name: "TestPacket"},
			field: model.Field{
				Name:     "testField",
				Type:     "match",
				MatchKey: "msgType",
				MatchPairs: []model.MatchPair{
					{Key: "1", Value: "Message1"},
				},
			},
			expected: `if msg_type == 1 then -- Message1
    dissect_message_1(buf, pinfo, tree, offset)
    pinfo.cols.info:set("Message1")
end`,
		},
		{
			name:   "char array",
			packet: model.Packet{Name: "TestPacket"},
			field:  model.Field{Name: "testField", Type: "char[10]"},
			expected: `tree:add(fields.test_packet_test_field, buf(offset, 10))
offset = offset + 10`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewLuaWspGenerator(config, binModel)
			result := generator.decodeField("tree", &tt.packet, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDecodeListSize(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
	}
	field := model.Field{
		Name:     "testList",
		Type:     "u32",
		IsRepeat: true,
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	tests := []struct {
		name     string
		prefix   string
		expected string
	}{
		{
			name:   "u16 prefix",
			prefix: "u16",
			expected: `local test_packet_test_list_size = buf(offset, 2):uint()
tree:add("testList Size: ".. test_packet_test_list_size, buf(offset, 2))
offset = offset + 2`,
		},
		{
			name:   "u32 prefix",
			prefix: "u32",
			expected: `local test_packet_test_list_size = buf(offset, 4):uint()
tree:add("testList Size: ".. test_packet_test_list_size, buf(offset, 4))
offset = offset + 4`,
		},
		{
			name:     "unsupported prefix",
			prefix:   "invalid",
			expected: "-- unsupported numeric type: u32\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.ListLenPrefixLenType = tt.prefix
			generator := NewLuaWspGenerator(config, binModel)
			result := generator.decodeListSize("tree", &packet, field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDecodeStringLen(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
	}
	field := model.Field{
		Name: "testString",
		Type: "string",
	}
	binModel := &model.BinaryModel{
		PacketsMap: map[string]model.Packet{
			"TestPacket": packet,
		},
	}

	tests := []struct {
		name     string
		prefix   string
		expected string
	}{
		{
			name:   "u8 prefix",
			prefix: "u8",
			expected: `local test_packet_test_string_len = buf(offset, 1):uint()
tree:add("testString Len: ".. test_packet_test_string_len, buf(offset, 1))
offset = offset + 1`,
		},
		{
			name:   "u16 prefix",
			prefix: "u16",
			expected: `local test_packet_test_string_len = buf(offset, 2):uint()
tree:add("testString Len: ".. test_packet_test_string_len, buf(offset, 2))
offset = offset + 2`,
		},
		{
			name:     "unsupported prefix",
			prefix:   "invalid",
			expected: "-- unsupported numeric type: string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.StringLenPrefixLenType = tt.prefix
			generator := NewLuaWspGenerator(config, binModel)
			result := generator.decodeStringLen("tree", &packet, field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenerateFieldDefinition(t *testing.T) {
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
			{
				Name: "testString",
				Type: "string",
			},
		},
	}
	binModel := model.NewBinaryModel()
	binModel.AddPacket(packet)

	generator := NewLuaWspGenerator(config, binModel)
	code := generator.generateFieldDefinition(binModel)

	assert.Contains(t, code, "local fields = {")
	assert.Contains(t, code, "-- Field from TestPacket")
	assert.Contains(t, code, "test_packet_test_field = ProtoField.uint32(\"test_packet.test_field\", \"testField\", base.DEC),")
	assert.Contains(t, code, "test_packet_test_string = ProtoField.string(\"test_packet.test_string\", \"testString\"),")
}

func TestGenerateFieldDefinitionFromPacket(t *testing.T) {
	config := &GeneratorConfig{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u8",
	}
	packet := model.Packet{
		Name: "TestPacket",
		Fields: []model.Field{
			{
				Name: "test",
				Type: "u32",
			},
			{
				Name: "testString",
				Type: "string",
			},
			{
				Name: "testChar",
				Type: "char",
			},
			{
				Name: "testBytes",
				Type: "bytes",
			},
			{
				Name: "testBool",
				Type: "bool",
			},
		},
	}
	binModel := model.NewBinaryModel()
	binModel.AddPacket(packet)

	generator := NewLuaWspGenerator(config, binModel)
	code := generator.generateFieldDefinitionFromPacket(binModel, &packet)

	assert.Contains(t, code, "-- Field from TestPacket")
	assert.Contains(t, code, "test_packet_test = ProtoField.uint32(\"test_packet.test\", \"test\", base.DEC),")
	assert.Contains(t, code, "test_packet_test_string = ProtoField.string(\"test_packet.test_string\", \"testString\"),")
	assert.Contains(t, code, "test_packet_test_char = ProtoField.char(\"test_packet.test_char\", \"testChar\", base.OCT),")
	assert.Contains(t, code, "test_packet_test_bytes = ProtoField.bytes(\"test_packet.test_bytes\", \"testBytes\"),")
	assert.Contains(t, code, "test_packet_test_bool = ProtoField.bool(\"test_packet.test_bool\", \"testBool\"),")
}

func TestIsMatchField(t *testing.T) {
	tests := []struct {
		name       string
		field      model.Field
		rootPacket model.Packet
		expected   bool
	}{
		{
			name: "is match field",
			field: model.Field{
				Name: "msgType",
			},
			rootPacket: model.Packet{
				Fields: []model.Field{
					{
						Name:     "testField",
						Type:     "match",
						MatchKey: "msgType",
					},
				},
			},
			expected: true,
		},
		{
			name: "not match field",
			field: model.Field{
				Name: "otherField",
			},
			rootPacket: model.Packet{
				Fields: []model.Field{
					{
						Name:     "testField",
						Type:     "match",
						MatchKey: "msgType",
					},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatchField(tt.field, &tt.rootPacket)
			assert.Equal(t, tt.expected, result)
		})
	}
}
