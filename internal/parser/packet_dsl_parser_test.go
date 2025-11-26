package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// writeTempFile helper writes content to a temporary file and returns its path
func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "test-*.dsl")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	return f.Name()
}

func TestParseOption(t *testing.T) {
	dsl := `options {
    StringPrefixLenType = u16;
    ArrayPrefixLenType = u16;
    LittleEndian = true;
    JavaPackage = "com.finproto.sample.messages";
    GoPackage = "sample_bin";
    GoModule = "github.com/xinchentechnote/fin-proto-go/sample-bin/messages";
    FixedStringPadFromLeft = true;
    FixedStringPadChar = '0';
}`
	path := writeTempFile(t, dsl)
	defer os.Remove(path)

	data, err := ParseFile(path)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, &model.Configuration{
		StringLenPrefixLenType: "u16",
		ListLenPrefixLenType:   "u16",
		LittleEndian:           true,
		JavaPackage:            "com.finproto.sample.messages",
		GoPackage:              "sample_bin",
		GoModule:               "github.com/xinchentechnote/fin-proto-go/sample-bin/messages",
		Padding: &model.Padding{
			PadChar: "'0'",
			PadLeft: true,
		},
	}, data.(*model.BinaryModel).Config)
}

func TestParseSimplePacket(t *testing.T) {
	dsl := `
root packet Logon {
	char fchar,
	u8 fu8,
	i8 fi8,
	u16 fu16,
	i16 fi16,
	u32 fu32,
	i32 fi32,
	u64 fu64,
	i64 fi64,
	f32 ff32,
	f64 ff64,
	char[10] ffixedstr,
	string fstr,
}`
	path := writeTempFile(t, dsl)
	defer os.Remove(path)

	data, err := ParseFile(path)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	binModel := data.(*model.BinaryModel)
	assert.Equal(t, 1, len(binModel.PacketsMap))
	pkt := binModel.PacketsMap["Logon"]
	assert.Equal(t, "Logon", pkt.Name)
	assert.Equal(t, 1, len(binModel.Packets))
	assert.Equal(t, pkt, binModel.Packets[0])
	assert.True(t, pkt.IsRoot)
	expectedFields := []*model.Field{
		{Name: "fchar", IsRepeat: false, Line: 3, Column: 1, Attr: &model.BasicFieldAttribute{Type: "char"}},
		{Name: "fu8", IsRepeat: false, Line: 4, Column: 1, Attr: &model.BasicFieldAttribute{Type: "u8"}},
		{Name: "fi8", IsRepeat: false, Line: 5, Column: 1, Attr: &model.BasicFieldAttribute{Type: "i8"}},
		{Name: "fu16", IsRepeat: false, Line: 6, Column: 1, Attr: &model.BasicFieldAttribute{Type: "u16"}},
		{Name: "fi16", IsRepeat: false, Line: 7, Column: 1, Attr: &model.BasicFieldAttribute{Type: "i16"}},
		{Name: "fu32", IsRepeat: false, Line: 8, Column: 1, Attr: &model.BasicFieldAttribute{Type: "u32"}},
		{Name: "fi32", IsRepeat: false, Line: 9, Column: 1, Attr: &model.BasicFieldAttribute{Type: "i32"}},
		{Name: "fu64", IsRepeat: false, Line: 10, Column: 1, Attr: &model.BasicFieldAttribute{Type: "u64"}},
		{Name: "fi64", IsRepeat: false, Line: 11, Column: 1, Attr: &model.BasicFieldAttribute{Type: "i64"}},
		{Name: "ff32", IsRepeat: false, Line: 12, Column: 1, Attr: &model.BasicFieldAttribute{Type: "f32"}},
		{Name: "ff64", IsRepeat: false, Line: 13, Column: 1, Attr: &model.BasicFieldAttribute{Type: "f64"}},
		{Name: "ffixedstr", IsRepeat: false, Line: 14, Column: 1, Attr: &model.FixedStringFieldAttribute{Length: 10}},
		{Name: "fstr", IsRepeat: false, Line: 15, Column: 1, Attr: &model.DynamicStringFieldAttribute{}},
	}
	assert.Equal(t, 13, len(pkt.Fields))
	assert.Equal(t, expectedFields, pkt.Fields)
}

