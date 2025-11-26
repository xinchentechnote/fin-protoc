package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
)

func TestFormatPacketDsl(t *testing.T) {
	dsl, _ := ReadFileToString("testdata/sample_binary.dsl")
	expectedDsl, _ := ReadFileToString("testdata/sample_binary_formatted.dsl")
	formattedDsl, err := FormatPacketDsl(dsl)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, strings.TrimSpace(strings.ReplaceAll(expectedDsl, "\r\n", "\n")), strings.TrimSpace(formattedDsl))
}

func TestFormatPacketDsl1(t *testing.T) {
	formattedDsl, err := FormatPacketDsl("root  packet    MyPacket  {  int64  name   `description` , }")
	assert.NoError(t, err)
	assert.Equal(t, formattedDsl, "root packet MyPacket {\n    int64 name `description`,\n}")
}

func TestVisitMetaDataDefinition(t *testing.T) {
	input := `
MetaData SampleMeta {
    string   key  ` + "`" + `a description` + "`" + ` ,
    int32   value  ` + "`" + `another desc` + "`" + `
}
`
	p, stream, err := NewPacketDslParserByContent(input)
	assert.NoError(t, err)
	formatter := NewPacketDslFormattor(stream)
	got := formatter.VisitMetaDataDefinition(p.MetaDataDefinition().(*gen.MetaDataDefinitionContext)).(string)

	expected := `MetaData SampleMeta {
    string key ` + "`" + `a description` + "`" + `,
    int32 value ` + "`" + `another desc` + "`" + `,
}`

	assert.Equal(t, got, expected)
}

func TestVisitMatchField(t *testing.T) {
	input := `
match    bodyType   as  body  {
     "A" : ValA,
     10 : Val10
}
`
	p, stream, err := NewPacketDslParserByContent(input)
	assert.NoError(t, err)
	formatter := NewPacketDslFormattor(stream)
	got := formatter.VisitMatchFieldDeclaration(p.MatchFieldDeclaration().(*gen.MatchFieldDeclarationContext)).(string)

	expected := `match bodyType as body {
    "A" : ValA,
    10 : Val10,
}`

	got = strings.TrimSpace(got)
	expected = strings.TrimSpace(expected)
	assert.Equal(t, got, expected)
}

func TestFormatStringList(t *testing.T) {
	values := []string{"1", "2", "3", "4", "5", "6", "7"}
	got := formatStringList(values, 3)

	expected := `[
    1, 2, 3,
    4, 5, 6,
    7
]`

	assert.Equal(t, got, expected)
}
