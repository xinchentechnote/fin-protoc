// Code generated from grammar/PacketDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // PacketDsl
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PacketDslParser.
type PacketDslVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PacketDslParser#packet.
	VisitPacket(ctx *PacketContext) interface{}

	// Visit a parse tree produced by PacketDslParser#optionDefinition.
	VisitOptionDefinition(ctx *OptionDefinitionContext) interface{}

	// Visit a parse tree produced by PacketDslParser#optionDeclaration.
	VisitOptionDeclaration(ctx *OptionDeclarationContext) interface{}

	// Visit a parse tree produced by PacketDslParser#packetDefinition.
	VisitPacketDefinition(ctx *PacketDefinitionContext) interface{}

	// Visit a parse tree produced by PacketDslParser#fieldDefinition.
	VisitFieldDefinition(ctx *FieldDefinitionContext) interface{}

	// Visit a parse tree produced by PacketDslParser#metaDataDefinition.
	VisitMetaDataDefinition(ctx *MetaDataDefinitionContext) interface{}

	// Visit a parse tree produced by PacketDslParser#metaDataDeclaration.
	VisitMetaDataDeclaration(ctx *MetaDataDeclarationContext) interface{}

	// Visit a parse tree produced by PacketDslParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by PacketDslParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by PacketDslParser#matchField.
	VisitMatchField(ctx *MatchFieldContext) interface{}

	// Visit a parse tree produced by PacketDslParser#matchPair.
	VisitMatchPair(ctx *MatchPairContext) interface{}

	// Visit a parse tree produced by PacketDslParser#list.
	VisitList(ctx *ListContext) interface{}
}
