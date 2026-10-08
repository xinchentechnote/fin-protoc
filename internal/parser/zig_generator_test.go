package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

func newZigTestModel() (*ZigGenerator, *model.BinaryModel) {
	binModel := &model.BinaryModel{
		PacketsMap: make(map[string]*model.Packet),
		Config: &model.Configuration{
			ListLenPrefixLenType:   "u16",
			StringLenPrefixLenType: "u8",
		},
	}
	return NewZigGenerator(binModel), binModel
}

func TestNewZigGenerator(t *testing.T) {
	generator, binModel := newZigTestModel()

	assert.NotNil(t, generator)
	assert.Equal(t, binModel, generator.binModel)
	assert.Equal(t, binModel.Config, generator.GetConfig())
}

func TestGenerateZig(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "testField", Attr: &model.BasicFieldAttribute{Type: "u32"}},
		},
	}
	generator, binModel := newZigTestModel()
	binModel.PacketsMap["TestPacket"] = packet
	binModel.Packets = append(binModel.Packets, packet)

	output, err := generator.Generate(binModel)

	assert.NoError(t, err)
	assert.NotEmpty(t, output)
	assert.Contains(t, output, "test_packet.zig")
	assert.Contains(t, output, "root.zig")
}

func TestGenerateZigFileForPacket(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "testField", Attr: &model.BasicFieldAttribute{Type: "u32"}},
		},
	}
	generator, _ := newZigTestModel()
	code := generator.generateZigFileForPacket(packet)

	assert.Contains(t, code, "pub const TestPacket = struct {")
	assert.Contains(t, code, "test_field: u32,")
	assert.Contains(t, code, "pub fn encode(self: *const TestPacket, buf: *codec.Buffer) !void {")
	assert.Contains(t, code, "pub fn decode(allocator: std.mem.Allocator, r: *codec.Reader) codec.DecodeError!TestPacket {")
	assert.Contains(t, code, `test "test_test_packet_codec" {`)
	assert.Contains(t, code, "const codec = @import(\"binary_codec\");")
}

func TestGenerateZigRootCode(t *testing.T) {
	packet := &model.Packet{Name: "TestPacket"}
	generator, binModel := newZigTestModel()
	binModel.PacketsMap["TestPacket"] = packet
	binModel.Packets = append(binModel.Packets, packet)

	code := generator.generateRootCode()

	assert.Contains(t, code, `pub const test_packet = @import("test_packet.zig");`)
	assert.Contains(t, code, "pub const TestPacket = test_packet.TestPacket;")
	assert.Contains(t, code, "test {\n    _ = test_packet;\n}")
}

func TestGenerateZigImportCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "objField", Attr: &model.ObjectFieldAttribute{PacketName: "OtherPacket"}},
			{Name: "matchField", Attr: &model.MatchFieldAttribute{
				MatchPairs:    []model.MatchPair{{Key: "1", Value: "MatchedPacket"}},
				MatchKeyField: &model.Field{Name: "matchField"},
			}},
		},
	}
	generator, _ := newZigTestModel()
	code := generator.generateImportCode(packet)

	assert.Contains(t, code, `const OtherPacket = @import("other_packet.zig").OtherPacket;`)
	assert.Contains(t, code, `const MatchedPacket = @import("matched_packet.zig").MatchedPacket;`)
}

func TestGenerateZigMatchFieldEnumCode(t *testing.T) {
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
	generator, _ := newZigTestModel()
	code := generator.generateMatchFieldEnumCode(packet)

	assert.Contains(t, code, "pub const TestPacketBodyEnum = union(enum) {")
	assert.Contains(t, code, "first_value: FirstValue,")
	assert.Contains(t, code, "second_value: SecondValue,")
	assert.Contains(t, code, "pub fn eql(self: TestPacketBodyEnum, other: TestPacketBodyEnum) bool {")
}

