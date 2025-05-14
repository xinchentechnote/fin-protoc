package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
)

func ReadFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func NewPacketDslParserByFile(filename string) (*gen.PacketDslParser, error) {
	data, err := ReadFileToString(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	return NewPacketDslParserByContent(data)
}

func NewPacketDslParserByContent(data string) (*gen.PacketDslParser, error) {
	// 创建一个新的输入流
	input := antlr.NewInputStream(string(data))

	// 创建一个新的词法分析器
	lexer := gen.NewPacketDslLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 创建一个新的语法分析器
	parser := gen.NewPacketDslParser(stream)
	return parser, nil
}
