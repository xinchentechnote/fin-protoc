package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
)

// FormatPacketDsl formats the packet DSL string.
func FormatPacketDsl(dsl string) (string, error) {
	parser, stream, err := NewPacketDslParserByContent(dsl)
	if err != nil {
		return "", fmt.Errorf("could not create parser: %v", err)
	}
	listener := NewSyntaxErrorListener()
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	// parese the file
	tree := parser.Packet()
	if listener.HasErrors() {
		return dsl, fmt.Errorf("syntax errors found: %v", listener.Errors)
	}
	formattor := NewPacketDslFormattor(stream)
	formattedDsl := tree.Accept(formattor).(string)
	return strings.TrimSpace(formattedDsl), nil
}

// NewPacketDslFormattor creates a new instance of PacketDslFormattor.
func NewPacketDslFormattor(tokenStream *antlr.CommonTokenStream) *PacketDslFormattor {
	return &PacketDslFormattor{
		BasePacketDslVisitor: &gen.BasePacketDslVisitor{},
		tokenStream:          tokenStream,
		lineComments:         make(map[antlr.Token]struct{}),
	}
}

// PacketDslFormattor is a visitor for formatting packet DSL.
type PacketDslFormattor struct {
	*gen.BasePacketDslVisitor
	tokenStream  *antlr.CommonTokenStream
	lineComments map[antlr.Token]struct{}
}

