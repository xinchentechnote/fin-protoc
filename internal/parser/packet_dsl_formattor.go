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

	// 判断是否有 root 修饰符
	if ctx.ROOT() != nil {
		formattedDsl.WriteString("root ")
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
	typeName := ""
	if ctx.Type_() != nil {
		typeName = ctx.Type_().GetText()
	}
	fieldName := ctx.IDENTIFIER().GetText()
	description := ""
	if ctx.STRING_LITERAL() != nil {
		description = ctx.STRING_LITERAL().GetText()
	}
	formattedDsl := fmt.Sprintf("%s %s %s", typeName, fieldName, description)
	if ctx.COMMA() != nil {
		formattedDsl = strings.TrimSpace(formattedDsl) + ","
	}
	return formattedDsl
}

// VisitType for visiting type definitions.
func (v *PacketDslFormattor) VisitType(ctx *gen.TypeContext) interface{} {
	fmt.Println("Visiting Type")
	return nil
}

// VisitMatchField for visiting match fields.
func (v *PacketDslFormattor) VisitMatchField(ctx *gen.MatchFieldContext) interface{} {
	fieldName := ctx.IDENTIFIER().GetText()
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("match %s {\n", fieldName))
	for _, pairCtx := range ctx.AllMatchPair() {
		key := ""
		// 处理 STRING / number / list
		switch {
		case pairCtx.STRING() != nil:
			key = pairCtx.STRING().GetText()
		case pairCtx.DIGITS() != nil:
			key = pairCtx.DIGITS().GetText()
		case pairCtx.List() != nil:
			items := []string{}
			for _, num := range pairCtx.List().AllDIGITS() {
				items = append(items, num.GetText())
			}
			for _, num := range pairCtx.List().AllSTRING() {
				items = append(items, num.GetText())
			}
			key = fmt.Sprintf("[%s]", strings.Join(items, ", "))
		}

		value := pairCtx.IDENTIFIER().GetText()
		if pairCtx.COMMA() != nil {
			value = strings.TrimSpace(value) + ","
		}
		sb.WriteString(fmt.Sprintf("\t%s : %s\n", key, value))
	}
	sb.WriteString("}")
	return sb.String()
}
