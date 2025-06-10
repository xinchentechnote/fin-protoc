package parser

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
)

// AddIndent4ln adds 4-space indent and a newline (similar to fmt.Println)
func AddIndent4ln(s string) string {
	return AddIndent4(s) + "\n"
}

// AddIndent4 add 4 Indent bydefault
func AddIndent4(s string) string {
	return AddIndent(s, 4)
}

// AddIndent add Indent
func AddIndent(s string, spaces int) string {
	indent := strings.Repeat(" ", spaces)
	return indent + strings.ReplaceAll(s, "\n", "\n"+indent)
}

// ReadFileToString reads the content of a file and returns it as a string.
func ReadFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NewPacketDslParserByFile creates a new PacketDslParser instance from a file.
func NewPacketDslParserByFile(filename string) (*gen.PacketDslParser, *antlr.CommonTokenStream, error) {
	data, err := ReadFileToString(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read file: %v", err)
	}
	return NewPacketDslParserByContent(data)
}

// NewPacketDslParserByContent creates a new PacketDslParser instance from a string.
func NewPacketDslParserByContent(data string) (*gen.PacketDslParser, *antlr.CommonTokenStream, error) {
	// 创建一个新的输入流
	input := antlr.NewInputStream(string(data))

	// 创建一个新的词法分析器
	lexer := gen.NewPacketDslLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 创建一个新的语法分析器
	parser := gen.NewPacketDslParser(stream)
	return parser, stream, nil
}

// ParseCharArrayType parse char[\d]
func ParseCharArrayType(fieldType string) (size string, ok bool) {
	pattern := `^char\[(\d+)\]$`
	matches := regexp.MustCompile(pattern).FindStringSubmatch(fieldType)
	if len(matches) == 2 {
		return matches[1], true
	}
	return "", false
}

// RenderToString render tmpl
func RenderToString(tmpl string, lang string, data interface{}) (string, error) {
	t := template.Must(template.New(lang).Parse(tmpl))

	var buf bytes.Buffer
	err := t.Execute(&buf, data) // Directly pass data instead of wrapping it
	if err != nil {
		return "", fmt.Errorf("template execution failed: %w", err)
	}

	return buf.String(), nil
}

// WriteCodeToFile write code to file
func WriteCodeToFile(path string, codeMap map[string][]byte) {
	for name, datas := range codeMap {
		filepath := path + "/" + name
		f, err := os.Create(filepath)
		if nil != err {
			fmt.Println("create file fail.")
		}
		defer f.Close()
		_, err = f.Write(datas)
		if nil != err {
			fmt.Println("write file fail.")
		}
		fmt.Println("Generated code for packet:", filepath)
	}
}
