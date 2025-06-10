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
fieldDefinition:
	REPEAT? inerObjectDeclaration COMMA?	# InerObjectField
	| REPEAT? metaDataDeclaration			# MetaField
	| REPEAT? IDENTIFIER COMMA?				# ObjectField
	| matchFieldDeclaration COMMA?			# MatchField;

// MetaData rule with declarations inside curly braces
metaDataDefinition:
	METADATA IDENTIFIER '{' metaDataDeclaration* '}';

// Metadata declaration with type and description
metaDataDeclaration: type? IDENTIFIER STRING_LITERAL? COMMA?;

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
matchFieldDeclaration: MATCH IDENTIFIER '{' matchPair+ '}';

matchPair: (DIGITS | STRING | list) COLON IDENTIFIER COMMA?;

inerObjectDeclaration: IDENTIFIER ('{' fieldDefinition+ '}');

list: '[' (DIGITS | STRING) (COMMA (DIGITS | STRING))* ']';

UINT8 options {
	caseInsensitive = true;
}: 'uint8' | 'u8';
UINT16 options {
	caseInsensitive = true;
}: 'uint16' | 'u16';
UINT32 options {
	caseInsensitive = true;
}: 'uint32' | 'u32';
UINT64 options {
	caseInsensitive = true;
}: 'uint64' | 'u64';

INT8 options {
	caseInsensitive = true;
}: 'int8' | 'i8';
INT16 options {
	caseInsensitive = true;
}: 'int16' | 'i16';
INT32 options {
	caseInsensitive = true;
}: 'int32' | 'i32';
INT64 options {
	caseInsensitive = true;
}: 'int64' | 'i64';

FLOAT32 options {
	caseInsensitive = true;
}: 'float32' | 'f32';

FLOAT64 options {
	caseInsensitive = true;
}: 'float64' | 'f64';

DIGITS: [0-9]+;

STRING: '"' (~["\\\r\n])* '"';

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