package parser

import (
	"fmt"
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

// VisitPacket handles a list of packetDefinition nodes and returns []model.Packet.
func (v *PacketDslVisitorImpl) VisitPacket(ctx *gen.PacketContext) interface{} {

	// MetaData map
	for _, metaDataDifinition := range ctx.AllMetaDataDefinition() {
		for _, metaDataDeclaration := range metaDataDifinition.AllMetaDataDeclaration() {
			name := metaDataDeclaration.GetName().GetText()
			typ := metaDataDeclaration.Type_().GetText()
			basicType := metaDataDeclaration.Type_().GetText()
			// Store metadata in the map
			v.BinModel.AddMetaData(model.MetaData{
				Name:        name,
				Typ:         typ,
				BasicType:   basicType,
				Description: metaDataDeclaration.STRING_LITERAL().GetText(),
				Line:        metaDataDeclaration.GetStart().GetLine(),
				Column:      metaDataDeclaration.GetStart().GetTokenSource().GetCharPositionInLine(),
			})

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
		packet := def.Accept(v).(model.Packet)
		v.BinModel.AddPacket(packet) // Store in PacketMap
	}
	return v.BinModel
}

// VisitPacketDefinition constructs a model.Packet from a packetDefinition context.
func (v *PacketDslVisitorImpl) VisitPacketDefinition(ctx *gen.PacketDefinitionContext) interface{} {
	// Packet name from IDENTIFIER
	name := ctx.IDENTIFIER().GetText()
	// Check if 'root' keyword is present
	isRoot := ctx.ROOT() != nil

	// Iterate over all fieldDefinition children
	var fields []model.Field
	var fieldMap = make(map[string]model.Field)
	var lengthField *model.Field
	var matchFields = make(map[string][]model.MatchPair)
	for _, fctx := range ctx.AllFieldDefinition() {
		fd := v.VisitFieldDefinition(fctx)
		if fd == nil {
			continue
		}
		fld := fd.(model.Field) // Ensure type assertion
		if fld.LengthOfField != "" {
			if !isRoot {
				v.BinModel.AddSyntaxError(model.SyntaxError{
					Line:            fctx.GetStart().GetLine(),
					Column:          fctx.GetStart().GetTokenSource().GetCharPositionInLine(),
					Msg:             "LengthOfField can only be declared in the root packet",
					OffendingSymbol: nil,
				})
				continue
			}
			if lengthField != nil {
				v.BinModel.AddSyntaxError(model.SyntaxError{
					Line:            fctx.GetStart().GetLine(),
					Column:          fctx.GetStart().GetTokenSource().GetCharPositionInLine(),
					Msg:             "Duplicate LengthOfField declaration",
					OffendingSymbol: nil,
				})
				continue
			}
			lengthField = &fld
		}

		fields = append(fields, fld)
		fieldMap[fld.Name] = fld

		if fld.Type == "match" {
			matchFields[fld.MatchKey] = fld.MatchPairs
		}
	}

	return model.Packet{
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

// VisitFieldDefinition dispatches to specific handlers based on field type.
func (v *PacketDslVisitorImpl) VisitFieldDefinition(ctx interface{}) interface{} {
	switch c := ctx.(type) {
	// Basic object field: REPEAT? IDENTIFIER
	case *gen.ObjectFieldContext:
		name := c.IDENTIFIER().GetText()
		typ := name
		if v.BinModel.MetaDataMap[name] != (model.MetaData{}) {
			// If metadata exists, use its basic type
			typ = v.BinModel.MetaDataMap[name].BasicType
		}
		return model.Field{
			Name:     name,
			Type:     typ,
			IsRepeat: c.REPEAT() != nil,
			Line:     c.GetStart().GetLine(),
			Column:   c.GetStart().GetTokenSource().GetCharPositionInLine(),
		}

	// Nested object field: REPEAT? IDENTIFIER '{' fieldDefinition+ '}'
	case *gen.InerObjectFieldContext:
		return v.VisitInerObjectField(c).(model.Field)
	case *gen.LengthFieldContext:
		return v.VisitLengthFieldDeclaration(c.LengthFieldDeclaration().(*gen.LengthFieldDeclarationContext)).(model.Field)
	case *gen.CheckSumFieldContext:
		return v.VisitCheckSumFieldDeclaration(c.CheckSumFieldDeclaration().(*gen.CheckSumFieldDeclarationContext)).(model.Field)

	// Metadata field: REPEAT? metaDataDeclaration
	case *gen.MetaFieldContext:
		mctx := c.MetaDataDeclaration().(*gen.MetaDataDeclarationContext)
		fld := v.VisitMetaDataDeclaration(mctx).(model.Field)
		if c.REPEAT() != nil {
			fld.IsRepeat = true
		}
		return fld

	// Match field: match IDENTIFIER '{' matchPair+ '}'
	case *gen.MatchFieldContext:
		return v.VisitMatchFieldDeclaration(c.MatchFieldDeclaration().(*gen.MatchFieldDeclarationContext)).(model.Field)

	default:
		v.BinModel.AddSyntaxError(model.SyntaxError{
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
	return model.Field{
		Name:          name,
		Type:          typ,
		IsRepeat:      false,
		LengthOfField: ctx.LengthOfAttribute().GetFrom().GetText(),
		Doc:           desc,
		Line:          ctx.GetStart().GetLine(),
		Column:        ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
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
	return model.Field{
		Name:         ctx.GetName().GetText(),
		Type:         typ,
		IsRepeat:     false,
		CheckSumType: ctx.CalculatedFromAttribute().GetFrom().GetText(),
		Doc:          desc,
		Line:         ctx.GetStart().GetLine(),
		Column:       ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitInerObjectField handles a nested object declaration and returns model.Field.
func (v *PacketDslVisitorImpl) VisitInerObjectField(ctx *gen.InerObjectFieldContext) interface{} {
	decl := ctx.InerObjectDeclaration()
	name := decl.IDENTIFIER().GetText()
	var subFields []model.Field
	// Iterate all sub-field definitions inside the nested object
	for _, fctx := range decl.AllFieldDefinition() {
		fld := v.VisitFieldDefinition(fctx)
		if fld == nil {
			continue
		}
		subFields = append(subFields, fld.(model.Field))
	}
	// Construct nested Packet model
	p := model.Packet{
		Name:   name,
		IsRoot: false,
		Fields: subFields,
	}
	return model.Field{
		Name:       name,
		Type:       name, // nested objects do not have a basic Type
		IsRepeat:   ctx.REPEAT() != nil,
		InerObject: &p,
		Line:       ctx.GetStart().GetLine(),
		Column:     ctx.GetStart().GetTokenSource().GetCharPositionInLine(),
	}
}

// VisitMetaDataDeclaration handles a metadata declaration of the form: type? IDENTIFIER `doc`?
func (v *PacketDslVisitorImpl) VisitMetaDataDeclaration(ctx *gen.MetaDataDeclarationContext) interface{} {
	// Field name
	name := ctx.GetName().GetText()
	// Determine type if present
	var typ string
	if ctx.Type_() != nil {
		typ = ctx.Type_().GetText()
	} else {
		typ = name
	}
	// If metadata exists, use its basic type
	if meta, exists := v.BinModel.MetaDataMap[typ]; exists {
		typ = meta.BasicType
	}
	// Extract documentation string if present, by removing backticks
	var doc string
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		doc = raw[1 : len(raw)-1]
	}
	return model.Field{
		Name:     name,
		Type:     typ,
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
			v.BinModel.AddSyntaxError(model.SyntaxError{
				Line:            pair.Line,
				Column:          pair.Column,
				Msg:             "Duplicate match key: " + pair.Key,
				OffendingSymbol: nil,
			})
			continue
		}
	}
	return model.Field{
		Name:       matchName,
		Type:       "match",
		MatchKey:   matchKey,
		IsRepeat:   false,
		MatchPairs: pairs,
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

// VisitMetaDataDefinition is not used currently; just logs visiting metadata definitions.
func (v *PacketDslVisitorImpl) VisitMetaDataDefinition(ctx *gen.MetaDataDefinitionContext) interface{} {
	fmt.Println("Visiting MetaDataDefinition:", ctx.GetText())
	return nil
}

// VisitType logs visiting a type context; can be extended for type validation.
func (v *PacketDslVisitorImpl) VisitType(ctx *gen.TypeContext) interface{} {
	fmt.Println("Visiting Type:", ctx.GetText())
	return nil
}
