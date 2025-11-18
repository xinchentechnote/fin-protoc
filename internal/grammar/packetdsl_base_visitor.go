// Code generated from grammar/PacketDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // PacketDsl
import "github.com/antlr4-go/antlr/v4"

type BasePacketDslVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePacketDslVisitor) VisitPacket(ctx *PacketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitOptionDefinition(ctx *OptionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitOptionDeclaration(ctx *OptionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitPacketDefinition(ctx *PacketDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitFieldDefinitionWithAttribute(ctx *FieldDefinitionWithAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitInerObjectField(ctx *InerObjectFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMetaField(ctx *MetaFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitObjectField(ctx *ObjectFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitLengthField(ctx *LengthFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitCheckSumField(ctx *CheckSumFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMatchField(ctx *MatchFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMetaDataDefinition(ctx *MetaDataDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitLengthFieldDeclaration(ctx *LengthFieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitCheckSumFieldDeclaration(ctx *CheckSumFieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitFieldAttribute(ctx *FieldAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitCalculatedFromAttribute(ctx *CalculatedFromAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitLengthOfAttribute(ctx *LengthOfAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitPaddingAttribute(ctx *PaddingAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitTagAttribute(ctx *TagAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMetaDataDeclaration(ctx *MetaDataDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitRefMetaDataDeclaration(ctx *RefMetaDataDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMatchFieldDeclaration(ctx *MatchFieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMatchPair(ctx *MatchPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitInerObjectDeclaration(ctx *InerObjectDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}
