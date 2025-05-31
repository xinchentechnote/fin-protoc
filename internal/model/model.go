package model

// Packet represents a message definition, which may be a root or nested packet.
type Packet struct {
	Name   string  // Name of the packet
	IsRoot bool    // True if the packet declaration included the 'root' keyword
	Fields []Field // List of fields belonging to this packet
}

// Field represents a single field within a packet. It can be a basic field, nested object, or match field.
type Field struct {
	Name       string      // Field name
	Type       string      // Type name if this is a basic field; empty for nested or match fields
	IsRepeat   bool        // True if the 'repeat' modifier is present
	InerObject Packet      // If the field is a nested object, this holds the nested Packet definition
	Doc        string      // Optional documentation string (from STRING_LITERAL), currently unused
	MatchPairs []MatchPair // If the field is a match field, holds all match key-value pairs
}

// MatchPair represents a single key-value mapping inside a match field.
type MatchPair struct {
	Key   string // Matching key literal, e.g. "0", "1", or "[A,B]"
	Value string // Corresponding field name for this match key
}
