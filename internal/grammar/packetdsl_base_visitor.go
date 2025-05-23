// Code generated from grammar/PacketDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // PacketDsl
import "github.com/antlr4-go/antlr/v4"

type BasePacketDslVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePacketDslVisitor) VisitPacket(ctx *PacketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitPacketDefinition(ctx *PacketDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitFieldDefinition(ctx *FieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMetaDataDefinition(ctx *MetaDataDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMetaDataDeclaration(ctx *MetaDataDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMatchField(ctx *MatchFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitMatchPair(ctx *MatchPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePacketDslVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}
