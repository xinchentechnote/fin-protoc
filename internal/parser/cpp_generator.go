package parser

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// CppGenerator a go code generator
type CppGenerator struct {
	config   *GeneratorConfig
	binModel *model.BinaryModel
}

// NewCppGenerator new
func NewCppGenerator(config *GeneratorConfig, binModel *model.BinaryModel) *CppGenerator {
	return &CppGenerator{
		config:   config,
		binModel: binModel,
	}
}

// Generate go code
func (g CppGenerator) Generate(binModel *model.BinaryModel) (map[string][]byte, error) {
	output := make(map[string][]byte)
	for _, pkt := range binModel.PacketsMap {
		code := g.generateCodeForPacket(&pkt)
		output[strcase.ToSnake(pkt.Name)+".cpp"] = []byte(code)
	}
	return output, nil
}

func (g CppGenerator) generateCodeForPacket(_ *model.Packet) string {
	var b strings.Builder

	b.WriteString("// Code generated by fin-protoc. DO NOT EDIT.\n")

	return b.String()
}
