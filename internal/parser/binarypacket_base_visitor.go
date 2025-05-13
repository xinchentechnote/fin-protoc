// Code generated from BinaryPacket.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BinaryPacket

import "github.com/antlr4-go/antlr/v4"

type BaseBinaryPacketVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBinaryPacketVisitor) VisitPacketDefinition(ctx *PacketDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBinaryPacketVisitor) VisitFieldDefinition(ctx *FieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBinaryPacketVisitor) VisitMetaDataDefinition(ctx *MetaDataDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBinaryPacketVisitor) VisitMetaDataDeclaration(ctx *MetaDataDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBinaryPacketVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBinaryPacketVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBinaryPacketVisitor) VisitMatchField(ctx *MatchFieldContext) interface{} {
	return v.VisitChildren(ctx)
}
