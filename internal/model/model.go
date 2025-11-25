package model

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	// StringPrefixLenType prefix length type for dynamic string
	StringPrefixLenType = "StringPrefixLenType"
	// ArrayPrefixLenType prefix length type for dynamic array(list)
	ArrayPrefixLenType = "ArrayPrefixLenType"
	// LittleEndian option for endianness
	LittleEndian = "LittleEndian"
	// JavaPackage java package name
	JavaPackage = "JavaPackage"
	// GoPackage go package name
	GoPackage = "GoPackage"
	// GoModule go module name
	GoModule = "GoModule"
	// FixedStringPadFromLeft pad from left if true, otherwise pad from right
	FixedStringPadFromLeft = "FixedStringPadFromLeft"
	// FixedStringPadChar pad char, default is space
	FixedStringPadChar = "FixedStringPadChar"
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
	"zchar[]": {},
	"string":  {},
}

var options = map[string][]string{
	StringPrefixLenType:    {"u8", "u16", "u32", "u64"},
	ArrayPrefixLenType:     {"u8", "u16", "u32", "u64"},
	LittleEndian:           {"true", "false"},
	JavaPackage:            {},
	GoPackage:              {},
	GoModule:               {},
	FixedStringPadFromLeft: {"true", "false"},
	FixedStringPadChar:     {"'0'", "' '", "'\x00'"},
}

// BinaryModel contains metaData, options,and packets
type BinaryModel struct {
	MetaDataMap  map[string]MetaData // Map of metadata definitions
	Options      map[string]string   // Map of options
	PacketsMap   map[string]*Packet  // Map of packet definitions, keyed by packet name
	Packets      []*Packet           // List of all packets
	RootPacket   *Packet             // root packet
	SyntaxErrors []*SyntaxError      // List of syntax errors encountered during parsing
}

// NewBinaryModel new BinaryModel
func NewBinaryModel() *BinaryModel {
	return &BinaryModel{
		MetaDataMap: make(map[string]MetaData),
		Options:     make(map[string]string),
		PacketsMap:  make(map[string]*Packet),
	}
}

// ResolveDependencies after parse
func (m *BinaryModel) ResolveDependencies() {
	for _, packet := range m.Packets {
		for _, field := range packet.Fields {
			if of, ok := field.Attr.(*ObjectFieldAttribute); ok {
				if of.RefPacket == nil {
					if refPacket, exists := m.PacketsMap[of.PacketName]; exists {
						of.RefPacket = refPacket
					} else {
						m.AddSyntaxError(&SyntaxError{
							Line:   field.Line,
							Column: field.Column,
							Msg:    "Unknown packet type " + of.PacketName + " for field " + field.Name,
						})
					}
				}
			}
		}
	}
}

// AddMetaData add MetaData
func (m *BinaryModel) AddMetaData(metaData MetaData) {
	if _, exists := m.MetaDataMap[metaData.Name]; exists {
		m.AddSyntaxError(&SyntaxError{
			Line:   metaData.Line,
			Column: metaData.Column,
			Msg:    "Duplicate metadata definition for " + metaData.Name,
		})
		return
	}
	if m.MetaDataMap[metaData.Name] != (MetaData{}) {
		metaData.Attr = m.MetaDataMap[metaData.Name].Attr
	}
	m.MetaDataMap[metaData.Name] = metaData
}

