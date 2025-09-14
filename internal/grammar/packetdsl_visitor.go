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

	// Visit a parse tree produced by PacketDslParser#fieldDefinitionWithAttribute.
	VisitFieldDefinitionWithAttribute(ctx *FieldDefinitionWithAttributeContext) interface{}

	// Visit a parse tree produced by PacketDslParser#InerObjectField.
	VisitInerObjectField(ctx *InerObjectFieldContext) interface{}

	// Visit a parse tree produced by PacketDslParser#MetaField.
	VisitMetaField(ctx *MetaFieldContext) interface{}

	// Visit a parse tree produced by PacketDslParser#ObjectField.
	VisitObjectField(ctx *ObjectFieldContext) interface{}

	// Visit a parse tree produced by PacketDslParser#LengthField.
	VisitLengthField(ctx *LengthFieldContext) interface{}

	// Visit a parse tree produced by PacketDslParser#CheckSumField.
	VisitCheckSumField(ctx *CheckSumFieldContext) interface{}

	// Visit a parse tree produced by PacketDslParser#MatchField.
	VisitMatchField(ctx *MatchFieldContext) interface{}

	// Visit a parse tree produced by PacketDslParser#metaDataDefinition.
	VisitMetaDataDefinition(ctx *MetaDataDefinitionContext) interface{}

	// Visit a parse tree produced by PacketDslParser#lengthFieldDeclaration.
	VisitLengthFieldDeclaration(ctx *LengthFieldDeclarationContext) interface{}

	// Visit a parse tree produced by PacketDslParser#checkSumFieldDeclaration.
	VisitCheckSumFieldDeclaration(ctx *CheckSumFieldDeclarationContext) interface{}

	// Visit a parse tree produced by PacketDslParser#fieldAttribute.
	VisitFieldAttribute(ctx *FieldAttributeContext) interface{}

	// Visit a parse tree produced by PacketDslParser#lengthOfAttribute.
	VisitLengthOfAttribute(ctx *LengthOfAttributeContext) interface{}

	// Visit a parse tree produced by PacketDslParser#calculatedFromAttribute.
	VisitCalculatedFromAttribute(ctx *CalculatedFromAttributeContext) interface{}

	// Visit a parse tree produced by PacketDslParser#metaDataDeclaration.
	VisitMetaDataDeclaration(ctx *MetaDataDeclarationContext) interface{}

	// Visit a parse tree produced by PacketDslParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by PacketDslParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by PacketDslParser#matchFieldDeclaration.
	VisitMatchFieldDeclaration(ctx *MatchFieldDeclarationContext) interface{}

	// Visit a parse tree produced by PacketDslParser#matchPair.
	VisitMatchPair(ctx *MatchPairContext) interface{}

	// Visit a parse tree produced by PacketDslParser#inerObjectDeclaration.
	VisitInerObjectDeclaration(ctx *InerObjectDeclarationContext) interface{}

	// Visit a parse tree produced by PacketDslParser#list.
	VisitList(ctx *ListContext) interface{}
}
