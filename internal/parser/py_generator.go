package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// PyType represents a Python type for a field in a packet.
type PyType struct {
	Name         string
	BasicType    string
	TestValue    string
	DefaultValue string
	Le           string
	Size         uint32
}

// pyBasicTypeMap maps field types to their Python representations.
var pyBasicTypeMap = map[string]PyType{
	"i8":  {"i8", "i8", "1", "0", "i8", 1},
	"i16": {"i16", "i16", "2", "0", "i16_le", 2},
	"i32": {"i32", "i32", "4", "0", "i32_le", 4},
	"i64": {"i64", "i64", "8", "0", "i64_le", 8},
	"u8":  {"u8", "u8", "1", "0", "u8", 1},
	"u16": {"u16", "u16", "2", "0", "u16_le", 2},
	"u32": {"u32", "u32", "4", "0", "u32_le", 4},
	"u64": {"u64", "u64", "8", "0", "u64_le", 8},
	"f32": {"f32", "f32", "4", "0", "f32_le", 4},
	"f64": {"f64", "f64", "8", "0", "f64_le", 8},
}

// PythonGenerator a go code generator
type PythonGenerator struct {
	config   *GeneratorConfig
	binModel *model.BinaryModel
}

// NewPythonGenerator new
func NewPythonGenerator(config *GeneratorConfig, binModel *model.BinaryModel) *PythonGenerator {
	return &PythonGenerator{
		config:   config,
		binModel: binModel,
	}
}

// Generate python code
func (g PythonGenerator) Generate(binModel *model.BinaryModel) (map[string][]byte, error) {
	output := make(map[string][]byte)
	// Generate code for each packet
	var code = g.generateCode(binModel)
	output[strcase.ToSnake(binModel.RootPacket.Name)+".py"] = []byte(code)
	// Generate test code for each packet
	var testCode = g.generateTestCode(binModel)
	output[strcase.ToSnake(binModel.RootPacket.Name)+"_test.py"] = []byte(testCode)
	return output, nil
}

func (g PythonGenerator) generateCode(binModel *model.BinaryModel) string {
	var code strings.Builder
	code.WriteString("# Code generated by fin-protoc. DO NOT EDIT.\n")
	code.WriteString(`from bytebuf import ByteBuf
from codec import *`)
	code.WriteString("\n\n")
	for _, pkt := range binModel.PacketsMap {
		code.WriteString(g.generateCodeForPacket(&pkt))
		code.WriteString("\n\n")
	}
	return code.String()
}

func (g PythonGenerator) generateCodeForPacket(p *model.Packet) string {
	var b strings.Builder
	// gen iner object
	for _, f := range p.Fields {
		if f.InerObject != nil {
			b.WriteString(g.generateCodeForPacket(f.InerObject))
		}
	}
	// gen class
	b.WriteString(fmt.Sprintf("class %s(BinaryCodec):\n", strcase.ToCamel(p.Name)))
	// gen __init_ method
	b.WriteString(AddIndent4ln(g.generateInitMethod(p)))
	// gen encode method
	b.WriteString(AddIndent4ln(g.generateEncodeMethod(p)))
	// gen decode method
	b.WriteString(AddIndent4ln(g.generateDecodeMethod(p)))
	// gen __eq__ method
	b.WriteString(AddIndent4ln(g.generateEqMethod(p)))
	return b.String()
}

func (g PythonGenerator) generateEqMethod(p *model.Packet) string {
	var b strings.Builder
	b.WriteString("def __eq__(self, other):\n")
	b.WriteString("    if not isinstance(other, self.__class__):\n")
	b.WriteString("        return False\n")
	if len(p.Fields) == 0 {
		b.WriteString("    return True\n")
		return b.String()
	}
	b.WriteString("    return all([\n")
	for i, f := range p.Fields {
		if _, ok := pyBasicTypeMap[f.GetType()]; ok {
			b.WriteString(fmt.Sprintf("        self.%s == other.%s", strcase.ToSnake(f.Name), strcase.ToSnake(f.Name)))
		} else if _, ok := ParseCharArrayType(f.GetType()); ok {
			b.WriteString(fmt.Sprintf("        self.%s == other.%s", strcase.ToSnake(f.Name), strcase.ToSnake(f.Name)))
		} else if f.InerObject != nil || (f.GetType() != "string" && f.GetType() != "char[]") {
			b.WriteString(fmt.Sprintf("        self.%s == other.%s", strcase.ToSnake(f.Name), strcase.ToSnake(f.Name)))
		} else {
			b.WriteString(fmt.Sprintf("        self.%s == other.%s", strcase.ToSnake(f.Name), strcase.ToSnake(f.Name)))
		}
		if i < len(p.Fields)-1 {
			b.WriteString(",\n")
		} else {
			b.WriteString("\n")
		}
	}
	b.WriteString("    ])\n")
	b.WriteString("    \n")
	return b.String()
}

