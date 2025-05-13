package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

// ParseFile 解析协议文件并返回协议定义的 AST
func ParseFile(filename string) (interface{}, error) {
	// 读取协议文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	// 创建一个新的输入流
	input := antlr.NewInputStream(string(data))

	// 创建一个新的词法分析器
	lexer := NewBinaryPacketLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 创建一个新的语法分析器
	parser := NewBinaryPacketParser(stream)

	// 设置语法规则，开始解析
	tree := parser.PacketDefinition()
	fmt.Println("Parsing completed, tree generated:", tree)
	visitor := NewCustomVisitor()
	astData := visitor.Visit(tree)

	return astData, nil
}

// CustomVisitor 自定义访问器，解析具体的协议数据
type CustomVisitor struct {
	*BaseBinaryPacketVisitor
}

// NewCustomVisitor 创建一个新的自定义访问器
func NewCustomVisitor() *CustomVisitor {
	return &CustomVisitor{
		BaseBinaryPacketVisitor: &BaseBinaryPacketVisitor{}, // Initialize it here if needed
	}
}

// VisitPacketDefinition overrides the default implementation for packet definitions.
func (v *CustomVisitor) VisitPacketDefinition(ctx *PacketDefinitionContext) interface{} {
	fmt.Println("Visiting PacketDefinition")
	return nil
}

// VisitFieldDefinition overrides the default implementation for field definitions.
func (v *CustomVisitor) VisitFieldDefinition(ctx *FieldDefinitionContext) interface{} {
	fmt.Println("Visiting FieldDefinition")
	return nil
}

// VisitMetaDataDefinition overrides the default implementation for metadata definitions.
func (v *CustomVisitor) VisitMetaDataDefinition(ctx *MetaDataDefinitionContext) interface{} {
	fmt.Println("Visiting MetaDataDefinition")
	return nil
}

// VisitMetaDataDeclaration for visiting metadata declarations.
func (v *CustomVisitor) VisitMetaDataDeclaration(ctx *MetaDataDeclarationContext) interface{} {
	fmt.Println("Visiting MetaDataDeclaration")
	return nil
}

// VisitType for visiting type definitions.
func (v *CustomVisitor) VisitType(ctx *TypeContext) interface{} {
	fmt.Println("Visiting Type")
	return nil
}

// VisitDescription for visiting description nodes.
func (v *CustomVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	fmt.Println("Visiting Description")
	return nil
}

// VisitMatchField for visiting match fields.
func (v *CustomVisitor) VisitMatchField(ctx *MatchFieldContext) interface{} {
	fmt.Println("Visiting MatchField")
	return nil
}