func TestParseMatchField(t *testing.T) {
	// DSL contains nested object and match field
	dsl := `
packet Parent {
	u16 MsgType,
	match MsgType as Body {
		0:OptionA,
		1:OptionB,
	},
}`
	path := writeTempFile(t, dsl)
	defer os.Remove(path)

	data, err := ParseFile(path)
	assert.NoError(t, err)
	packets := data.(*model.BinaryModel).PacketsMap
	pkt := packets["Parent"]
	assert.Equal(t, 2, len(pkt.Fields))

	msgTypeField := &model.Field{
		Name: "MsgType", IsRepeat: false, Line: 3, Column: 1, Attr: &model.BasicFieldAttribute{Type: "u16"},
	}
	expectedFields := []*model.Field{
		msgTypeField,
		{Name: "Body", IsRepeat: false, Line: 0, Column: 0, Attr: &model.MatchFieldAttribute{
			MatchKeyField: msgTypeField,
			MatchPairs: []model.MatchPair{
				{Key: "0", Value: "OptionA", Line: 5, Column: 1},
				{Key: "1", Value: "OptionB", Line: 6, Column: 1},
			},
		}},
	}
	assert.Equal(t, expectedFields, pkt.Fields)
}

func TestParseNested(t *testing.T) {
	dsl := `
packet Parent {
	NestedChild {
		i32 ChildField,
	},
}`
	path := writeTempFile(t, dsl)
	defer os.Remove(path)

	data, err := ParseFile(path)
	assert.NoError(t, err)
	packets := data.(*model.BinaryModel).PacketsMap
	pkt := packets["Parent"]
	assert.Equal(t, 1, len(pkt.Fields))

	nestedPacket := &model.Packet{
		Name: "NestedChild",
		Fields: []*model.Field{
			{Name: "ChildField", IsRepeat: false, Line: 4, Column: 1, Attr: &model.BasicFieldAttribute{Type: "i32"}},
		},
		Line: 3, Column: 1,
	}
	expectedFields := []*model.Field{
		{Name: "NestedChild", IsRepeat: false, Line: 3, Column: 1, Attr: &model.ObjectFieldAttribute{IsIner: true, PacketName: "NestedChild", RefPacket: nestedPacket}},
	}
	assert.Equal(t, expectedFields, pkt.Fields)
}

func TestLenghtFeildAttribute(t *testing.T) {
	dsl := `packet Body {}
root packet RootPacket {
	u16 MsgType,
	@lengthOf(Body)
	u16 BodyLen,
	Body,
	@calculatedFrom("CRC32")
	u32 Checksum,
}`
	path := writeTempFile(t, dsl)
	defer os.Remove(path)

	result, err := ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile returned error: %v", err)
	}
	assert.NotNil(t, result)
	binModel := result.(*model.BinaryModel)
	body := &model.Packet{Name: "Body", Column: 1, Line: 1, FieldMap: make(map[string]*model.Field), MatchFields: make(map[string][]model.MatchPair)}
	bodyField := &model.Field{Name: "Body", IsRepeat: false, Line: 6, Column: 1, Attr: &model.ObjectFieldAttribute{PacketName: "Body", RefPacket: body}}
	bodyLen := &model.Field{Name: "BodyLen", IsRepeat: false, Line: 5, Column: 1, Attr: &model.LengthFieldAttribute{LengthType: "u16", TragetField: bodyField}}
	bodyLen.LenAttr = &model.LengthOfAttribute{LengthField: bodyLen}
	bodyField.LenAttr = bodyLen.Attr
	expectedFields := []*model.Field{
		{Name: "MsgType", IsRepeat: false, Line: 3, Column: 1, Attr: &model.BasicFieldAttribute{Type: "u16"}},
		bodyLen,
		bodyField,
		{Name: "Checksum", IsRepeat: false, Line: 8, Column: 1, Attr: &model.CheckSumFieldAttribute{CheckSumType: "\"CRC32\"", Type: "u32"}},
	}
	assert.Equal(t, expectedFields, binModel.PacketsMap["RootPacket"].Fields)
}