func (g PythonGenerator) generateDecodeMethod(p *model.Packet) string {
	var b strings.Builder
	b.WriteString("def decode(self, buffer: ByteBuf):\n")
	if len(p.Fields) == 0 {
		b.WriteString("    pass\n")
		return b.String()
	}
	for _, f := range p.Fields {
		if f.IsRepeat {
			if g.config.LittleEndian {
				b.WriteString(fmt.Sprintf("    size = get_len_le(buffer, '%s')\n", g.config.ListLenPrefixLenType))
			} else {
				b.WriteString(fmt.Sprintf("    size = get_len(buffer, '%s')\n", g.config.ListLenPrefixLenType))
			}
			b.WriteString("    for i in range(size):\n")
			b.WriteString(AddIndent4ln(g.generateDecodeField(&f)))
		} else {
			b.WriteString(g.generateDecodeField(&f))
		}
	}
	return b.String()
}

func (g PythonGenerator) generateDecodeField(f *model.Field) string {
	var b strings.Builder
	if typ, ok := pyBasicTypeMap[f.GetType()]; ok {
		read := typ.BasicType
		if g.config.LittleEndian {
			read = typ.Le
		}
		if f.IsRepeat {
			b.WriteString(fmt.Sprintf("    self.%s.append(buffer.read_%s())\n", strcase.ToSnake(f.Name), read))
		} else {
			b.WriteString(fmt.Sprintf("    self.%s = buffer.read_%s()\n", strcase.ToSnake(f.Name), read))
		}
	}
	if size, ok := ParseCharArrayType(f.GetType()); ok {
		if f.IsRepeat {
			b.WriteString(fmt.Sprintf("    self.%s.append(buffer.read_bytes(%s).decode('utf-8').strip('\\x00'))\n", strcase.ToSnake(f.Name), size))
		} else {
			b.WriteString(fmt.Sprintf("    self.%s = buffer.read_bytes(%s).decode('utf-8').strip('\\x00')\n", strcase.ToSnake(f.Name), size))
		}
	}
	if f.GetType() == "string" || f.GetType() == "char[]" {
		var le string
		if g.config.LittleEndian {
			le = "_le"
		}
		if f.IsRepeat {
			b.WriteString(fmt.Sprintf("    self.%s.append(get_string%s(buffer,'%s'))\n", strcase.ToSnake(f.Name), le, g.config.StringLenPrefixLenType))
		} else {
			b.WriteString(fmt.Sprintf("    self.%s = get_string%s(buffer,'%s')\n", strcase.ToSnake(f.Name), le, g.config.StringLenPrefixLenType))
		}
	}
	if f.InerObject != nil {
		if f.IsRepeat {
			b.WriteString(fmt.Sprintf("    _%s = %s()\n", strcase.ToSnake(f.Name), strcase.ToCamel(f.InerObject.Name)))
			b.WriteString(fmt.Sprintf("    _%s.decode(buffer)\n", strcase.ToSnake(f.Name)))
			b.WriteString(fmt.Sprintf("    self.%s.append(_%s)\n", strcase.ToSnake(f.Name), strcase.ToSnake(f.Name)))
		} else {
			b.WriteString(fmt.Sprintf("    self.%s = %s()\n", strcase.ToSnake(f.Name), strcase.ToCamel(f.InerObject.Name)))
			b.WriteString(fmt.Sprintf("    self.%s.decode(buffer)\n", strcase.ToSnake(f.Name)))
		}
	}
	if _, ok := g.binModel.PacketsMap[f.GetType()]; ok {
		if f.IsRepeat {
			b.WriteString(fmt.Sprintf("    _%s = %s()\n", strcase.ToSnake(f.GetType()), strcase.ToCamel(f.GetType())))
			b.WriteString(fmt.Sprintf("    _%s.decode(buffer)\n", strcase.ToSnake(f.GetType())))
			b.WriteString(fmt.Sprintf("    self.%s.append(_%s)\n", strcase.ToSnake(f.Name), strcase.ToSnake(f.GetType())))
		} else {
			b.WriteString(fmt.Sprintf("    self.%s = %s()\n", strcase.ToSnake(f.Name), strcase.ToCamel(f.GetType())))
			b.WriteString(fmt.Sprintf("    self.%s.decode(buffer)\n", strcase.ToSnake(f.Name)))
		}
	}
	if f.GetType() == "match" {
		for _, pair := range f.MatchPairs {
			b.WriteString(fmt.Sprintf("    if self.%s == %s:\n", strcase.ToSnake(f.MatchKey), pair.Key))
			b.WriteString(fmt.Sprintf("        self.%s = %s()\n", strcase.ToSnake(f.Name), pair.Value))
		}
		b.WriteString(fmt.Sprintf("    self.%s.decode(buffer)\n", strcase.ToSnake(f.Name)))
	}
	return b.String()
}

