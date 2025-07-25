package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
)

// FormatPacketDsl formats the packet DSL string.
func FormatPacketDsl(dsl string) (string, error) {
	// 创建一个新的语法分析器
	parser, stream, err := NewPacketDslParserByContent(dsl)
	if err != nil {
		return "", fmt.Errorf("could not create parser: %v", err)
	}
	listener := NewSyntaxErrorListener()
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	// 设置语法规则，开始解析
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

func (v *PacketDslFormattor) getHiddenRight(token antlr.Token) string {
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
			v.lineComments[t] = struct{}{}

			sb.WriteString(t.GetText())
			sb.WriteString("\n")
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
	formattedDsl.WriteString(v.getHiddenRight(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitPacketDefinition overrides the default implementation for packet definitions.
func (v *PacketDslFormattor) VisitPacketDefinition(ctx *gen.PacketDefinitionContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
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
		formatted := v.VisitFieldDefinition(fieldCtx).(string)
		formattedDsl.WriteString(AddIndent4ln(formatted))
	}
	formattedDsl.WriteString("}")
	formattedDsl.WriteString(v.getHiddenRight(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitOptionDefinition overrides the default implementation for option definitions.
func (v *PacketDslFormattor) VisitOptionDefinition(ctx *gen.OptionDefinitionContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	// 处理 option 声明
	formattedDsl.WriteString("options {\n")
	for _, decl := range ctx.AllOptionDeclaration() {
		if d, ok := decl.(*gen.OptionDeclarationContext); ok {
			formattedDsl.WriteString("\t" + v.VisitOptionDeclaration(d).(string) + "\n")
		}
	}
	formattedDsl.WriteString("}")
	formattedDsl.WriteString(v.getHiddenRight(ctx.GetStop()))
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
	formattedDsl.WriteString(v.getHiddenRight(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitFieldDefinition overrides the default implementation for field definitions.
func (v *PacketDslFormattor) VisitFieldDefinition(ctx interface{}) interface{} {
	switch c := ctx.(type) {
	case *gen.ObjectFieldContext:
		repeat := ""
		if c.REPEAT() != nil {
			repeat = "repeat "
		}
		return repeat + c.IDENTIFIER().GetText()
	case *gen.InerObjectFieldContext:
		return v.VisitInerObjectField(c).(string)
	case *gen.MetaFieldContext:
		repeat := ""
		if c.REPEAT() != nil {
			repeat = "repeat "
		}
		return repeat + v.VisitMetaDataDeclaration(c.MetaDataDeclaration().(*gen.MetaDataDeclarationContext)).(string)
	case *gen.MatchFieldContext:
		code := v.VisitMatchFieldDeclaration(c.MatchFieldDeclaration().(*gen.MatchFieldDeclarationContext)).(string)
		if (ctx.(*gen.MatchFieldContext)).COMMA() != nil {
			code += ","
		}
		return code
	default:
		return "error"
	}
}

// VisitInerObjectField overrides the default implementation for inner object fields.
func (v *PacketDslFormattor) VisitInerObjectField(ctx *gen.InerObjectFieldContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	inerObjectDeclaration := ctx.InerObjectDeclaration()
	if ctx.REPEAT() != nil {
		formattedDsl.WriteString("repeat ")
	}
	if nil != inerObjectDeclaration.IDENTIFIER() {
		formattedDsl.WriteString(inerObjectDeclaration.IDENTIFIER().GetText() + " ")
	}
	formattedDsl.WriteString("{\n")

	// 遍历所有字段声明
	for _, decl := range inerObjectDeclaration.AllFieldDefinition() {
		result := v.VisitFieldDefinition(decl).(string)
		formattedDsl.WriteString("\t\t" + result + "\n")
	}

	formattedDsl.WriteString("\t},")
	formattedDsl.WriteString(v.getHiddenRight(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitMetaDataDefinition overrides the default implementation for metadata definitions.
func (v *PacketDslFormattor) VisitMetaDataDefinition(ctx *gen.MetaDataDefinitionContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	metaName := ctx.IDENTIFIER().GetText()
	formattedDsl.WriteString(fmt.Sprintf("MetaData %s {\n", metaName))

	for _, decl := range ctx.AllMetaDataDeclaration() {
		if d, ok := decl.(*gen.MetaDataDeclarationContext); ok {
			result := v.VisitMetaDataDeclaration(d).(string)
			formattedDsl.WriteString("\t" + result + "\n")
		}
	}

	formattedDsl.WriteString("}")
	formattedDsl.WriteString(v.getHiddenRight(ctx.GetStop()))
	return formattedDsl.String()
}

// VisitMetaDataDeclaration for visiting metadata declarations.
func (v *PacketDslFormattor) VisitMetaDataDeclaration(ctx *gen.MetaDataDeclarationContext) interface{} {
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	typeName := ""
	if ctx.Type_() != nil {
		typeName = ctx.Type_().GetText()
	}
	fieldName := ctx.GetName().GetText()
	if ctx.GetFrom() != nil {
		fieldName += " = " + "lengthof(" + ctx.GetFrom().GetText() + ")"
	}
	description := ""
	if ctx.STRING_LITERAL() != nil {
		description = ctx.STRING_LITERAL().GetText()
	}

	formattedDsl.WriteString(strings.TrimSpace(fmt.Sprintf("%s %s %s", typeName, fieldName, description)))
	if ctx.COMMA() != nil {
		formattedDsl.WriteString(",")
	}
	lineComment := v.getHiddenRight(ctx.GetStop())
	if lineComment != "" {
		formattedDsl.WriteString("\n" + lineComment)
	}
	return formattedDsl.String()
}

// VisitMatchFieldDeclaration for visiting match fields.
func (v *PacketDslFormattor) VisitMatchFieldDeclaration(ctx *gen.MatchFieldDeclarationContext) interface{} {
	matchKey := ctx.GetMatchKey().GetText()
	matchName := ctx.GetMatchName().GetText()
	var formattedDsl strings.Builder
	formattedDsl.WriteString(v.getHiddenLeft(ctx.GetStart()))
	formattedDsl.WriteString(fmt.Sprintf("match %s as %s {\n", matchKey, matchName))
	for _, pairCtx := range ctx.AllMatchPair() {
		lineComment := strings.TrimRight(v.getHiddenLeft(pairCtx.GetStart()), "\n")
		if lineComment != "" {
			formattedDsl.WriteString(AddIndent4ln(lineComment))
		}
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
			key = formatStringList(items, 5)
		}

		value := pairCtx.IDENTIFIER().GetText()
		if pairCtx.COMMA() != nil {
			value = strings.TrimSpace(value) + ","
		}
		formattedDsl.WriteString(AddIndent4ln(fmt.Sprintf("%s : %s", key, value)))
		lineComment = strings.TrimRight(v.getHiddenRight(pairCtx.GetStop()), "\n")
		if lineComment != "" {
			formattedDsl.WriteString(AddIndent4ln(lineComment))
		}
	}
	formattedDsl.WriteString("}")
	formattedDsl.WriteString(v.getHiddenRight(ctx.GetStop()))
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
