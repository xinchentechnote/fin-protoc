package parser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

func TestFormatPacketDsl(t *testing.T) {
	dsl, _ := parser.ReadFileToString("testdata/need_format.dsl")
	expectedDsl, _ := parser.ReadFileToString("testdata/formatted.dsl")
	formattedDsl, err := parser.FormatPacketDsl(dsl)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, strings.TrimSpace(strings.ReplaceAll(expectedDsl, "\r\n", "\n")), strings.TrimSpace(formattedDsl))
}

func TestFormatPacketDsl1(t *testing.T) {
	formattedDsl, err := parser.FormatPacketDsl("root packet MyPacket {  Int64 name   `description` }")
	assert.NoError(t, err)
	fmt.Println(formattedDsl)
}

func TestVisitMetaDataDeclaration(t *testing.T) {
	cases := []struct {
		name     string
		input    string // metaDataDeclaration 规则对应的字符串
		expected string
	}{
		{
			name:     "只要类型+名字",
			input:    "Int32  count ",
			expected: "Int32 count",
		},
		{
			name:     "带 String 描述",
			input:    "string  fieldX   `some description`",
			expected: "string fieldX `some description`",
		},
		{
			name:     "带描述并且逗号",
			input:    "uInt16  headerLen  `length` , ",
			expected: "uInt16 headerLen `length`,",
		},
		{
			name:     "无类型（只 IDENTIFIER）",
			input:    "fieldOnly  `desc` , ",
			expected: "fieldOnly `desc`,",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := parser.NewPacketDslParserByContent(tc.input)
			if err != nil {
				t.Fatalf("构造 MetaDataDeclarationContext 失败: %v", err)
			}

			formatter := parser.NewPacketDslFormattor()
			got := formatter.VisitMetaDataDeclaration(p.MetaDataDeclaration().(*gen.MetaDataDeclarationContext)).(string)
			got = strings.TrimSpace(got)
			expected := strings.TrimSpace(tc.expected)
			if got != expected {
				t.Errorf("VisitMetaDataDeclaration 返回不符合预期。\n期望: %q\n实际: %q", expected, got)
			}
		})
	}
}

// ----- 测试 VisitMetaDataDefinition ----- //

func TestVisitMetaDataDefinition(t *testing.T) {
	input := `
MetaData SampleMeta {
    string   key ` + "`" + `a description` + "`" + ` ,
    Int32   value ` + "`" + `another desc` + "`" + `
}
`
	// 1. 构造上下文
	p, err := parser.NewPacketDslParserByContent(input)
	if err != nil {
		t.Fatalf("构造 MetaDataDefinitionContext 时出错: %v", err)
	}

	formatter := parser.NewPacketDslFormattor()
	got := formatter.VisitMetaDataDefinition(p.MetaDataDefinition().(*gen.MetaDataDefinitionContext)).(string)

	expected := `MetaData SampleMeta {
	string key ` + "`" + `a description` + "`" + `,
	Int32 value ` + "`" + `another desc` + "`" + `
}`

	// 4. 去掉首尾空行并比较
	got = strings.TrimSpace(got)
	expected = strings.TrimSpace(expected)
	if got != expected {
		t.Errorf("VisitMetaDataDefinition 输出不符合预期。\n期望:\n----\n%s\n----\n实际:\n----\n%s\n----", expected, got)
	}
}

func TestVisitMatchField(t *testing.T) {
	// 单个 matchPair
	input := `
match category  {
    "A" : ValA ,
    10 : Val10
}
`
	p, err := parser.NewPacketDslParserByContent(input)
	if err != nil {
		t.Fatalf("构造 MatchFieldContext 失败: %v", err)
	}
	formatter := parser.NewPacketDslFormattor()
	got := formatter.VisitMatchField(p.MatchField().(*gen.MatchFieldContext)).(string)

	expected := `match category {
	"A" : ValA,
	10 : Val10
}`

	got = strings.TrimSpace(got)
	expected = strings.TrimSpace(expected)
	if got != expected {
		t.Errorf("VisitMatchField 输出不符合预期。\n期望:\n----\n%s\n----\n实际:\n----\n%s\n----", expected, got)
	}
}
