grammar PacketDsl;

packet: (
		packetDefinition
		| metaDataDefinition
		| optionDefinition
	)*;

optionDefinition: 'options' '{' optionDeclaration* '}';

optionDeclaration: IDENTIFIER '=' value SEMICOLON?;

// Root rule for packet definition
packetDefinition:
	ROOT? PACKET IDENTIFIER '{' fieldDefinition* '}';

// Field definitions are either identifiers, metadata declarations, or match fields
fieldDefinitionWithAttribute: fieldAttribute* fieldDefinition;

fieldDefinition:
	REPEAT? inerObjectDeclaration COMMA	# InerObjectField
	| REPEAT? metaDataDeclaration		# MetaField
	| REPEAT? IDENTIFIER COMMA			# ObjectField
	| lengthFieldDeclaration			# LengthField
	| checkSumFieldDeclaration			# CheckSumField
	| matchFieldDeclaration COMMA		# MatchField;

// MetaData rule with declarations inside curly braces
metaDataDefinition:
	METADATA IDENTIFIER '{' metaDataDeclaration* '}';

lengthFieldDeclaration:
	type? name = IDENTIFIER lengthOfAttribute STRING_LITERAL? COMMA;

checkSumFieldDeclaration:
	type? name = IDENTIFIER (calculatedFromAttribute) STRING_LITERAL? COMMA;

fieldAttribute: lengthOfAttribute | calculatedFromAttribute;

lengthOfAttribute: '@lengthOf(' from = IDENTIFIER ')';

calculatedFromAttribute: '@calculatedFrom(' from = STRING ')';

// Metadata declaration with type and description
metaDataDeclaration:
	type? name = IDENTIFIER STRING_LITERAL? COMMA;

value: type | STRING | DIGITS;

// Types for fields
type:
	IDENTIFIER
	| UINT8
	| UINT16
	| UINT32
	| UINT64
	| INT8
	| INT16
	| INT32
	| INT64
	| FLOAT32
	| FLOAT64
	| 'string'
	| 'char'
	| 'char[' DIGITS? ']';

// Match field rule for defining match criteria
matchFieldDeclaration:
	MATCH matchKey = IDENTIFIER 'as' matchName = IDENTIFIER '{' matchPair+ '}';

matchPair: (DIGITS | STRING | list) COLON IDENTIFIER COMMA?;

inerObjectDeclaration: IDENTIFIER ('{' fieldDefinition+ '}');

list: '[' (DIGITS | STRING) (COMMA (DIGITS | STRING))* ']';

UINT8: 'uint8' | 'u8';
UINT16: 'uint16' | 'u16';
UINT32: 'uint32' | 'u32';
UINT64: 'uint64' | 'u64';

INT8: 'int8' | 'i8';
INT16: 'int16' | 'i16';
INT32: 'int32' | 'i32';
INT64: 'int64' | 'i64';

FLOAT32: 'float32' | 'f32';

FLOAT64: 'float64' | 'f64';

DIGITS: [0-9]+;

STRING: '"' ( ~["\\\r\n] | '\\' .)* '"';

// Root and Packet keywords
ROOT: 'root';
PACKET: 'packet';
REPEAT: 'repeat';
METADATA: 'MetaData';
MATCH: 'match';

// COLON定义，匹配英文
COLON: ':';

// COMMA匹配逗号
COMMA: ',';

//分号
SEMICOLON: ';';

// Identifier matching letters and numbers (for identifiers like field names, types, etc.)
IDENTIFIER: [a-zA-Z_][a-zA-Z_0-9]*;

STRING_LITERAL: '`' (~'`' | '\r' | '\n')* '`';

LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
// Skip whitespace (spaces, tabs, newlines)
WS: [ \t\r\n]+ -> skip;