// AddOption add option
func (m *BinaryModel) AddOption(name, value string, line int, column int) {
	if values, ok := options[name]; ok {
		if len(values) > 0 && !contains(values, value) {
			m.AddSyntaxError(&SyntaxError{line, column, "Option " + name + " is not allowed to be " + value + ", Expected one of:" + strings.Join(values, ","), nil})
		}
	} else {
		var keys []string
		for key := range options {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		m.AddSyntaxError(&SyntaxError{line, column, "Option " + name + " is not allowed in this context, Expected one of:" + strings.Join(keys, ","), nil})
		return
	}

	if _, ok := m.Options[name]; ok {
		m.AddSyntaxError(&SyntaxError{line, column, "Option " + name + " is already defined", nil})
		return
	}
	m.Options[name] = value
}

// AddSyntaxError add syntax error
func (m *BinaryModel) AddSyntaxError(error *SyntaxError) {
	m.SyntaxErrors = append(m.SyntaxErrors, error)
}

// AddPacket add packet
func (m *BinaryModel) AddPacket(packet *Packet) {
	if _, exists := m.PacketsMap[packet.Name]; exists {
		m.AddSyntaxError(&SyntaxError{
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
			m.AddSyntaxError(&SyntaxError{
				Line:   packet.Line,   // Assuming the first field's line is the packet's line
				Column: packet.Column, // Column is not defined in Packet
				Msg:    "Multiple root packets are not allowed",
			})
			return
		}
		m.RootPacket = packet
	}
}

// MetaData metaData
type MetaData struct {
	Name        string         // Name of the metadata
	Attr        FieldAttribute // Type attribute
	Description string         // Description of the metadata
	Line        int            // Line number where the metadata is defined
	Column      int            // Column number where the metadata is defined
}

// Packet represents a message definition, which may be a root or nested packet.
type Packet struct {
	Name        string                 // Name of the packet
	IsRoot      bool                   // True if the packet declaration included the 'root' keyword
	LengthField *Field                 // Length field for root packet (only one allowed and only applies to root packet)
	Fields      []*Field               // List of fields belonging to this packet
	FieldMap    map[string]*Field      // Map of fields belonging to this packet
	MatchFields map[string][]MatchPair // Map of match field names to their key-value pairs
	Line        int                    // Line number where the packet is defined
	Column      int                    // Column number where the packet is defined
}

// Padding represents padding configuration for a field.
type Padding struct {
	PadChar string // Character used for padding
	PadLeft bool   // True if padding is applied from the left, false if from the right
}

// IsDefault checks if the padding configuration is the default (space character, pad from right)
func (p Padding) IsDefault() bool {
	return p.PadChar == "' '" && !p.PadLeft
}

// FieldAttribute interface for field attributes
type FieldAttribute interface {
	GetType() string
}

// BasicFieldAttribute basic field attribute
type BasicFieldAttribute struct {
	Type string
}

// GetType return field type
func (b BasicFieldAttribute) GetType() string {
	return getBasicType(b.Type)
}

// LengthFieldAttribute length field attribute
type LengthFieldAttribute struct {
	TragetField *Field
	LengthType  string
}

// GetType return field type
func (l LengthFieldAttribute) GetType() string {
	return getBasicType(l.LengthType)
}

// LengthOfAttribute describes how a length field is computed.
//
// This attribute defines that `LengthField` stores the size of `TargetField`.
// The final length is calculated as: len(TargetField) + Offset.
type LengthOfAttribute struct {
	// TargetField is the field whose serialized size will be measured.
	TargetField *Field

	// LengthField is the field that stores the computed length value.
	LengthField *Field

	// Offset is an additional constant added to the computed length, default 0.
	// FinalLength = len(TargetField) + Offset.
	Offset int32
}

// GetType return field type
func (l LengthOfAttribute) GetType() string {
	return l.TargetField.Attr.GetType()
}

// CheckSumFieldAttribute check sum field attribute
type CheckSumFieldAttribute struct {
	CheckSumType string
	Type         string
}

// GetType return field type
func (c CheckSumFieldAttribute) GetType() string {
	return getBasicType(c.Type)
}

// FixedStringFieldAttribute fixed string field attribute
type FixedStringFieldAttribute struct {
	Length  int
	Padding *Padding
}

// GetType return field type
func (f FixedStringFieldAttribute) GetType() string {
	return "string"
}

// DynamicStringFieldAttribute dynamic string field attribute
type DynamicStringFieldAttribute struct {
	PrefixLenType string
	Type          string
}

// GetType return field type
func (d DynamicStringFieldAttribute) GetType() string {
	return "string"
}

// ObjectFieldAttribute object field attribute
type ObjectFieldAttribute struct {
	IsIner     bool
	PacketName string
	RefPacket  *Packet
}

// GetType return field type
func (o ObjectFieldAttribute) GetType() string {
	return "object"
}

// MatchFieldAttribute match field attribute
type MatchFieldAttribute struct {
	MatchKeyField *Field      // Reference to the match key field
	MatchPairs    []MatchPair // If the field is a match field, holds all match key-value pairs
}

// GetType return field type
func (m MatchFieldAttribute) GetType() string {
	return "match"
}

func getBasicType(fieldType string) string {
	switch strings.ToLower(fieldType) {
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
	default:
		return fieldType
	}
}

// Field represents a single field within a packet. It can be a basic field, nested object, or match field.
type Field struct {
	Name     string         // Field name
	Attr     FieldAttribute // Field attribute interface
	LenAttr  FieldAttribute // Length attribute or LengthOf attribute
	IsRepeat bool           // True if the 'repeat' modifier is present
	Doc      string         // Optional documentation string (from STRING_LITERAL), currently unused
	Tag      int            // Tag value for step or fix protocols
	Line     int            // Line number where the field is defined
	Column   int            // Column number where the field is defined
}

// ParseZCharArrayType parse char[\d]
func ParseZCharArrayType(fieldType string) (size int, ok bool) {
	pattern := `^zchar\[(\d+)\]$`
	matches := regexp.MustCompile(pattern).FindStringSubmatch(fieldType)
	if len(matches) == 2 {
		//must is number
		result, _ := strconv.Atoi(matches[1])
		return result, true
	}
	return 0, false
}

// ParseCharArrayType parse char[\d]
func ParseCharArrayType(fieldType string) (size int, ok bool) {
	pattern := `^char\[(\d+)\]$`
	matches := regexp.MustCompile(pattern).FindStringSubmatch(fieldType)
	if len(matches) == 2 {
		//must is number
		result, _ := strconv.Atoi(matches[1])
		return result, true
	}
	return 0, false
}

// GetType return field type
func (f Field) GetType() string {
	switch c := f.Attr.(type) {
	case *FixedStringFieldAttribute, *DynamicStringFieldAttribute:
		return "string"
	case *ObjectFieldAttribute:
		return c.RefPacket.Name
	case *MatchFieldAttribute:
		return "match"
	default:
		switch strings.ToLower(f.Attr.GetType()) {
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
			return f.Attr.GetType()
		}
	}
}

// MatchPair represents a single key-value mapping inside a match field.
type MatchPair struct {
	Key    string // Matching key literal, e.g. "0", "1", or "[A,B]"
	Value  string // Corresponding field name for this match key
	Line   int    // Line number where the match pair is defined
	Column int    // Column number where the match pair is defined
}
