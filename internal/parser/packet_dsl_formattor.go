package parser

import (
	"fmt"
	"strings"

	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
)

// FormatPacketDsl formats the packet DSL string.
func FormatPacketDsl(dsl string) (string, error) {
	// 创建一个新的语法分析器
	parser, err := NewPacketDslParserByContent(dsl)
	if err != nil {
		return "", fmt.Errorf("could not create parser: %v", err)
	}
	// 设置语法规则，开始解析
	tree := parser.Packet()
	formattor := NewPacketDslFormattor()
	formattedDsl := tree.Accept(formattor).(string)
	return strings.TrimSpace(formattedDsl), nil
}

// NewPacketDslFormattor creates a new instance of PacketDslFormattor.
func NewPacketDslFormattor() *PacketDslFormattor {
	return &PacketDslFormattor{
		BasePacketDslVisitor: &gen.BasePacketDslVisitor{},
	}
}

// PacketDslFormattor is a visitor for formatting packet DSL.
type PacketDslFormattor struct {
	*gen.BasePacketDslVisitor
}

// VisitPacket overrides the default implementation for protocol definitions.
func (v *PacketDslFormattor) VisitPacket(ctx *gen.PacketContext) interface{} {
	formattedDsl := ""
	childrens := ctx.GetChildren()
	if len(childrens) != 0 {
		for _, children := range childrens {
			switch c := children.(type) {
			case *gen.PacketDefinitionContext:
				formattedDsl += v.VisitPacketDefinition(c).(string)
			case *gen.MetaDataDefinitionContext:
				formattedDsl += v.VisitMetaDataDefinition(c).(string)
			default:
				formattedDsl += "error"
			}
			formattedDsl += "\n\n"
		}
	}
	return formattedDsl
}

// VisitPacketDefinition overrides the default implementation for packet definitions.
func (v *PacketDslFormattor) VisitPacketDefinition(ctx *gen.PacketDefinitionContext) interface{} {
	var formattedDsl strings.Builder

	// 判断是否有 ROOT 修饰符
	if ctx.ROOT() != nil {
		formattedDsl.WriteString("ROOT ")
	}

	// packet 关键字 + 包名
	formattedDsl.WriteString("packet ")
	formattedDsl.WriteString(ctx.IDENTIFIER().GetText())
	formattedDsl.WriteString(" {\n")

	// 遍历字段定义（通过 fieldDefinition*）
	for _, fieldCtx := range ctx.AllFieldDefinition() {
		formatted := v.VisitFieldDefinition(fieldCtx.(*gen.FieldDefinitionContext)).(string)
		formattedDsl.WriteString("    " + formatted + "\n")
	}
	formattedDsl.WriteString("}")
	return formattedDsl.String()
}

// VisitFieldDefinition overrides the default implementation for field definitions.
func (v *PacketDslFormattor) VisitFieldDefinition(ctx *gen.FieldDefinitionContext) interface{} {
	switch {
	case ctx.IDENTIFIER() != nil:
		return ctx.IDENTIFIER().GetText()

	case ctx.MetaDataDeclaration() != nil:
		return v.VisitMetaDataDeclaration(ctx.MetaDataDeclaration().(*gen.MetaDataDeclarationContext)).(string)

	case ctx.MatchField() != nil:
		return v.VisitMatchField(ctx.MatchField().(*gen.MatchFieldContext)).(string)
	default:
		return ""
	}
}

// VisitMetaDataDefinition overrides the default implementation for metadata definitions.
func (v *PacketDslFormattor) VisitMetaDataDefinition(ctx *gen.MetaDataDefinitionContext) interface{} {
	var sb strings.Builder

	// 获取 MetaData 名称
	metaName := ctx.IDENTIFIER().GetText()
	sb.WriteString(fmt.Sprintf("MetaData %s {\n", metaName))

	// 遍历所有 metaDataDeclaration 子节点
	for _, decl := range ctx.AllMetaDataDeclaration() {
		if d, ok := decl.(*gen.MetaDataDeclarationContext); ok {
			result := v.VisitMetaDataDeclaration(d).(string)
			sb.WriteString("\t" + result + "\n")
		}
	}

	sb.WriteString("}")
	return sb.String()
}

// VisitMetaDataDeclaration for visiting metadata declarations.
func (v *PacketDslFormattor) VisitMetaDataDeclaration(ctx *gen.MetaDataDeclarationContext) interface{} {
	typeName := ctx.Type_().GetText()
	fieldName := ctx.IDENTIFIER().GetText()
	description := ctx.STRING_LITERAL().GetText()
	description = strings.Trim(description, "`") // 去除反引号

	return fmt.Sprintf("%s %s `%s`", typeName, fieldName, description)
}

// VisitType for visiting type definitions.
func (v *PacketDslFormattor) VisitType(ctx *gen.TypeContext) interface{} {
	fmt.Println("Visiting Type")
	return nil
}

// VisitMatchField for visiting match fields.
func (v *PacketDslFormattor) VisitMatchField(ctx *gen.MatchFieldContext) interface{} {
	fmt.Println("Visiting MatchField")
	return nil
}