func (g PythonGenerator) generateEncodeMethod(p *model.Packet) string {
	var b strings.Builder
	b.WriteString("def encode(self, buffer: ByteBuf):\n")
	if len(p.Fields) == 0 {
		b.WriteString("    pass\n")
		return b.String()
	}

	for _, f := range p.Fields {
		if f.LengthOfField != "" {
			b.WriteString(fmt.Sprintf("    %s_buf = ByteBuf()\n", strcase.ToSnake(f.LengthOfField)))
			b.WriteString(fmt.Sprintf("    self.%s.encode(%s_buf)\n", strcase.ToSnake(f.LengthOfField), strcase.ToSnake(f.LengthOfField)))
			b.WriteString(fmt.Sprintf("    self.%s = %s_buf.readable_bytes_len()\n", strcase.ToSnake(f.Name), strcase.ToSnake(f.LengthOfField)))
		}
		if f.Name == p.LengthOfField {
			b.WriteString(fmt.Sprintf("    buffer.write_bytes(%s_buf.to_bytes())\n", strcase.ToSnake(f.Name)))
			continue
		}
		if f.IsRepeat {
			typ := pyBasicTypeMap[g.config.ListLenPrefixLenType]
			b.WriteString(fmt.Sprintf("    size = len(self.%s)\n", strcase.ToSnake(f.Name)))
			if g.config.LittleEndian {
				b.WriteString(fmt.Sprintf("    buffer.write_%s(size)\n", typ.Le))
			} else {
				b.WriteString(fmt.Sprintf("    buffer.write_%s(size)\n", typ.BasicType))
			}
			b.WriteString("    for i in range(size):\n")
			b.WriteString(AddIndent4ln(g.generateEncodeField(&f)))
		} else {
			b.WriteString(g.generateEncodeField(&f))
		}
	}
	return b.String()
}

func (g PythonGenerator) generateEncodeField(f *model.Field) string {
	var b strings.Builder
	fieldName := strcase.ToSnake(f.Name)
	if f.IsRepeat {
		fieldName += "[i]"
	}

	if typ, ok := pyBasicTypeMap[f.GetType()]; ok {
		if g.config.LittleEndian {
			b.WriteString(fmt.Sprintf("    buffer.write_%s(self.%s)\n", typ.Le, fieldName))
		} else {
			b.WriteString(fmt.Sprintf("    buffer.write_%s(self.%s)\n", typ.BasicType, fieldName))
		}
	} else if size, ok := ParseCharArrayType(f.GetType()); ok {
		b.WriteString(fmt.Sprintf("    write_fixed_string(buffer, self.%s, %s)\n", fieldName, size))
	} else if f.GetType() == "string" || f.GetType() == "char[]" {
		if g.config.LittleEndian {
			b.WriteString(fmt.Sprintf("    put_string_le(buffer, self.%s, '%s')\n", fieldName, g.config.StringLenPrefixLenType))
		} else {
			b.WriteString(fmt.Sprintf("    put_string(buffer, self.%s, '%s')\n", fieldName, g.config.StringLenPrefixLenType))
		}
	} else if f.InerObject != nil {
		b.WriteString(fmt.Sprintf("    self.%s.encode(buffer)\n", fieldName))
	} else if _, ok := g.binModel.PacketsMap[f.GetType()]; ok {
		b.WriteString(fmt.Sprintf("    self.%s.encode(buffer)\n", fieldName))
	} else if f.GetType() == "match" {
		b.WriteString(fmt.Sprintf("    if self.%s is not None:\n", fieldName))
		b.WriteString(fmt.Sprintf("        self.%s.encode(buffer)\n", fieldName))
	} else {
		b.WriteString("-- unsupported type: " + f.GetType() + "\n")
	}
	return b.String()
}

func (g PythonGenerator) generateInitMethod(p *model.Packet) string {
	var b strings.Builder
	b.WriteString("def __init__(self):\n")
	if len(p.Fields) == 0 {
		b.WriteString("    pass\n")
		return b.String()
	}
	for _, f := range p.Fields {
		if f.IsRepeat {
			b.WriteString(fmt.Sprintf("    self.%s = []\n", strcase.ToSnake(f.Name)))
			continue
		}
		if _, ok := pyBasicTypeMap[f.GetType()]; ok {
			b.WriteString(fmt.Sprintf("    self.%s = %s\n", strcase.ToSnake(f.Name), pyBasicTypeMap[f.GetType()].DefaultValue))
			continue
		}
		if _, ok := ParseCharArrayType(f.GetType()); ok {
			b.WriteString(fmt.Sprintf("    self.%s = ''\n", strcase.ToSnake(f.Name)))
			continue
		}
		switch f.GetType() {
		case "string", "char[]":
			b.WriteString(fmt.Sprintf("    self.%s = ''\n", strcase.ToSnake(f.Name)))
		case "match":
			b.WriteString(fmt.Sprintf("    self.%s = None\n", strcase.ToSnake(f.Name)))
		default:
			b.WriteString(fmt.Sprintf("    self.%s = None\n", strcase.ToSnake(f.Name)))
		}
	}
	return b.String()
}

