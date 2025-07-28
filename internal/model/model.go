package model

import (
	"sort"
	"strings"
)

func contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// SyntaxError represents a single syntax error found during parsing
type SyntaxError struct {
	Line            int
	Column          int
	Msg             string
	OffendingSymbol interface{}
}

var _ = map[string]struct{}{
	"u8":      {},
	"u16":     {},
	"u32":     {},
	"u64":     {},
	"i8":      {},
	"i16":     {},
	"i32":     {},
	"i64":     {},
	"f32":     {},
	"f64":     {},
	"uint8":   {},
	"uint16":  {},
	"uint32":  {},
	"uint64":  {},
	"int8":    {},
	"int16":   {},
	"int32":   {},
	"int64":   {},
	"float32": {},
	"float64": {},
	"char":    {},
	"char[]":  {},
	"string":  {},
}

var options = map[string][]string{
	"StringPreFixLenType":  {"u8", "u16", "u32", "u64"},
	"RepeatPreFixSizeType": {"u8", "u16", "u32", "u64"},
	"LittleEndian":         {"true", "false"},
	"JavaPackage":          {},
	"GoPackage":            {},
	"GoModule":             {},
}

// BinaryModel contains metaData, options,and packets
type BinaryModel struct {
	MetaDataMap  map[string]MetaData // Map of metadata definitions
	Options      map[string]string   // Map of options
	PacketsMap   map[string]Packet   // Map of packet definitions, keyed by packet name
	Packets      []Packet            // List of all packets
	RootPacket   *Packet             // root packet
	SyntaxErrors []SyntaxError       // List of syntax errors encountered during parsing
}

// NewBinaryModel new BinaryModel
func NewBinaryModel() *BinaryModel {
	return &BinaryModel{
		MetaDataMap: make(map[string]MetaData),
		Options:     make(map[string]string),
		PacketsMap:  make(map[string]Packet),
	}
}

// AddMetaData add MetaData
func (m *BinaryModel) AddMetaData(metaData MetaData) {
	if _, exists := m.MetaDataMap[metaData.Name]; exists {
		m.AddSyntaxError(SyntaxError{
			Line:   metaData.Line,
			Column: metaData.Column,
			Msg:    "Duplicate metadata definition for " + metaData.Name,
		})
		return
	}
	if m.MetaDataMap[metaData.Typ] != (MetaData{}) {
		metaData.BasicType = m.MetaDataMap[metaData.Typ].BasicType
	}
	m.MetaDataMap[metaData.Name] = metaData
}

// AddOption add option
func (m *BinaryModel) AddOption(name, value string, line int, column int) {
	if values, ok := options[name]; ok {
		if !contains(values, value) {
			m.AddSyntaxError(SyntaxError{line, column, "Option " + name + " is not allowed to be " + value + ", Expected one of:" + strings.Join(values, ","), nil})
		}
	} else {
		var keys []string
		for key := range options {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		m.AddSyntaxError(SyntaxError{line, column, "Option " + name + " is not allowed in this context, Expected one of:" + strings.Join(keys, ","), nil})
		return
	}

	if _, ok := m.Options[name]; ok {
		m.AddSyntaxError(SyntaxError{line, column, "Option " + name + " is already defined", nil})
		return
	}
	m.Options[name] = value
}

// AddSyntaxError add syntax error
func (m *BinaryModel) AddSyntaxError(error SyntaxError) {
	m.SyntaxErrors = append(m.SyntaxErrors, error)
}

// AddPacket add packet
func (m *BinaryModel) AddPacket(packet Packet) {
	if _, exists := m.PacketsMap[packet.Name]; exists {
		m.AddSyntaxError(SyntaxError{
			Line:   packet.Line,
			Column: packet.Column,
			Msg:    "Duplicate packet definition for " + packet.Name,
		})
		return
	}
	m.PacketsMap[packet.Name] = packet
	m.Packets = append(m.Packets, packet)
	if packet.IsRoot {
		if m.RootPacket != nil {
			m.AddSyntaxError(SyntaxError{
				Line:   packet.Line,   // Assuming the first field's line is the packet's line
				Column: packet.Column, // Column is not defined in Packet
				Msg:    "Multiple root packets are not allowed",
			})
			return
		}
		m.RootPacket = &packet
	}
}

// MetaData metaData
type MetaData struct {
	Name        string // Name of the metadata
	Typ         string // Type of the metadata,
	BasicType   string // Basic type if applicable, e.g., "i32", "u16"
	Description string // Description of the metadata
	Line        int    // Line number where the metadata is defined
	Column      int    // Column number where the metadata is defined
}

// Packet represents a message definition, which may be a root or nested packet.
type Packet struct {
	Name          string                 // Name of the packet
	IsRoot        bool                   // True if the packet declaration included the 'root' keyword
	LengthOfField string                 // Length field name for root packet (only one allowed and only applies to root packet)
	Fields        []Field                // List of fields belonging to this packet
	FieldMap      map[string]Field       // Map of fields belonging to this packet
	MatchFields   map[string][]MatchPair // Map of match field names to their key-value pairs
	Line          int                    // Line number where the packet is defined
	Column        int                    // Column number where the packet is defined
}

// Field represents a single field within a packet. It can be a basic field, nested object, or match field.
type Field struct {
	Name          string      // Field name
	Type          string      // Type name if this is a basic field; empty for nested or match fields
	LengthOfField string      // This is a length field, its value is assigned from the length of another field (LengthOfField), usually used in root packets.
	CheckSumType  string      //
	IsRepeat      bool        // True if the 'repeat' modifier is present
	InerObject    *Packet     // If the field is a nested object, this holds the nested Packet definition
	Doc           string      // Optional documentation string (from STRING_LITERAL), currently unused
	MatchKey      string      // If the field is a match field, this holds the typeName of match
	MatchPairs    []MatchPair // If the field is a match field, holds all match key-value pairs
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
	Key    string // Matching key literal, e.g. "0", "1", or "[A,B]"
	Value  string // Corresponding field name for this match key
	Line   int    // Line number where the match pair is defined
	Column int    // Column number where the match pair is defined
}
