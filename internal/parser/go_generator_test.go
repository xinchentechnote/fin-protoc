package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// removeTempFile is the cleanup counterpart of writeTempFile.
func removeTempFile(t *testing.T, path string) {
	t.Helper()
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		t.Logf("failed to remove temp file %s: %v", path, err)
	}
}

// parseDSL parses dsl content via a temp file and returns the built model.
func parseDSL(t *testing.T, dsl string) *model.BinaryModel {
	t.Helper()
	path := writeTempFile(t, dsl)
	defer removeTempFile(t, path)

	result, err := ParseFile(path)
	require.NoError(t, err)
	require.Empty(t, result.(*model.BinaryModel).SyntaxErrors)
	return result.(*model.BinaryModel)
}

func TestGoGeneratorGenerate(t *testing.T) {
	dsl := `options {
	GoPackage = "messages";
	GoModule = "github.com/example/messages";
}
packet Logon {
	u32 ClientId ` + "`client id`" + `,
}
root packet RootPacket {
	u16 MsgType,
	match MsgType as Body {
		1: Logon,
	},
}`
	binModel := parseDSL(t, dsl)

	output, err := NewGoGenerator(binModel).Generate(binModel)
	require.NoError(t, err)

	assert.Contains(t, output, "root_packet.go")
	assert.Contains(t, output, "root_packet_test.go")
	assert.Contains(t, output, "logon.go")
	assert.Contains(t, output, "logon_test.go")

	rootCode := string(output["root_packet.go"])
	assert.Contains(t, rootCode, "package messages")
	assert.Contains(t, rootCode, "type RootPacket struct")
	// match field generates the message factory keyed by the match key field
	assert.Contains(t, rootCode, "func NewRootPacketMessageByMsgType(key uint16) (codec.BinaryCodec, error)")
	assert.Contains(t, rootCode, "func (p *RootPacket) Encode(buf *bytes.Buffer) error")
	assert.Contains(t, rootCode, "func (p *RootPacket) Decode(buf *bytes.Buffer) error")
}

// TestGoGeneratorLengthFieldBackfill guards the length backfill slice width:
// it must match the declared length field type size, not a hardcoded value.
func TestGoGeneratorLengthFieldBackfill(t *testing.T) {
	tests := []struct {
		name          string
		dsl           string
		expectedLine  string
		outdatedLine  string
	}{
		{
			name: "big endian u16 backfills 2 bytes",
			dsl: `packet Body {}
root packet RootPacket {
	@lengthOf(Body)
	u16 BodyLen,
	Body,
}`,
			expectedLine: "binary.BigEndian.PutUint16(buf.Bytes()[bodyPos:bodyPos + 2], p.BodyLen)",
			outdatedLine: "binary.BigEndian.PutUint16(buf.Bytes()[bodyPos:bodyPos + 4], p.BodyLen)",
		},
		{
			name: "little endian u32 backfills 4 bytes",
			dsl: `options {
	LittleEndian = true;
}
packet Body {}
root packet RootPacket {
	@lengthOf(Body)
	u32 BodyLen,
	Body,
}`,
			expectedLine: "binary.LittleEndian.PutUint32(buf.Bytes()[bodyPos:bodyPos + 4], p.BodyLen)",
			outdatedLine: "binary.LittleEndian.PutUint32(buf.Bytes()[bodyPos:bodyPos + 2], p.BodyLen)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binModel := parseDSL(t, tt.dsl)

			output, err := NewGoGenerator(binModel).Generate(binModel)
			require.NoError(t, err)
			code := string(output["root_packet.go"])
			assert.Contains(t, code, tt.expectedLine)
			// a hardcoded slice width for the wrong type must never come back
			assert.NotContains(t, code, tt.outdatedLine)
		})
	}
}

func TestGoGeneratorGetFieldType(t *testing.T) {
	child := &model.Packet{Name: "Child"}
	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{name: "basic", field: &model.Field{Attr: &model.BasicFieldAttribute{Type: "u32"}}, expected: "uint32"},
		{name: "length", field: &model.Field{Attr: &model.LengthFieldAttribute{LengthType: "u16"}}, expected: "uint16"},
		{name: "checksum", field: &model.Field{Attr: &model.CheckSumFieldAttribute{Type: "u32"}}, expected: "uint32"},
		{name: "fixed string", field: &model.Field{Attr: &model.FixedStringFieldAttribute{Length: 10}}, expected: "string"},
		{name: "dynamic string", field: &model.Field{Attr: &model.DynamicStringFieldAttribute{}}, expected: "string"},
		{name: "match", field: &model.Field{Attr: &model.MatchFieldAttribute{}}, expected: "codec.BinaryCodec"},
		{name: "object", field: &model.Field{Attr: &model.ObjectFieldAttribute{RefPacket: child}}, expected: "Child"},
	}
	g := NewGoGenerator(&model.BinaryModel{Config: &model.Configuration{}})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, g.getFieldType(tt.field))
		})
	}
}

