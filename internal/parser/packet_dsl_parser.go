package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
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
	lexer := gen.NewPacketDslLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 创建一个新的语法分析器
	parser := gen.NewPacketDslParser(stream)

	// 设置语法规则，开始解析
	tree := parser.Packet()
	fmt.Println("Parsing completed, tree generated:", tree)
	visitor := NewPacketDslVisitor()
	astData := tree.Accept(visitor)
	return astData, nil
}

// PacketDslVisitorImpl 自定义访问器，解析具体的协议数据
type PacketDslVisitorImpl struct {
	*gen.BasePacketDslVisitor
}

// NewPacketDslVisitor 创建一个新的自定义访问器
func NewPacketDslVisitor() *PacketDslVisitorImpl {
	return &PacketDslVisitorImpl{
		BasePacketDslVisitor: &gen.BasePacketDslVisitor{}, // Initialize it here if needed
	}
}

// VisitPacket overrides the default implementation for protocol definitions.
func (v *PacketDslVisitorImpl) VisitPacket(ctx *gen.PacketContext) interface{} {
	packetDefinitions := ctx.AllPacketDefinition()
	if len(packetDefinitions) != 0 {
		for _, packetDefinition := range packetDefinitions {
			packetDef := packetDefinition.Accept(v).(*gen.PacketDefinitionContext)
			fmt.Println("Packet Definition:", packetDef.GetText())
		}
	}
	fmt.Println("Visiting Proto")
	return nil
}

// VisitPacketDefinition overrides the default implementation for packet definitions.
func (v *PacketDslVisitorImpl) VisitPacketDefinition(ctx *gen.PacketDefinitionContext) interface{} {
	fieldDefinitions := ctx.AllFieldDefinition()
	if len(fieldDefinitions) != 0 {
		for _, fieldDefinition := range fieldDefinitions {
			fmt.Println("Field Definition:", fieldDefinition.GetText())
		}

	}
	fmt.Println("Visiting PacketDefinition")
	return ctx
}

// VisitFieldDefinition overrides the default implementation for field definitions.
func (v *PacketDslVisitorImpl) VisitFieldDefinition(ctx *gen.FieldDefinitionContext) interface{} {
	fmt.Println("Visiting FieldDefinition")
	return ctx
}

// VisitMetaDataDefinition overrides the default implementation for metadata definitions.
func (v *PacketDslVisitorImpl) VisitMetaDataDefinition(ctx *gen.MetaDataDefinitionContext) interface{} {
	fmt.Println("Visiting MetaDataDefinition")
	return nil
}

// VisitMetaDataDeclaration for visiting metadata declarations.
func (v *PacketDslVisitorImpl) VisitMetaDataDeclaration(ctx *gen.MetaDataDeclarationContext) interface{} {
	fmt.Println("Visiting MetaDataDeclaration")
	return nil
}

// VisitType for visiting type definitions.
func (v *PacketDslVisitorImpl) VisitType(ctx *gen.TypeContext) interface{} {
	fmt.Println("Visiting Type")
	return nil
}

// VisitDescription for visiting description nodes.
func (v *PacketDslVisitorImpl) VisitDescription(ctx *gen.DescriptionContext) interface{} {
	fmt.Println("Visiting Description")
	return nil
}

// VisitMatchField for visiting match fields.
func (v *PacketDslVisitorImpl) VisitMatchField(ctx *gen.MatchFieldContext) interface{} {
	fmt.Println("Visiting MatchField")
	return nil
}
