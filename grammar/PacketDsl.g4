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
	'MetaData' IDENTIFIER '{' metaDataDeclaration+ '}';

// Metadata declaration with type and description
metaDataDeclaration:
	type IDENTIFIER ':' '`' description '`';

// Types for fields
type:
	'Int64'
	| 'uInt16'
	| 'uInt32'
	| 'string'
	| 'Int32'
	| 'uInt8'
	| 'char'
	| 'char[' DIGIT* ']';

// Descriptions are anything that doesn't contain backticks
description: ~('`')+;

// Root and Packet keywords
ROOT: 'ROOT';
PACKET: 'packet';

// Identifier matching letters and numbers (for identifiers like field names, types, etc.)
IDENTIFIER: [a-zA-Z_][a-zA-Z_0-9]*;

// Digits for handling arrays and other numeric types
DIGIT: [0-9];

// Match field rule for defining match criteria
matchField:
	'match' IDENTIFIER '{' ('"' IDENTIFIER '"')+ ': ' IDENTIFIER '}';

// Skip whitespace (spaces, tabs, newlines)
WS: [ \t\r\n]+ -> skip;