func TestZigGetFieldType(t *testing.T) {
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
			expected: "[]const u8",
		},
		{
			name:     "match type",
			parent:   "Parent",
			field:    &model.Field{Name: "Body", Attr: &model.MatchFieldAttribute{}},
			expected: "ParentBodyEnum",
		},
		{
			name:     "char array",
			parent:   "Parent",
			field:    &model.Field{Attr: &model.FixedStringFieldAttribute{Length: 10}},
			expected: "[]const u8",
		},
		{
			name:     "primitive type",
			parent:   "Parent",
			field:    &model.Field{Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "u32",
		},
		{
			name:     "char scalar",
			parent:   "Parent",
			field:    &model.Field{Attr: &model.BasicFieldAttribute{Type: "char"}},
			expected: "u8",
		},
	}
	g, _ := newZigTestModel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.GetFieldType(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestZigGetFieldName(t *testing.T) {
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
		{
			name:     "zig keyword escapes",
			field:    &model.Field{Name: "Type"},
			expected: `@"type"`,
		},
	}
	g, _ := newZigTestModel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.GetFieldName(tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestZigEncodeField(t *testing.T) {
	g, _ := newZigTestModel()

	tests := []struct {
		name     string
		packet   *model.Packet
		field    *model.Field
		expected string
	}{
		{
			name:   "primitive",
			packet: &model.Packet{Name: "Test"},
			field:  &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "try buf.put(u32, self.number);",
		},
		{
			name:   "char scalar",
			packet: &model.Packet{Name: "Test"},
			field:  &model.Field{Name: "side", Attr: &model.BasicFieldAttribute{Type: "char"}},
			expected: "try buf.put(u8, self.side);",
		},
		{
			name:   "char array",
			packet: &model.Packet{Name: "Test"},
			field:  &model.Field{Name: "chars", Attr: &model.FixedStringFieldAttribute{Length: 10}},
			expected: "try codec.putCharArray(buf, self.chars, 10);",
		},
		{
			name:   "char array nul pad",
			packet: &model.Packet{Name: "Test"},
			field: &model.Field{Name: "code", Attr: &model.FixedStringFieldAttribute{Length: 8,
				Padding: &model.Padding{PadChar: "'\x00'", PadLeft: false}}},
			expected: `try codec.putCharArrayWithPad(buf, self.code, 8, '\x00', false);`,
		},
		{
			name:   "primitive list",
			packet: &model.Packet{Name: "Test"},
			field:  &model.Field{Name: "numbers", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
			expected: "try codec.putList(u32, u16, buf, self.numbers);",
		},
		{
			name:   "checksum",
			packet: &model.Packet{Name: "Test"},
			field:  &model.Field{Name: "checksum", Attr: &model.CheckSumFieldAttribute{CheckSumType: `"SSE_BIN"`, Type: "u32"}},
			expected: `try buf.put(u32, codec.computeChecksum(u32, "SSE_BIN", buf.written(), self.checksum));`,
		},
		{
			name: "length field",
			packet: &model.Packet{Name: "Test", LengthField: &model.Field{Name: "BodyLen",
				Attr: &model.LengthFieldAttribute{LengthType: "u32"}}},
			field: &model.Field{Name: "BodyLen", Attr: &model.LengthFieldAttribute{LengthType: "u32"}},
			expected: "const body_len_pos = buf.len;\ntry buf.put(u32, 0);",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.EncodeField(tt.packet, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestZigDecodeField(t *testing.T) {
	g, _ := newZigTestModel()

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
			expected: "const strings = try codec.getStringList(u16,u8, allocator, r);",
		},
		{
			name:     "primitive list",
			parent:   "Parent",
			field:    &model.Field{Name: "numbers", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
			expected: "const numbers = try codec.getList(u32, u16, allocator, r);",
		},
		{
			name:     "string",
			parent:   "Parent",
			field:    &model.Field{Name: "text", Attr: &model.DynamicStringFieldAttribute{}},
			expected: "const text = try codec.getString(u8, r);",
		},
		{
			name:     "primitive",
			parent:   "Parent",
			field:    &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "const number = try r.get(u32);",
		},
		{
			name:   "match field numeric key",
			parent: "Parent",
			field: &model.Field{
				Name: "matchField",
				Attr: &model.MatchFieldAttribute{MatchKeyField: &model.Field{Name: "matchField"},
					MatchPairs: []model.MatchPair{
						{Key: "1", Value: "Value"},
					},
				}},
			expected: `const match_field: ParentmatchFieldEnum = switch (match_field) {
    1 => .{ .value = try Value.decode(allocator, r) },
    else => return error.UnknownMessageType,
};`,
		},
		{
			name:   "match field string key",
			parent: "Parent",
			field: &model.Field{
				Name: "matchField",
				Attr: &model.MatchFieldAttribute{MatchKeyField: &model.Field{Name: "matchField",
					Attr: &model.DynamicStringFieldAttribute{}},
					MatchPairs: []model.MatchPair{
						{Key: `"1"`, Value: "Value"},
					},
				}},
			expected: `const match_field: ParentmatchFieldEnum = if (std.mem.eql(u8, match_field, "1"))
    .{ .value = try Value.decode(allocator, r) }
else
    return error.UnknownMessageType;`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.DecodeField(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestZigEqlField(t *testing.T) {
	g, _ := newZigTestModel()

	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{
			name:     "primitive",
			field:    &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "if (self.number != other.number) return false;",
		},
		{
			name:     "string",
			field:    &model.Field{Name: "text", Attr: &model.DynamicStringFieldAttribute{}},
			expected: "if (!std.mem.eql(u8, self.text, other.text)) return false;",
		},
		{
			name:     "primitive list",
			field:    &model.Field{Name: "numbers", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
			expected: "if (self.numbers.len != other.numbers.len) return false;\nfor (self.numbers, other.numbers) |a, b| {\n    if (a != b) return false;\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.EqlField(tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestZigGenerateUnitTestCode(t *testing.T) {
	packet := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "testField", Attr: &model.BasicFieldAttribute{Type: "u32"}},
		},
	}
	generator, _ := newZigTestModel()
	code := generator.generateUnitTestCode(packet)

	assert.Contains(t, code, `test "test_test_packet_codec" {`)
	assert.Contains(t, code, "const msg = TestPacket{")
	assert.Contains(t, code, ".test_field = 123456,")
	assert.Contains(t, code, "const decoded = try TestPacket.decode(allocator, &reader);")
	assert.Contains(t, code, "try std.testing.expect(msg.eql(decoded));")
}

func TestZigTestValue(t *testing.T) {
	g, _ := newZigTestModel()

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
			expected: `&[_][]const u8{ "example", "test" }`,
		},
		{
			name:     "u32 list",
			parent:   "Parent",
			field:    &model.Field{Name: "numbers", IsRepeat: true, Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "&[_]u32{ 123456, 654321 }",
		},
		{
			name:     "string",
			parent:   "Parent",
			field:    &model.Field{Name: "text", Attr: &model.DynamicStringFieldAttribute{}},
			expected: `"example"`,
		},
		{
			name:     "u32",
			parent:   "Parent",
			field:    &model.Field{Name: "number", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			expected: "123456",
		},
		{
			name:     "char array",
			parent:   "Parent",
			field:    &model.Field{Name: "chars", Attr: &model.FixedStringFieldAttribute{Length: 3}},
			expected: `"aaa"`,
		},
		{
			name:     "fixed string list",
			parent:   "Parent",
			field:    &model.Field{Name: "codes", IsRepeat: true, Attr: &model.FixedStringFieldAttribute{Length: 2}},
			expected: `&[_][]const u8{ "a", "a" }`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.testValue(tt.parent, tt.field)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestZigHasQuotes(t *testing.T) {
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
	g, _ := newZigTestModel()
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := g.HasQuotes(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
