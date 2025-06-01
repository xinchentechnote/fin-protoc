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

// GetType return field type
func (f Field) GetType() string {
	if f.Type != "" {
		switch f.Type {
		case "i8", "int8":
			return "i8"
		case "i16", "int16":
			return "i16"
		case "i32", "int32":
			return "i32"
		case "i64", "int64":
			return "i64"
		case "u8", "uint8":
			return "u8"
		case "u16", "uint16":
			return "u16"
		case "u32", "uint32":
			return "u32"
		case "u64", "uint64":
			return "u64"
		case "f32", "float32":
			return "f32"
		case "f64", "float64":
			return "f64"
		case "string":
			return "string"
		default:
			return f.Type // Return the type as is if it doesn't match any known types
		}
	}
	if f.InerObject.Name != "" {
		return f.InerObject.Name
	}
	return "unknown"
}

// MatchPair represents a single key-value mapping inside a match field.
type MatchPair struct {
	Key   string // Matching key literal, e.g. "0", "1", or "[A,B]"
	Value string // Corresponding field name for this match key
}