func (v *PacketDslFormattor) getHiddenLeft(token antlr.Token) string {
	hidden := v.tokenStream.GetHiddenTokensToLeft(token.GetTokenIndex(), antlr.TokenHiddenChannel)
	if hidden == nil {
		return ""
	}
	var sb strings.Builder
	for _, t := range hidden {
		if t.GetTokenType() == gen.PacketDslParserLINE_COMMENT {
			if _, ok := v.lineComments[t]; ok {
				continue
			}
			v.lineComments[t] = struct{}{}
			sb.WriteString(t.GetText())
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (v *PacketDslFormattor) getHiddenRightAtSameLine(token antlr.Token) string {
	hidden := v.tokenStream.GetHiddenTokensToRight(token.GetTokenIndex(), antlr.TokenHiddenChannel)
	if hidden == nil {
		return ""
	}
	var sb strings.Builder
	for _, t := range hidden {
		if t.GetTokenType() == gen.PacketDslParserLINE_COMMENT {
			if _, ok := v.lineComments[t]; ok {
				continue
			}
			if t.GetLine() != token.GetLine() {
				continue
			}
			v.lineComments[t] = struct{}{}
			sb.WriteString(t.GetText())
		}
	}
	return strings.TrimRight(sb.String(), "\n")
}

// VisitPacket overrides the default implementation for protocol definitions.
func (v *PacketDslFormattor) VisitPacket(ctx *gen.PacketContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	childrens := ctx.GetChildren()
	if len(childrens) != 0 {
		for idx, child := range childrens {
			switch c := child.(type) {
			case *gen.PacketDefinitionContext:
				formattedDsl.WriteString(v.VisitPacketDefinition(c).(string))
			case *gen.MetaDataDefinitionContext:
				formattedDsl.WriteString(v.VisitMetaDataDefinition(c).(string))
			case *gen.OptionDefinitionContext:
				formattedDsl.WriteString(v.VisitOptionDefinition(c).(string))
			default:
				if termNode, ok := child.(antlr.TerminalNode); ok {
					formattedDsl.WriteString(v.VisitTerminal(termNode).(string))
				} else {
					text := fmt.Sprint(child)
					formattedDsl.WriteString(text)
				}
			}
			if idx < len(childrens)-1 {
				formattedDsl.WriteString("\n\n")
			}
		}
	}
	formattedDsl.WriteString(v.getHiddenRightAtSameLine(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitPacketDefinition overrides the default implementation for packet definitions.
func (v *PacketDslFormattor) VisitPacketDefinition(ctx *gen.PacketDefinitionContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	if ctx.ROOT() != nil {
		formattedDsl.WriteString("root ")
	}

	// packet key world and name
	formattedDsl.WriteString("packet ")
	formattedDsl.WriteString(ctx.IDENTIFIER().GetText())
	formattedDsl.WriteString(" {\n")

	// iterate over all field definitions
	for _, fieldCtx := range ctx.AllFieldDefinitionWithAttribute() {
		if fc, ok := fieldCtx.(*gen.FieldDefinitionWithAttributeContext); ok {
			formatted := v.VisitFieldDefinitionWithAttribute(fc).(string)
			formattedDsl.WriteString(AddIndent4ln(formatted))
		}
	}

	formattedDsl.WriteString("}")
	formattedDsl.WriteString(v.getHiddenRightAtSameLine(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitFieldDefinitionWithAttribute overrides the default implementation for field definitions with attributes.
func (v *PacketDslFormattor) VisitFieldDefinitionWithAttribute(ctx *gen.FieldDefinitionWithAttributeContext) interface{} {
	var formattedDsl strings.Builder
	if len(ctx.AllFieldAttribute()) > 0 {
		for _, fieldAttr := range ctx.AllFieldAttribute() {
			switch {
			case fieldAttr.CalculatedFromAttribute() != nil:
				formattedDsl.WriteString(v.VisitCalculatedFromAttribute(fieldAttr.CalculatedFromAttribute().(*gen.CalculatedFromAttributeContext)).(string))
			case fieldAttr.LengthOfAttribute() != nil:
				formattedDsl.WriteString(v.VisitLengthOfAttribute(fieldAttr.LengthOfAttribute().(*gen.LengthOfAttributeContext)).(string))
			case fieldAttr.PaddingAttribute() != nil:
				formattedDsl.WriteString(v.VisitPaddingAttribute(fieldAttr.PaddingAttribute().(*gen.PaddingAttributeContext)).(string))
			case fieldAttr.TagAttribute() != nil:
				tagValue := fieldAttr.TagAttribute().(*gen.TagAttributeContext).DIGITS().GetText()
				formattedDsl.WriteString(fmt.Sprintf("@tag(%s)", tagValue))
			default:
				formattedDsl.WriteString(fieldAttr.GetText())
			}
			formattedDsl.WriteString("\n")
		}
	}
	formattedDsl.WriteString(v.VisitFieldDefinition(ctx.FieldDefinition()).(string))
	return formattedDsl.String()
}

// VisitLengthOfAttribute formats lengthOf attribute
func (v *PacketDslFormattor) VisitLengthOfAttribute(ctx *gen.LengthOfAttributeContext) interface{} {
	return fmt.Sprintf("@lengthOf(%s)", ctx.GetFrom().GetText())
}

// VisitCalculatedFromAttribute formats calculatedFrom attribute
func (v *PacketDslFormattor) VisitCalculatedFromAttribute(ctx *gen.CalculatedFromAttributeContext) interface{} {
	return fmt.Sprintf("@calculatedFrom(%s)", ctx.GetFrom().GetText())
}

// VisitPaddingAttribute formats padding attribute
func (v *PacketDslFormattor) VisitPaddingAttribute(ctx *gen.PaddingAttributeContext) interface{} {
	return fmt.Sprintf("%s(%s)", ctx.PADDING_ATTR().GetText(), ctx.PADDING_CHAR().GetText())
}

// VisitOptionDefinition overrides the default implementation for option definitions.
func (v *PacketDslFormattor) VisitOptionDefinition(ctx *gen.OptionDefinitionContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	// format options
	formattedDsl.WriteString("options {\n")
	for _, decl := range ctx.AllOptionDeclaration() {
		if d, ok := decl.(*gen.OptionDeclarationContext); ok {
			formattedDsl.WriteString(AddIndent4ln(v.VisitOptionDeclaration(d).(string)))
		}
	}
	formattedDsl.WriteString("}")
	formattedDsl.WriteString(v.getHiddenRightAtSameLine(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitOptionDeclaration overrides the default implementation for option declarations.
func (v *PacketDslFormattor) VisitOptionDeclaration(ctx *gen.OptionDeclarationContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	optionName := ctx.IDENTIFIER().GetText()
	optionValue := ctx.Value().GetText()
	formattedDsl.WriteString(fmt.Sprintf("%s = %s", optionName, optionValue))
	if ctx.SEMICOLON() != nil {
		formattedDsl.WriteString(";")
	}
	formattedDsl.WriteString(v.getHiddenRightAtSameLine(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitFieldDefinition overrides the default implementation for field definitions.
func (v *PacketDslFormattor) VisitFieldDefinition(ctx interface{}) interface{} {
	var b strings.Builder
	b.WriteString(v.getHiddenLeft(ctx.(antlr.ParserRuleContext).GetStart()))
	switch c := ctx.(type) {
	case *gen.ObjectFieldContext:
		field := ""
		if c.REPEAT() != nil {
			field = "repeat "
		}
		field += c.GetFtype().GetText()
		if c.GetFname() != nil {
			field += " " + c.GetFname().GetText()
		}
		b.WriteString(field + ",")
	case *gen.InerObjectFieldContext:
		b.WriteString(v.VisitInerObjectField(c).(string))
	case *gen.LengthFieldContext:
		b.WriteString(v.VisitLengthFieldDeclaration(c.LengthFieldDeclaration().(*gen.LengthFieldDeclarationContext)).(string))
	case *gen.CheckSumFieldContext:
		b.WriteString(v.VisitCheckSumFieldDeclaration(c.CheckSumFieldDeclaration().(*gen.CheckSumFieldDeclarationContext)).(string))
	case *gen.MetaFieldContext:
		repeat := ""
		if c.REPEAT() != nil {
			repeat = "repeat "
		}
		b.WriteString(repeat + v.VisitMetaDataDeclaration(c.MetaDataDeclaration().(*gen.MetaDataDeclarationContext)).(string))
	case *gen.MatchFieldContext:
		code := v.VisitMatchFieldDeclaration(c.MatchFieldDeclaration().(*gen.MatchFieldDeclarationContext)).(string)
		if (ctx.(*gen.MatchFieldContext)).COMMA() != nil {
			code += ","
		}
		b.WriteString(code)
	default:
		return "error"
	}
	b.WriteString(v.getHiddenRightAtSameLine(ctx.(antlr.ParserRuleContext).GetStop()))
	return b.String()
}

// VisitInerObjectField overrides the default implementation for inner object fields.
func (v *PacketDslFormattor) VisitInerObjectField(ctx *gen.InerObjectFieldContext) interface{} {
	var formattedDsl strings.Builder
	inerObjectDeclaration := ctx.InerObjectDeclaration()
	if ctx.REPEAT() != nil {
		formattedDsl.WriteString("repeat ")
	}
	if nil != inerObjectDeclaration.IDENTIFIER() {
		formattedDsl.WriteString(inerObjectDeclaration.IDENTIFIER().GetText() + " ")
	}
	formattedDsl.WriteString("{\n")

	// iterate over all field definitions
	for _, decl := range inerObjectDeclaration.AllFieldDefinition() {
		result := v.VisitFieldDefinition(decl).(string)
		formattedDsl.WriteString(AddIndent4ln(result))
	}

	formattedDsl.WriteString("},")
	return formattedDsl.String()
}

// VisitMetaDataDefinition overrides the default implementation for metadata definitions.
func (v *PacketDslFormattor) VisitMetaDataDefinition(ctx *gen.MetaDataDefinitionContext) interface{} {
	var formattedDsl strings.Builder
	metaName := ctx.IDENTIFIER().GetText()
	formattedDsl.WriteString(fmt.Sprintf("MetaData %s {\n", metaName))

	for _, decl := range ctx.AllMetaDataDeclaration() {
		if d, ok := decl.(*gen.MetaDataDeclarationContext); ok {
			result := v.VisitMetaDataDeclaration(d).(string)
			formattedDsl.WriteString(AddIndent4ln(result))
		}
	}

	formattedDsl.WriteString("}")
	return formattedDsl.String()
}

// VisitLengthFieldDeclaration format lengthof field
func (v *PacketDslFormattor) VisitLengthFieldDeclaration(ctx *gen.LengthFieldDeclarationContext) interface{} {
	desc := ""
	if ctx.STRING_LITERAL() != nil {
		desc = " " + ctx.STRING_LITERAL().GetText()
	}
	typ := ""
	if ctx.Type_() != nil {
		typ = ctx.Type_().GetText() + " "
	}
	lengthOfAttribute := ctx.LengthOfAttribute()
	return fmt.Sprintf("%s%s @lengthOf(%s)%s,", typ, ctx.GetName().GetText(), lengthOfAttribute.GetFrom().GetText(), desc)
}

// VisitCheckSumFieldDeclaration format check sum field
func (v *PacketDslFormattor) VisitCheckSumFieldDeclaration(ctx *gen.CheckSumFieldDeclarationContext) interface{} {
	desc := ""
	if ctx.STRING_LITERAL() != nil {
		desc = " " + ctx.STRING_LITERAL().GetText()
	}
	typ := ""
	if ctx.Type_() != nil {
		typ = ctx.Type_().GetText() + " "
	}
	checkSumAttribute := ctx.CalculatedFromAttribute()
	return fmt.Sprintf("%s%s @calculatedFrom(%s)%s,", typ, ctx.GetName().GetText(), checkSumAttribute.GetFrom().GetText(), desc)
}

// VisitMetaDataDeclaration for visiting metadata declarations.
func (v *PacketDslFormattor) VisitMetaDataDeclaration(ctx *gen.MetaDataDeclarationContext) interface{} {
	var formattedDsl strings.Builder
	typeName := ""
	if ctx.Type_() != nil {
		typeName = ctx.Type_().GetText()
	}
	fieldName := ctx.GetName().GetText()
	description := ""
	if ctx.STRING_LITERAL() != nil {
		description = ctx.STRING_LITERAL().GetText()
	}

	formattedDsl.WriteString(strings.TrimSpace(fmt.Sprintf("%s %s %s", typeName, fieldName, description)))
	if ctx.COMMA() != nil {
		formattedDsl.WriteString(",")
	}
	return formattedDsl.String()
}

// VisitMatchFieldDeclaration for visiting match fields.
func (v *PacketDslFormattor) VisitMatchFieldDeclaration(ctx *gen.MatchFieldDeclarationContext) interface{} {
	matchKey := ctx.GetMatchKey().GetText()
	matchName := ctx.GetMatchName().GetText()
	var formattedDsl strings.Builder
	formattedDsl.WriteString(fmt.Sprintf("match %s as %s {\n", matchKey, matchName))
	for _, pairCtx := range ctx.AllMatchPair() {
		lineComment := strings.TrimRight(v.getHiddenLeft(pairCtx.GetStart()), "\n")
		if lineComment != "" {
			formattedDsl.WriteString(AddIndent4ln(lineComment))
		}
		key := ""
		switch {
		case pairCtx.STRING() != nil:
			key = pairCtx.STRING().GetText()
		case pairCtx.DIGITS() != nil:
			key = pairCtx.DIGITS().GetText()
		case pairCtx.List() != nil:
			items := []string{}
			// digit list
			for _, num := range pairCtx.List().AllDIGITS() {
				items = append(items, num.GetText())
			}
			// string list
			for _, num := range pairCtx.List().AllSTRING() {
				items = append(items, num.GetText())
			}
			key = formatStringList(items, 5)
		}

		value := pairCtx.IDENTIFIER().GetText()
		if pairCtx.COMMA() != nil {
			value = strings.TrimSpace(value) + ","
		}
		formattedDsl.WriteString(AddIndent4ln(fmt.Sprintf("%s : %s", key, value)))
		lineComment = strings.TrimRight(v.getHiddenRightAtSameLine(pairCtx.GetStop()), "\n")
		if lineComment != "" {
			formattedDsl.WriteString(AddIndent4ln(lineComment))
		}
	}
	formattedDsl.WriteString("}")
	return formattedDsl.String()
}

// VisitTerminal Visit Terminal
func (v *PacketDslFormattor) VisitTerminal(node antlr.TerminalNode) interface{} {
	token := node.GetSymbol()
	switch token.GetTokenType() {
	case gen.PacketDslParserLINE_COMMENT:
		return token.GetText()
	default:
		return token.GetText()
	}
}

func formatStringList(values []string, itemsPerLine int) string {
	var b strings.Builder
	if len(values) <= itemsPerLine {
		b.WriteString(strings.Join(values, ", "))
		return "[" + b.String() + "]"
	}
	for idx, value := range values {
		if idx != 0 && idx%itemsPerLine == 0 {
			b.WriteString("\n")
		}
		b.WriteString(value)
		if idx != len(values)-1 {
			b.WriteString(",")
			if (idx+1)%itemsPerLine != 0 {
				b.WriteString(" ")
			}
		}
	}
	return "[\n" + AddIndent4ln(b.String()) + "]"
}
