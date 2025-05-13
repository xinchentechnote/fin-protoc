// Code generated from BinaryPacket.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BinaryPacket

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BinaryPacketParser.
type BinaryPacketVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BinaryPacketParser#packetDefinition.
	VisitPacketDefinition(ctx *PacketDefinitionContext) interface{}

	// Visit a parse tree produced by BinaryPacketParser#fieldDefinition.
	VisitFieldDefinition(ctx *FieldDefinitionContext) interface{}

	// Visit a parse tree produced by BinaryPacketParser#metaDataDefinition.
	VisitMetaDataDefinition(ctx *MetaDataDefinitionContext) interface{}

	// Visit a parse tree produced by BinaryPacketParser#metaDataDeclaration.
	VisitMetaDataDeclaration(ctx *MetaDataDeclarationContext) interface{}

	// Visit a parse tree produced by BinaryPacketParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by BinaryPacketParser#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by BinaryPacketParser#matchField.
	VisitMatchField(ctx *MatchFieldContext) interface{}
}
