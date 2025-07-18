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

	// Invoke the root rule 'Packet' to parse the file
	tree := parser.Packet()

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
			v.BinModel.AddOption(name, value)
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
	var lengthOfField string
	var matchFields = make(map[string][]model.MatchPair)
	for _, fctx := range ctx.AllFieldDefinition() {
		fld := v.VisitFieldDefinition(fctx).(model.Field)

		if fld.LengthOfField != "" {
			if !isRoot {
				panic("LengthOfField can only be declared in the root packet")
			}
			if lengthOfField != "" {
				panic("duplicate LengthOfField declaration")
			}
			lengthOfField = fld.LengthOfField
		}

		fields = append(fields, fld)
		fieldMap[fld.Name] = fld

		if fld.Type == "match" {
			matchFields[fld.MatchKey] = fld.MatchPairs
		}
	}

	return model.Packet{
		Name:          name,
		IsRoot:        isRoot,
		LengthOfField: lengthOfField,
		Fields:        fields,
		FieldMap:      fieldMap,
		MatchFields:   matchFields,
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
		}

	// Nested object field: REPEAT? IDENTIFIER '{' fieldDefinition+ '}'
	case *gen.InerObjectFieldContext:
		return v.VisitInerObjectField(c).(model.Field)

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
		panic(fmt.Sprintf("unexpected context type in VisitFieldDefinition: %T", ctx))
	}
}

// VisitInerObjectField handles a nested object declaration and returns model.Field.
func (v *PacketDslVisitorImpl) VisitInerObjectField(ctx *gen.InerObjectFieldContext) interface{} {
	decl := ctx.InerObjectDeclaration()
	name := decl.IDENTIFIER().GetText()
	var subFields []model.Field
	// Iterate all sub-field definitions inside the nested object
	for _, fctx := range decl.AllFieldDefinition() {
		fld := v.VisitFieldDefinition(fctx).(model.Field)
		subFields = append(subFields, fld)
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
	var lengthOfField string
	if ctx.GetFrom() != nil {
		lengthOfField = ctx.GetFrom().GetText()
	}
	// Extract documentation string if present, by removing backticks
	var doc string
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		doc = raw[1 : len(raw)-1]
	}
	return model.Field{
		Name:          name,
		Type:          typ,
		LengthOfField: lengthOfField,
		IsRepeat:      false,
		Doc:           doc,
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
			pairs = append(pairs, model.MatchPair{Key: k.GetText(), Value: val})
		}
		for _, k := range ctx.List().AllSTRING() {
			pairs = append(pairs, model.MatchPair{Key: k.GetText(), Value: val})
		}
		return pairs
	}

	return append(pairs, model.MatchPair{Key: key, Value: val})
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
