package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// grammar_full.dsl exercises every PacketDsl grammar rule; the model
// assertions below pin the parse results, and every generator must produce
// fallback-marker-free output for it. The CI job `codegen-verify` additionally
// compiles and runs the generated code for all six languages.
const grammarFullDSL = "testdata/grammar_full.dsl"

func parseGrammarFull(t *testing.T) *model.BinaryModel {
	t.Helper()
	result, err := ParseFile(grammarFullDSL)
	require.NoError(t, err)
	binModel := result.(*model.BinaryModel)
	require.Empty(t, binModel.SyntaxErrors)
	return binModel
}

func TestGrammarFullParseOptions(t *testing.T) {
	m := parseGrammarFull(t)

	assert.Equal(t, &model.Configuration{
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u16",
		LittleEndian:           false,
		JavaPackage:            "com.grammar.full.messages",
		GoPackage:              "messages",
		GoModule:               "fin-protoc-codegen-verify",
		Padding:                &model.Padding{PadChar: "' '", PadLeft: false},
	}, m.Config)
}

func TestGrammarFullParseMetaData(t *testing.T) {
	m := parseGrammarFull(t)

	// u64, char[12], zchar[8], string and one ref metadata
	assert.Len(t, m.MetaDataMap, 5)
	assert.Equal(t, &model.BasicFieldAttribute{Type: "u64"}, m.MetaDataMap["Price"].Attr)
	assert.Equal(t, &model.FixedStringFieldAttribute{Length: 12}, m.MetaDataMap["SecurityID"].Attr)
	assert.Equal(t, &model.FixedStringFieldAttribute{Length: 8,
		Padding: &model.Padding{PadChar: "'\x00'", PadLeft: false}}, m.MetaDataMap["ZCode"].Attr)
	assert.Equal(t, &model.DynamicStringFieldAttribute{}, m.MetaDataMap["Note"].Attr)
	// the ref metadata aliases the earlier entry's type
	assert.Equal(t, &model.BasicFieldAttribute{Type: "u64"}, m.MetaDataMap["RefPrice"].Attr)
}