func (g PythonGenerator) generateTestCode(binModel *model.BinaryModel) string {
	var b strings.Builder
	b.WriteString("# Code generated by fin-protoc. DO NOT EDIT.\n")
	b.WriteString("import unittest\n\n")
	b.WriteString(fmt.Sprintf("from %s import *\n\n", strcase.ToSnake(binModel.RootPacket.Name)))
	for _, pkt := range binModel.PacketsMap {
		b.WriteString(g.generateTestCodeForPacket(&pkt))
		b.WriteString("\n\n")
	}
	b.WriteString("if __name__ == '__main__':\n")
	b.WriteString("    unittest.main()\n")
	return b.String()
}

func (g PythonGenerator) generateTestCodeForPacket(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("class Test%s(unittest.TestCase):\n", strcase.ToCamel(packet.Name)))
	b.WriteString("    def setUp(self):\n")
	b.WriteString(AddIndent4ln(AddIndent4(g.generateNewInstance("self.packet", packet))))
	b.WriteString("\n")
	b.WriteString("    def test_encode_decode(self):\n")
	b.WriteString("        buf = ByteBuf()\n")
	b.WriteString("        self.packet.encode(buf)\n")
	b.WriteString(fmt.Sprintf("        decoded_packet = %s()\n", strcase.ToCamel(packet.Name)))
	b.WriteString("        decoded_packet.decode(buf)\n")
	b.WriteString("        self.assertEqual(decoded_packet, self.packet)\n")
	b.WriteString("\n")
	return b.String()
}

func (g PythonGenerator) generateNewInstance(name string, packet *model.Packet) string {
	var b strings.Builder
	for _, f := range packet.Fields {
		if f.InerObject != nil {
			b.WriteString(g.generateNewInstance(strcase.ToSnake(f.InerObject.Name), f.InerObject))
		}
		if rp, ok := g.binModel.PacketsMap[f.GetType()]; ok {
			b.WriteString(g.generateNewInstance(strcase.ToSnake(f.Name), &rp))
		}
		if f.GetType() == "match" {
			mp := g.binModel.PacketsMap[f.MatchPairs[0].Value]
			b.WriteString(g.generateNewInstance(strcase.ToSnake(f.Name), &mp))
		}
	}
	b.WriteString(fmt.Sprintf("%s = %s()\n", name, strcase.ToCamel(packet.Name)))
	for _, f := range packet.Fields {
		if packet.MatchFields[f.MatchKey] != nil {
			if len(f.MatchPairs) > 0 {
				key := f.MatchPairs[0].Key
				b.WriteString(fmt.Sprintf("%s.%s = %s\n", name, strcase.ToSnake(f.MatchKey), key))
				b.WriteString(fmt.Sprintf("%s.%s = %s\n", name, strcase.ToSnake(f.Name), g.generateTestValue(&f)))
			}
			continue
		}
		if packet.MatchFields[f.Name] != nil {
			continue
		}
		b.WriteString(fmt.Sprintf("%s.%s = %s\n", name, strcase.ToSnake(f.Name), g.generateTestValue(&f)))
	}

	return b.String()
}

func (g PythonGenerator) generateTestValue(f *model.Field) string {
	var testValue string
	if typ, ok := pyBasicTypeMap[f.GetType()]; ok {
		testValue = typ.TestValue
	}
	if f.InerObject != nil {
		testValue = strcase.ToSnake(f.InerObject.Name)
	}
	if _, ok := g.binModel.PacketsMap[f.GetType()]; ok {
		testValue = strcase.ToSnake(f.Name)
	}
	if f.GetType() == "string" || f.GetType() == "char[]" {
		testValue = "\"hello\""
	}
	if size, ok := ParseCharArrayType(f.GetType()); ok {
		s, _ := strconv.Atoi(size)
		testValue = "\"" + strings.Repeat("x", s) + "\""
	}
	if f.GetType() == "match" {
		testValue = strcase.ToSnake(f.Name)
	}
	if f.IsRepeat {
		return "[" + testValue + "]"
	}
	return testValue
}
