package model

import "strings"

// BinaryModel contains metaData, options,and packets
type BinaryModel struct {
	MetaDataMap map[string]MetaData // Map of metadata definitions
	Options     map[string]string   // Map of options
	Packets     map[string]Packet   // Map of packet definitions, keyed by packet name
	RootPacket  Packet              // root packet
}

// NewBinaryModel new BinaryModel
func NewBinaryModel() *BinaryModel {
	return &BinaryModel{
		MetaDataMap: make(map[string]MetaData),
		Options:     make(map[string]string),
		Packets:     make(map[string]Packet),
	}
}

// AddMetaData add MetaData
func (m *BinaryModel) AddMetaData(metaData MetaData) {
	if m.MetaDataMap[metaData.Typ] != (MetaData{}) {
		metaData.BasicType = m.MetaDataMap[metaData.Typ].BasicType
	}
	m.MetaDataMap[metaData.Name] = metaData
}

// AddOption add option
func (m *BinaryModel) AddOption(name, value string) {
	m.Options[name] = value
}

// AddPacket add packet
func (m *BinaryModel) AddPacket(packet Packet) {
	m.Packets[packet.Name] = packet
	if packet.IsRoot {
		m.RootPacket = packet
	}
}

// MetaData metaData
type MetaData struct {
	Name        string // Name of the metadata
	Typ         string // Type of the metadata,
	BasicType   string // Basic type if applicable, e.g., "i32", "u16"
	Description string // Description of the metadata
}

// Packet represents a message definition, which may be a root or nested packet.
type Packet struct {
	Name        string                 // Name of the packet
	IsRoot      bool                   // True if the packet declaration included the 'root' keyword
	Fields      []Field                // List of fields belonging to this packet
	MatchFields map[string][]MatchPair // Map of match field names to their key-value pairs
}

// Field represents a single field within a packet. It can be a basic field, nested object, or match field.
type Field struct {
	Name       string      // Field name
	Type       string      // Type name if this is a basic field; empty for nested or match fields
	IsRepeat   bool        // True if the 'repeat' modifier is present
	InerObject *Packet     // If the field is a nested object, this holds the nested Packet definition
	Doc        string      // Optional documentation string (from STRING_LITERAL), currently unused
	MatchType  string      // If the field is a match field, this holds the typeName of match
	MatchPairs []MatchPair // If the field is a match field, holds all match key-value pairs
}

// GetType return field type
func (f Field) GetType() string {
	if f.Type != "" {
		switch strings.ToLower(f.Type) {
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
		case "string", "char[]":
			return "string"
		default:

			return f.Type // Return the type as is if it doesn't match any known types
		}
	}
	if f.InerObject != nil && f.InerObject.Name != "" {
		return f.InerObject.Name
	} else if len(f.MatchPairs) > 0 {
		return "match"
	}
	return "unknown"
}

// MatchPair represents a single key-value mapping inside a match field.
type MatchPair struct {
	Key   string // Matching key literal, e.g. "0", "1", or "[A,B]"
	Value string // Corresponding field name for this match key
}
