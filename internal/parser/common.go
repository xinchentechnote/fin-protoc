package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
)

// ReadFileToString reads the content of a file and returns it as a string.
func ReadFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NewPacketDslParserByFile creates a new PacketDslParser instance from a file.
func NewPacketDslParserByFile(filename string) (*gen.PacketDslParser, error) {
	data, err := ReadFileToString(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	return NewPacketDslParserByContent(data)
}

// NewPacketDslParserByContent creates a new PacketDslParser instance from a string.
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