func TestGrammarFullParsePackets(t *testing.T) {
	m := parseGrammarFull(t)

	assert.Len(t, m.Packets, 7) // EmptyMsg, Logon, StrA, StrB, Item, InlineItem, GrammarFull
	require.NotNil(t, m.RootPacket)
	assert.Equal(t, "GrammarFull", m.RootPacket.Name)

	root := m.RootPacket
	byName := root.FieldMap

	// every declared field made it into the packet
	for _, name := range []string{
		"Side", "F_u8", "F_i8", "MsgType", "F_i16", "F_u32", "F_i32", "F_u64", "F_i64", "F_f32", "F_f64",
		"A_u8", "A_i8", "A_u16", "A_i16", "A_u32", "A_i32", "A_u64", "A_i64", "A_f32", "A_f64",
		"CharArr", "Note", "MetaNote", "RefPriceField", "Tagged",
		"PadSpace", "PadNul", "PadDefaultLeft",
		"ItemLen", "ItemBlock",
		"RepeatU32", "RepeatStr", "RepeatItems", "InnerBlk", "InlineObj",
		"NamedObj", "BareEmpty", "Body", "SessKind", "Sess", "Checksum",
	} {
		assert.Contains(t, byName, name)
	}

	// alias type names normalize to canonical short names
	assert.Equal(t, "u32", byName["A_u32"].GetType())
	assert.Equal(t, "f64", byName["A_f64"].GetType())
	// char[] is a dynamic string like string
	assert.Equal(t, "string", byName["CharArr"].GetType())
	assert.Equal(t, "string", byName["MetaNote"].GetType())
	// metadata-typed fields resolve to their metadata type
	assert.Equal(t, "u64", byName["RefPriceField"].GetType())

	// @tag
	assert.Equal(t, 9, byName["Tagged"].Tag)

	// padding attribute variants
	assert.Equal(t, &model.Padding{PadChar: "' '", PadLeft: false}, byName["PadSpace"].Attr.(*model.FixedStringFieldAttribute).Padding)
	assert.Equal(t, &model.Padding{PadChar: "'\x00'", PadLeft: false}, byName["PadNul"].Attr.(*model.FixedStringFieldAttribute).Padding)
	assert.Equal(t, &model.Padding{PadChar: "' '", PadLeft: true}, byName["PadDefaultLeft"].Attr.(*model.FixedStringFieldAttribute).Padding)
	assert.Equal(t, &model.Padding{PadChar: "'0'", PadLeft: true},
		m.PacketsMap["Logon"].FieldMap["UserName"].Attr.(*model.FixedStringFieldAttribute).Padding)

	// length field + its target
	lengthAttr, ok := byName["ItemLen"].Attr.(*model.LengthFieldAttribute)
	require.True(t, ok)
	assert.Equal(t, "ItemBlock", lengthAttr.TargetField.Name)
	assert.Equal(t, "u16", lengthAttr.LengthType)

	// object references resolve to their packets (explicit name, bare and inline)
	assert.Equal(t, "Item", byName["ItemBlock"].Attr.(*model.ObjectFieldAttribute).RefPacket.Name)
	assert.Equal(t, "InlineItem", byName["NamedObj"].Attr.(*model.ObjectFieldAttribute).RefPacket.Name)
	assert.Equal(t, "EmptyMsg", byName["BareEmpty"].Attr.(*model.ObjectFieldAttribute).RefPacket.Name)
	assert.True(t, byName["InlineObj"].Attr.(*model.ObjectFieldAttribute).IsIner)
	assert.Equal(t, 2, len(byName["InlineObj"].Attr.(*model.ObjectFieldAttribute).RefPacket.Fields))

	// repeat markers
	assert.True(t, byName["RepeatU32"].IsRepeat)
	assert.True(t, byName["RepeatItems"].IsRepeat)
	assert.True(t, byName["InnerBlk"].IsRepeat)
	assert.False(t, byName["InlineObj"].IsRepeat)

	// match fields: digits key with digit list, string key with string list
	body, ok := byName["Body"].Attr.(*model.MatchFieldAttribute)
	require.True(t, ok)
	assert.Equal(t, []model.MatchPair{
		{Key: "1", Value: "Logon", Line: 126},
		{Key: "2", Value: "EmptyMsg", Line: 127},
		{Key: "3", Value: "EmptyMsg", Line: 127},
	}, body.MatchPairs)
	sess, ok := byName["Sess"].Attr.(*model.MatchFieldAttribute)
	require.True(t, ok)
	assert.Equal(t, []model.MatchPair{
		{Key: `"fast"`, Value: "StrA", Line: 132},
		{Key: `"slow"`, Value: "StrB", Line: 133},
		{Key: `"idle"`, Value: "StrB", Line: 133},
	}, sess.MatchPairs)

	// checksum field
	checksum, ok := byName["Checksum"].Attr.(*model.CheckSumFieldAttribute)
	require.True(t, ok)
	assert.Equal(t, `"CRC32"`, checksum.CheckSumType)
	assert.Equal(t, "u32", checksum.Type)
}

// TestGrammarFullGenerators runs every generator over the fixture and fails
// on fallback markers like "-- ... not supported --" that would otherwise
// silently produce uncompilable output.
func TestGrammarFullGenerators(t *testing.T) {
	m := parseGrammarFull(t)

	generators := []struct {
		name     string
		generate func(*model.BinaryModel) (map[string][]byte, error)
	}{
		{"Go", NewGoGenerator(m).Generate},
		{"Java", NewJavaGenerator(m).Generate},
		{"Rust", NewRustGenerator(m).Generate},
		{"Lua", NewLuaWspGenerator(m).Generate},
		{"Python", NewPythonGenerator(m).Generate},
		{"C++", NewCppGenerator(m).Generate},
		{"Zig", NewZigGenerator(m).Generate},
	}

	// fallback markers emitted by the generators for unsupported constructs;
	// deliberately specific so legitimate error messages like Go's
	// "unknown message type" or Java's "Unsupported MsgType:" don't match
	fallbackMarkers := []string{
		"unknow type", "unknown type", "not supported",
		"unsupported type", "unsupport type", "unsupported numeric",
	}
	for _, g := range generators {
		t.Run(g.name, func(t *testing.T) {
			output, err := g.generate(m)
			require.NoError(t, err)
			require.NotEmpty(t, output)

			for name, code := range output {
				assert.NotEmpty(t, code, "%s: %s is empty", g.name, name)
				lower := strings.ToLower(string(code))
				for _, marker := range fallbackMarkers {
					assert.NotContains(t, lower, marker,
						"%s: %s contains fallback marker %q", g.name, name, marker)
				}
			}
		})
	}
}
