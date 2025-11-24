package parser

import (
	"fmt"
	"strconv"
	"strings"

	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// ParseFile parses a DSL file and returns a slice of Packet models or an error.
func ParseFile(filename string) (interface{}, error) {
	// Create a new parser by reading the file
	parser, _, err := NewPacketDslParserByFile(filename)
	if err != nil {
		return nil, err
	}
	listener := NewSyntaxErrorListener()
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	// Invoke the root rule 'Packet' to parse the file
	tree := parser.Packet()
	if listener.HasErrors() {
		return nil, fmt.Errorf("syntax errors found: %v", listener.Errors)
	}
	// Use the custom visitor to build the model from the parse tree
	visitor := NewPacketDslVisitor()
	astData := tree.Accept(visitor)
	return astData, nil
}

// PacketDslVisitorImpl visits parse tree nodes and constructs model.Packet and model.Field.
type PacketDslVisitorImpl struct {
	*gen.BasePacketDslVisitor
	BinModel *model.BinaryModel
}

// NewPacketDslVisitor returns a new PacketDslVisitorImpl.
func NewPacketDslVisitor() *PacketDslVisitorImpl {
	return &PacketDslVisitorImpl{
		BasePacketDslVisitor: &gen.BasePacketDslVisitor{},
		BinModel:             model.NewBinaryModel(),
	}
}

func (v *PacketDslVisitorImpl) metaDataDeclarationToMetaData(ctx *gen.MetaDataDeclarationContext) interface{} {
	return model.MetaData{
		Name:        ctx.GetName().GetText(),
		Typ:         ctx.Type_().GetText(),
		BasicType:   ctx.Type_().GetText(),
		Description: ctx.STRING_LITERAL().GetText(),
		Line:        ctx.GetStart().GetLine(),
		Column:      ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitPacket handles a list of packetDefinition nodes and returns []model.Packet.
func (v *PacketDslVisitorImpl) VisitPacket(ctx *gen.PacketContext) interface{} {

	// MetaData map
	for _, metaDataDifinition := range ctx.AllMetaDataDefinition() {
		for _, decl := range metaDataDifinition.GetChildren() {
			switch c := decl.(type) {
			case *gen.RefMetaDataDeclarationContext:
				result := v.VisitRefMetaDataDeclaration(c).(model.MetaData)
				v.BinModel.AddMetaData(result)
			case *gen.MetaDataDeclarationContext:
				result := v.metaDataDeclarationToMetaData(c).(model.MetaData)
				v.BinModel.AddMetaData(result)
			default:
				continue
			}
		}
	}

	// Options map
	for _, optionDefinition := range ctx.AllOptionDefinition() {
		for _, option := range optionDefinition.AllOptionDeclaration() {
			name := option.IDENTIFIER().GetText()
			value := option.Value().GetText()
			if option.Value().STRING() != nil {
				value = strings.Trim(value, "\"")
			}
			// Store option in the map
			v.BinModel.AddOption(name, value, option.GetStart().GetLine(), option.GetStart().GetTokenSource().GetCharPositionInLine())
		}
	}

	fmt.Println("Options:", v.BinModel.Options)

	for _, def := range ctx.AllPacketDefinition() {
		packet := def.Accept(v).(*model.Packet)
		v.BinModel.AddPacket(packet) // Store in PacketMap
	}
	v.BinModel.ResolveDependencies()
	return v.BinModel
}

// VisitPacketDefinition constructs a model.Packet from a packetDefinition context.
func (v *PacketDslVisitorImpl) VisitPacketDefinition(ctx *gen.PacketDefinitionContext) interface{} {
	// Packet name from IDENTIFIER
	name := ctx.IDENTIFIER().GetText()
	// Check if 'root' keyword is present
	isRoot := ctx.ROOT() != nil

	// Iterate over all fieldDefinition children
	var fields []*model.Field
	var fieldMap = make(map[string]*model.Field)
	var lengthField *model.Field
	var matchFields = make(map[string][]model.MatchPair)
	for _, fctx := range ctx.AllFieldDefinitionWithAttribute() {
		if fc, ok := fctx.(*gen.FieldDefinitionWithAttributeContext); ok {
			fd := v.VisitFieldDefinitionWithAttribute(fc)
			if fd == nil {
				continue
			}
			fld := fd.(*model.Field) // Ensure type assertion
			if _, ok := fld.Attr.(*model.LengthFieldAttribute); ok {
				if !isRoot {
					v.BinModel.AddSyntaxError(&model.SyntaxError{
						Line:            fctx.GetStart().GetLine(),
						Column:          fctx.GetStart().GetTokenSource().GetCharPositionInLine(),
						Msg:             "LengthOfField can only be declared in the root packet",
						OffendingSymbol: nil,
					})
					continue
				}
				if lengthField != nil {
					v.BinModel.AddSyntaxError(&model.SyntaxError{
						Line:            fctx.GetStart().GetLine(),
						Column:          fctx.GetStart().GetTokenSource().GetCharPositionInLine(),
						Msg:             "Duplicate LengthOfField declaration",
						OffendingSymbol: nil,
					})
					continue
				}
				lengthField = fld
			}

			fields = append(fields, fld)
			fieldMap[fld.Name] = fld

			if mf, ok := fld.Attr.(*model.MatchFieldAttribute); ok {
				matchFields[mf.MatchKeyField.Name] = mf.MatchPairs
			}
		}
	}

	for _, f := range fields {
		if lengthField != nil && f.Name == lengthField.Attr.(*model.LengthFieldAttribute).TragetField.Name {
			lengthField.LenAttr = &model.LengthOfAttribute{
				LengthField: lengthField,
			}
			f.LenAttr = lengthField.Attr
		}
		switch c := f.Attr.(type) {
		case *model.ObjectFieldAttribute:
			if !c.IsIner {
				c.RefPacket = v.BinModel.PacketsMap[c.PacketName]
			}
		case *model.LengthFieldAttribute:
			f.Attr = &model.LengthFieldAttribute{
				LengthType:  f.GetType(),
				TragetField: fieldMap[c.TragetField.Name],
			}
		case *model.MatchFieldAttribute:
			c.MatchKeyField = fieldMap[c.MatchKeyField.Name]

		}
	}

	return &model.Packet{
		Name:        name,
		IsRoot:      isRoot,
		LengthField: lengthField,
		Fields:      fields,
		FieldMap:    fieldMap,
		MatchFields: matchFields,
		Line:        ctx.GetStart().GetLine(),
		Column:      ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitFieldDefinitionWithAttribute visit field definition with attribute
func (v *PacketDslVisitorImpl) VisitFieldDefinitionWithAttribute(ctx *gen.FieldDefinitionWithAttributeContext) interface{} {

	fd := v.VisitFieldDefinition(ctx.FieldDefinition())

	f := fd.(*model.Field)

	for _, fieldAttr := range ctx.AllFieldAttribute() {
		switch {
		case fieldAttr.CalculatedFromAttribute() != nil:
			f.Attr = &model.CheckSumFieldAttribute{CheckSumType: fieldAttr.CalculatedFromAttribute().GetFrom().GetText()}
		case fieldAttr.LengthOfAttribute() != nil:
			f.Attr = &model.LengthFieldAttribute{
				TragetField: &model.Field{Name: fieldAttr.LengthOfAttribute().GetFrom().GetText()},
				LengthType:  f.GetType(),
			}
		case fieldAttr.PaddingAttribute() != nil:
			f.Attr.(*model.FixedStringFieldAttribute).Padding = &model.Padding{
				PadChar: fieldAttr.PaddingAttribute().PADDING_CHAR().GetText(),
				PadLeft: strings.Contains(fieldAttr.PaddingAttribute().PADDING_ATTR().GetText(), "left"),
			}
		case fieldAttr.TagAttribute() != nil:
			tagValue := fieldAttr.TagAttribute().DIGITS().GetText()
			tagInt, _ := strconv.Atoi(tagValue)
			f.Tag = tagInt
		}

	}

	return fd

}

// VisitFieldDefinition dispatches to specific handlers based on field type.
func (v *PacketDslVisitorImpl) VisitFieldDefinition(ctx interface{}) interface{} {
	switch c := ctx.(type) {
	// Basic object field: REPEAT? IDENTIFIER
	case *gen.ObjectFieldContext:
		typ := c.GetFtype().GetText()
		name := typ
		if c.GetFname() != nil {
			name = c.GetFname().GetText()
		}
		var attr model.FieldAttribute
		if v.BinModel.MetaDataMap[typ] != (model.MetaData{}) {
			// If metadata exists, use its basic type
			typ = v.BinModel.MetaDataMap[typ].BasicType
			attr = &model.BasicFieldAttribute{
				Type: typ,
			}
		} else {
			attr = &model.ObjectFieldAttribute{
				IsIner:     false,
				PacketName: typ,
			}
		}
		return &model.Field{
			Name:     name,
			Attr:     attr,
			IsRepeat: c.REPEAT() != nil,
			Line:     c.GetStart().GetLine(),
			Column:   c.GetStart().GetTokenSource().GetCharPositionInLine(),
		}

	// Nested object field: REPEAT? IDENTIFIER '{' fieldDefinition+ '}'
	case *gen.InerObjectFieldContext:
		return v.VisitInerObjectField(c)
	case *gen.LengthFieldContext:
		return v.VisitLengthFieldDeclaration(c.LengthFieldDeclaration().(*gen.LengthFieldDeclarationContext))
	case *gen.CheckSumFieldContext:
		return v.VisitCheckSumFieldDeclaration(c.CheckSumFieldDeclaration().(*gen.CheckSumFieldDeclarationContext))
	// Metadata field: REPEAT? metaDataDeclaration
	case *gen.MetaFieldContext:
		mctx := c.MetaDataDeclaration().(*gen.MetaDataDeclarationContext)
		fld := v.metaDataDeclarationToField(mctx).(*model.Field)
		if c.REPEAT() != nil {
			fld.IsRepeat = true
		}
		return fld

	// Match field: match IDENTIFIER '{' matchPair+ '}'
	case *gen.MatchFieldContext:
		return v.VisitMatchFieldDeclaration(c.MatchFieldDeclaration().(*gen.MatchFieldDeclarationContext))
	default:
		v.BinModel.AddSyntaxError(&model.SyntaxError{
			Line:            ctx.(*gen.FieldDefinitionContext).GetStart().GetLine(),
			Column:          ctx.(*gen.FieldDefinitionContext).GetStart().GetTokenSource().GetCharPositionInLine(),
			Msg:             "Unexpected field definition type",
			OffendingSymbol: nil,
		})
	}
	return nil
}

// VisitLengthFieldDeclaration parse lengthof field
func (v *PacketDslVisitorImpl) VisitLengthFieldDeclaration(ctx *gen.LengthFieldDeclarationContext) interface{} {
	desc := ""
	if ctx.STRING_LITERAL() != nil {
		desc = ctx.STRING_LITERAL().GetText()
	}
	name := ctx.GetName().GetText()
	typ := ctx.GetName().GetText()
	if ctx.Type_() != nil {
		typ = ctx.Type_().GetText()
	}
	if v.BinModel.MetaDataMap[name] != (model.MetaData{}) {
		// If metadata exists, use its basic type
		typ = v.BinModel.MetaDataMap[name].BasicType
	}
	return &model.Field{
		Name:     name,
		IsRepeat: false,
		Attr: &model.LengthFieldAttribute{
			TragetField: &model.Field{Name: ctx.LengthOfAttribute().GetFrom().GetText()},
			LengthType:  typ,
		},
		Doc:    desc,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitCheckSumFieldDeclaration parse check sum field
func (v *PacketDslVisitorImpl) VisitCheckSumFieldDeclaration(ctx *gen.CheckSumFieldDeclarationContext) interface{} {
	desc := ""
	if ctx.STRING_LITERAL() != nil {
		desc = ctx.STRING_LITERAL().GetText()
	}
	name := ctx.GetName().GetText()
	typ := ctx.GetName().GetText()
	if ctx.Type_() != nil {
		typ = ctx.Type_().GetText()
	}
	if v.BinModel.MetaDataMap[name] != (model.MetaData{}) {
		// If metadata exists, use its basic type
		typ = v.BinModel.MetaDataMap[name].BasicType
	}
	return &model.Field{
		Name:     ctx.GetName().GetText(),
		IsRepeat: false,
		Attr: &model.CheckSumFieldAttribute{
			Type:         typ,
			CheckSumType: ctx.CalculatedFromAttribute().GetFrom().GetText(),
		},
		Doc:    desc,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitInerObjectField handles a nested object declaration and returns model.Field.
func (v *PacketDslVisitorImpl) VisitInerObjectField(ctx *gen.InerObjectFieldContext) interface{} {
	decl := ctx.InerObjectDeclaration()
	name := decl.IDENTIFIER().GetText()
	var subFields []*model.Field
	// Iterate all sub-field definitions inside the nested object
	for _, fctx := range decl.AllFieldDefinition() {
		fld := v.VisitFieldDefinition(fctx)
		if fld == nil {
			continue
		}
		f := fld.(*model.Field)
		subFields = append(subFields, f)
	}
	// Construct nested Packet model
	p := model.Packet{
		Name:   name,
		IsRoot: false,
		Fields: subFields,
	}
	return &model.Field{
		Name:     name,
		IsRepeat: ctx.REPEAT() != nil,
		Attr: &model.ObjectFieldAttribute{
			IsIner:     true,
			RefPacket:  &p,
			PacketName: p.Name,
		},
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

func (v *PacketDslVisitorImpl) metaDataDeclarationToField(ctx *gen.MetaDataDeclarationContext) interface{} {
	// Field name
	name := ctx.GetName().GetText()
	// Determine type if present
	var attr model.FieldAttribute
	if ctx.Type_().BasicType() != nil {
		attr = &model.BasicFieldAttribute{
			Type: ctx.Type_().GetText(),
		}
	} else if ctx.Type_().FixedString() != nil {
		size, _ := strconv.Atoi(ctx.Type_().FixedString().DIGITS().GetText())
		if strings.Contains(ctx.Type_().GetText(), "zchar") {
			attr = &model.FixedStringFieldAttribute{
				Length:  size,
				Padding: &model.Padding{PadChar: "'\x00'", PadLeft: false},
			}
		} else {
			attr = &model.FixedStringFieldAttribute{
				Length: size,
			}
		}
	} else if ctx.Type_().DynamicString() != nil {
		attr = &model.DynamicStringFieldAttribute{Type: "string"}
	}
	// Extract documentation string if present, by removing backticks
	var doc string
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		doc = raw[1 : len(raw)-1]
	}
	return &model.Field{
		Name:     name,
		Attr:     attr,
		IsRepeat: false,
		Doc:      doc,
		Line:     ctx.GetStart().GetLine(),
		Column:   ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitMatchFieldDeclaration handles match fields: match IDENTIFIER '{' matchPair+ '}'.
func (v *PacketDslVisitorImpl) VisitMatchFieldDeclaration(ctx *gen.MatchFieldDeclarationContext) interface{} {
	matchKey := ctx.GetMatchKey().GetText()
	matchName := ctx.GetMatchName().GetText()
	var pairs []model.MatchPair
	for _, pairCtx := range ctx.AllMatchPair() {
		mp := pairCtx.Accept(v).([]model.MatchPair)
		pairs = append(pairs, mp...)
	}
	var pairsMap = make(map[string]interface{})
	for _, pair := range pairs {
		if _, exists := pairsMap[pair.Key]; exists {
			v.BinModel.AddSyntaxError(&model.SyntaxError{
				Line:            pair.Line,
				Column:          pair.Column,
				Msg:             "Duplicate match key: " + pair.Key,
				OffendingSymbol: nil,
			})
			continue
		}
	}
	return &model.Field{
		Name:     matchName,
		IsRepeat: false,
		Attr: &model.MatchFieldAttribute{
			MatchKeyField: &model.Field{Name: matchKey},
			MatchPairs:    pairs,
		},
	}
}

// VisitMatchPair handles an individual match key-value: (DIGITS|STRING|list) ':' IDENTIFIER
func (v *PacketDslVisitorImpl) VisitMatchPair(ctx *gen.MatchPairContext) interface{} {
	var pairs []model.MatchPair
	val := ctx.IDENTIFIER().GetText()
	var key string
	if ctx.DIGITS() != nil {
		key = ctx.DIGITS().GetText()
	} else if ctx.STRING() != nil {
		key = ctx.STRING().GetText()
	} else if ctx.List() != nil {
		for _, k := range ctx.List().AllDIGITS() {
			pairs = append(pairs, model.MatchPair{Key: k.GetText(), Value: val, Line: k.GetSymbol().GetLine(), Column: k.GetSymbol().GetTokenSource().GetCharPositionInLine()})
		}
		for _, k := range ctx.List().AllSTRING() {
			pairs = append(pairs, model.MatchPair{Key: k.GetText(), Value: val, Line: k.GetSymbol().GetLine(), Column: k.GetSymbol().GetTokenSource().GetCharPositionInLine()})
		}
		return pairs
	}

	return append(pairs, model.MatchPair{Key: key, Value: val, Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetTokenSource().GetCharPositionInLine()})
}

// VisitRefMetaDataDeclaration handles reference metadata declarations.
func (v *PacketDslVisitorImpl) VisitRefMetaDataDeclaration(ctx *gen.RefMetaDataDeclarationContext) interface{} {
	return model.MetaData{
		Name:        ctx.GetName().GetText(),
		Typ:         ctx.GetTyp().GetText(),
		BasicType:   ctx.GetTyp().GetText(),
		Description: ctx.STRING_LITERAL().GetText(),
		Line:        ctx.GetStart().GetLine(),
		Column:      ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitType logs visiting a type context; can be extended for type validation.
func (v *PacketDslVisitorImpl) VisitType(ctx *gen.TypeContext) interface{} {
	fmt.Println("Visiting Type:", ctx.GetText())
	return nil
}