func TestGoGeneratorGenerateStructCode(t *testing.T) {
	child := &model.Packet{Name: "Child"}
	p := &model.Packet{
		Name: "TestPacket",
		Fields: []*model.Field{
			{Name: "MsgType", Attr: &model.BasicFieldAttribute{Type: "u32"}},
			{Name: "Numbers", Attr: &model.BasicFieldAttribute{Type: "u32"}, IsRepeat: true},
			{Name: "Child", Attr: &model.ObjectFieldAttribute{RefPacket: child}},
		},
	}
	g := NewGoGenerator(&model.BinaryModel{Config: &model.Configuration{}})

	code := g.generateStructCode(p)

	assert.Contains(t, code, "type TestPacket struct {")
	assert.Contains(t, code, "MsgType uint32 `json:\"MsgType\"`")
	assert.Contains(t, code, "Numbers []uint32 `json:\"Numbers\"`")
	assert.Contains(t, code, "Child *Child `json:\"Child\"`")
	assert.Contains(t, code, "func NewTestPacket() *TestPacket")
}

func TestGoGeneratorGetOrder(t *testing.T) {
	be := NewGoGenerator(&model.BinaryModel{Config: &model.Configuration{LittleEndian: false}})
	assert.Equal(t, "", be.getOrder())

	le := NewGoGenerator(&model.BinaryModel{Config: &model.Configuration{LittleEndian: true}})
	assert.Equal(t, "LE", le.getOrder())
}

func TestGoGeneratorGetPadding(t *testing.T) {
	t.Run("no padding configured", func(t *testing.T) {
		g := NewGoGenerator(&model.BinaryModel{Config: &model.Configuration{}})
		assert.Nil(t, g.GetPadding(&model.Field{Attr: &model.BasicFieldAttribute{Type: "u32"}}))
	})

	t.Run("config level padding", func(t *testing.T) {
		g := NewGoGenerator(&model.BinaryModel{Config: &model.Configuration{
			Padding: &model.Padding{PadChar: "'0'", PadLeft: true},
		}})
		padding := g.GetPadding(&model.Field{Attr: &model.BasicFieldAttribute{Type: "u32"}})
		require.NotNil(t, padding)
		assert.Equal(t, "'0'", padding.PadChar)
		assert.True(t, padding.PadLeft)
	})

	t.Run("field level padding overrides config", func(t *testing.T) {
		g := NewGoGenerator(&model.BinaryModel{Config: &model.Configuration{
			Padding: &model.Padding{PadChar: "'0'", PadLeft: true},
		}})
		field := &model.Field{Attr: &model.FixedStringFieldAttribute{
			Length:  6,
			Padding: &model.Padding{PadChar: "'\x00'", PadLeft: false},
		}}
		padding := g.GetPadding(field)
		require.NotNil(t, padding)
		// nul padding is normalized to the escaped source form
		assert.Equal(t, "'\\x00'", padding.PadChar)
		assert.False(t, padding.PadLeft)
	})
}

func TestGoGeneratorTestValue(t *testing.T) {
	child := &model.Packet{Name: "Child"}
	tests := []struct {
		name     string
		field    *model.Field
		expected string
	}{
		{name: "u32", field: &model.Field{Name: "count", Attr: &model.BasicFieldAttribute{Type: "u32"}}, expected: "4"},
		{name: "u16", field: &model.Field{Name: "len", Attr: &model.BasicFieldAttribute{Type: "u16"}}, expected: "2"},
		{name: "fixed string padded to length", field: &model.Field{Name: "name", Attr: &model.FixedStringFieldAttribute{Length: 10}}, expected: `"xxxxxxxxxx"`},
		{name: "dynamic string", field: &model.Field{Name: "text", Attr: &model.DynamicStringFieldAttribute{}}, expected: `"hello"`},
		{name: "repeat u32", field: &model.Field{Name: "numbers", IsRepeat: true, Attr: &model.BasicFieldAttribute{Type: "u32"}}, expected: "[]uint32{4}"},
		{name: "repeat object", field: &model.Field{Name: "child", IsRepeat: true, Attr: &model.ObjectFieldAttribute{RefPacket: child}}, expected: "[]*msg.Child{child}"},
	}
	g := NewGoGenerator(&model.BinaryModel{PacketsMap: map[string]*model.Packet{"Child": child}, Config: &model.Configuration{}})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, g.generateTestValue(tt.field))
		})
	}
}
