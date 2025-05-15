grammar PacketDsl;

packet: (packetDefinition | metaDataDefinition)*;

// Root rule for packet definition
packetDefinition:
	ROOT? PACKET IDENTIFIER '{' fieldDefinition* '}';

// Field definitions are either identifiers, metadata declarations, or match fields
fieldDefinition:
	IDENTIFIER
	| metaDataDeclaration
	| matchField;

// MetaData rule with declarations inside curly braces
metaDataDefinition:
	'MetaData' IDENTIFIER '{' metaDataDeclaration* '}';

// Metadata declaration with type and description
metaDataDeclaration:
    type? IDENTIFIER STRING_LITERAL? COMMA?;
// Types for fields
type:
	IDENTIFIER 
	| 'Int64'
	| 'uInt16'
	| 'uInt32'
	| 'string'
	| 'Int32'
	| 'uInt8'
	| 'char'
	| 'char[' DIGITS? ']';

// Match field rule for defining match criteria
matchField
    : 'match' IDENTIFIER '{' matchPair (COMMA matchPair)* '}'
    ;

matchPair
    : (DIGITS | STRING | list) COLON IDENTIFIER COMMA?
    ;

list
    : '[' (DIGITS | STRING) (COMMA (DIGITS | STRING))* ']'
    ;


DIGITS: [0-9]+;

STRING: '"' (~["\\\r\n])* '"';

// Root and Packet keywords
ROOT: 'root';
PACKET: 'packet';

// COLON定义，匹配英文和中文冒号
COLON: ':' | '：';

// COMMA匹配逗号
COMMA: ',';

// Identifier matching letters and numbers (for identifiers like field names, types, etc.)
IDENTIFIER: [a-zA-Z_][a-zA-Z_0-9]*;


STRING_LITERAL: '`' (~'`'|'\r'|'\n')* '`';


LINE_COMMENT: '//' ~[\r\n]* -> skip;
// Skip whitespace (spaces, tabs, newlines)
WS: [ \t\r\n]+ -> skip;