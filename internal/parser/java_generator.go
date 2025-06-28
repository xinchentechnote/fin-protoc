package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// JavaType java basic type
type JavaType struct {
	Name      string
	BasicType string
	BoxType   string
	TestValue string
	Le        string
	Size      uint32
}

// javaBasicTypeMap dsl basic type to java type
var javaBasicTypeMap = map[string]JavaType{
	// 无符号整数类型
	"u8": {
		Name:      "u8",
		BasicType: "byte",
		BoxType:   "Byte",
		Le:        "Byte",
		Size:      1,
		TestValue: "(byte)1",
	},
	"char": {
		Name:      "char",
		BasicType: "byte",
		BoxType:   "Byte",
		Le:        "Byte",
		Size:      1,
		TestValue: "(byte)1",
	},
	"u16": {
		Name:      "u16",
		BasicType: "short",
		BoxType:   "Short",
		Le:        "ShortLE",
		Size:      2,
		TestValue: "(short)2",
	},
	"u32": {
		Name:      "u32",
		BasicType: "int",
		BoxType:   "Integer",
		Le:        "IntLE",
		Size:      4,
		TestValue: "4",
	},
	"u64": {
		Name:      "u64",
		BasicType: "long",
		BoxType:   "Long",
		Le:        "LongLE",
		Size:      8,
		TestValue: "8L",
	},

	// 有符号整数类型
	"i8": {
		Name:      "i8",
		BasicType: "byte",
		BoxType:   "Byte",
		Le:        "Byte",
		Size:      1,
		TestValue: "(byte)1",
	},
	"i16": {
		Name:      "i16",
		BasicType: "short",
		BoxType:   "Short",
		Le:        "ShortLE",
		Size:      2,
		TestValue: "(short)2",
	},
	"i32": {
		Name:      "i32",
		BasicType: "int",
		BoxType:   "Integer",
		Le:        "IntLE",
		Size:      4,
		TestValue: "4",
	},
	"i64": {
		Name:      "i64",
		BasicType: "long",
		BoxType:   "Long",
		Le:        "LongLE",
		Size:      8,
		TestValue: "8L",
	},

	// 浮点类型
	"f32": {
		Name:      "f32",
		BasicType: "float",
		BoxType:   "Float",
		Le:        "FloatLE",
		Size:      4,
		TestValue: "4",
	},
	"f64": {
		Name:      "f64",
		BasicType: "double",
		BoxType:   "Double",
		Le:        "DoubleLE",
		Size:      8,
		TestValue: "8",
	},
}

// JavaGenerator a java code generator
type JavaGenerator struct {
	config   *GeneratorConfig
	binModel *model.BinaryModel
}

// NewJavaGenerator new
func NewJavaGenerator(config *GeneratorConfig, binModel *model.BinaryModel) *JavaGenerator {
	return &JavaGenerator{
		config:   config,
		binModel: binModel,
	}
}

// Generate gen hava code entry
func (g JavaGenerator) Generate(binModel *model.BinaryModel) (map[string][]byte, error) {
	output := make(map[string][]byte)
	packagePath := strings.ReplaceAll(g.config.JavaPackage, ".", "/")
	for _, pkg := range binModel.PacketsMap {
		//gen message class
		code := g.GenerateJavaClassFileForPacket(&pkg, false)
		filePath := fmt.Sprintf("main/java/%s/%s.java", packagePath, pkg.Name)
		output[filePath] = []byte(code)
		//gen unit test class
		code = g.GenerateJavaTestClassFileForPacket(&pkg)
		filePath = fmt.Sprintf("test/java/%s/%sTest.java", packagePath, pkg.Name)
		output[filePath] = []byte(code)
	}
	return output, nil
}

// GenerateJavaClassFileForPacket gen java class file
func (g JavaGenerator) GenerateJavaClassFileForPacket(packet *model.Packet, isInerClass bool) string {
	var b strings.Builder
	if !isInerClass {
		//package
		b.WriteString(fmt.Sprintf("package %s;\n", g.config.JavaPackage))
		//import
		b.WriteString(`import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.function.Supplier;
import com.finproto.codec.BinaryCodec;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.util.internal.StringUtil;
`)
		b.WriteString("\n")
		b.WriteString("\n")
	}
	static := ""
	if isInerClass {
		static = " static"
	}
	//class
	b.WriteString(fmt.Sprintf("public%s class %s implements BinaryCodec {\n", static, packet.Name))
	//field
	for _, f := range packet.Fields {
		b.WriteString(AddIndent4ln(fmt.Sprintf("private %s %s;", g.GetFieldType(&f), g.GetFieldNameLower(&f))))
	}
	b.WriteString("\n")

	//get and set
	for _, f := range packet.Fields {
		b.WriteString(AddIndent4ln(g.GenerateGetterAndSetter(&f)))
	}
	b.WriteString("\n")

	//encode
	b.WriteString(AddIndent4ln(g.GenerateEncode(packet)))
	b.WriteString("\n")
	//decode
	b.WriteString(AddIndent4ln(g.GenerateDecode(packet)))

	//create
	for _, f := range packet.Fields {
		if f.Type == "match" {
			b.WriteString(AddIndent4ln(g.GenerateCreateMethod(packet, &f)))
		}
	}

	b.WriteString("\n")
	//hashcode and equals
	b.WriteString(AddIndent4ln(g.GenerateHashCode(packet)))
	b.WriteString(AddIndent4ln(g.GenerateEquals(packet)))

	b.WriteString("\n")
	//tostring
	b.WriteString(AddIndent4ln(g.GenerateToString(packet)))
	//inerClass
	for _, f := range packet.Fields {
		if f.InerObject != nil {
			b.WriteString(AddIndent4ln(g.GenerateJavaClassFileForPacket(f.InerObject, true)))
		}
	}

	b.WriteString("\n")
	b.WriteString("}")
	return b.String()
}

// GenerateCreateMethod gen create method
func (g JavaGenerator) GenerateCreateMethod(p *model.Packet, f *model.Field) string {
	var b strings.Builder
	typ := p.FieldMap[f.MatchKey].GetType()
	matchKeyTyp, ok := javaBasicTypeMap[typ]
	cast := ""
	if ok {
		typ = matchKeyTyp.BoxType
		cast = fmt.Sprintf("(%s) ", matchKeyTyp.BasicType)
	}
	_, ok = ParseCharArrayType(typ)
	if ok {
		typ = "String"
	}
	b.WriteString(fmt.Sprintf("private static final Map<%s, Supplier<BinaryCodec>> %sMap = new HashMap<>();", typ, g.GetFieldNameLower(f)))
	b.WriteString("\n")
	b.WriteString("static {\n")
	for _, pair := range f.MatchPairs {
		b.WriteString(fmt.Sprintf("    %sMap.put(%s%s, %s::new);\n", g.GetFieldNameLower(f), cast, pair.Key, pair.Value))
	}
	b.WriteString("}\n")
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("private BinaryCodec create%s(%s %s) {\n", g.GetFieldName(f), typ, strcase.ToLowerCamel(f.MatchKey)))
	b.WriteString(fmt.Sprintf("    Supplier<BinaryCodec> supplier = %sMap.get(%s);\n", g.GetFieldNameLower(f), strcase.ToLowerCamel(f.MatchKey)))
	b.WriteString("    if (null == supplier) {\n")
	b.WriteString(fmt.Sprintf("        throw new IllegalArgumentException(\"Unsupported %s:\" + %s);\n", f.MatchKey, strcase.ToLowerCamel(f.MatchKey)))
	b.WriteString("    }\n")
	b.WriteString("    return supplier.get();\n")
	b.WriteString("}")
	return b.String()
}

// GetFieldName get CamelName
func (g JavaGenerator) GetFieldName(f *model.Field) string {
	return strcase.ToCamel(f.Name)
}

// GetFieldNameLower get lower CamelName
func (g JavaGenerator) GetFieldNameLower(f *model.Field) string {
	return strcase.ToLowerCamel(g.GetFieldName(f))
}

// GetFieldType get java type
func (g JavaGenerator) GetFieldType(f *model.Field) string {
	switch f.GetType() {
	case "string":
		return "String"
	case "i8", "u8", "char", "i16", "u16", "i32", "u32", "i64", "u64", "f32", "f64":
		return javaBasicTypeMap[f.GetType()].BasicType
	default:
		_, ok := ParseCharArrayType(f.Type)
		if ok {
			return "String"
		}
		if f.InerObject != nil {
			return f.InerObject.Name
		}

		if f.Type == "match" {
			return "BinaryCodec"
		}
		return f.Type
	}
}

// GenerateGetterAndSetter gen getter and setter method
func (g JavaGenerator) GenerateGetterAndSetter(f *model.Field) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("public %s get%s() {\n", g.GetFieldType(f), g.GetFieldName(f)))
	b.WriteString(AddIndent4ln(fmt.Sprintf("return this.%s;", g.GetFieldNameLower(f))))
	b.WriteString("}\n")
	b.WriteString("\n")
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("public void set%s(%s %s) {\n", g.GetFieldName(f), g.GetFieldType(f), g.GetFieldNameLower(f)))
	b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = %s;", g.GetFieldNameLower(f), g.GetFieldNameLower(f))))
	b.WriteString("}\n")
	return b.String()
}

// GenerateToString gen toString method
func (g JavaGenerator) GenerateToString(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public String toString() {\n")
	if len(packet.Fields) > 0 {
		b.WriteString(AddIndent4(fmt.Sprintf("return \"%s [\"", packet.Name)))
		for idx, f := range packet.Fields {
			if idx == 0 {
				b.WriteString(fmt.Sprintf(" + \"%s=\" + this.%s", g.GetFieldNameLower(&f), g.GetFieldNameLower(&f)))
			} else {
				b.WriteString(fmt.Sprintf(" + \", %s=\" + this.%s", g.GetFieldNameLower(&f), g.GetFieldNameLower(&f)))
			}
		}
		b.WriteString(" + \"]\";\n")
	} else {
		b.WriteString(AddIndent4(fmt.Sprintf("return \"%s []\";", packet.Name)))
	}
	b.WriteString("}\n")
	return b.String()
}

// GenerateEquals gen equals method
func (g JavaGenerator) GenerateEquals(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public boolean equals(Object obj) {\n")
	b.WriteString(AddIndent4ln("if(this == obj) {"))
	b.WriteString(AddIndent4ln(AddIndent4("return true;")))
	b.WriteString(AddIndent4ln("}"))
	b.WriteString(AddIndent4ln("if(null == obj || getClass() != obj.getClass()) {"))
	b.WriteString(AddIndent4ln(AddIndent4("return false;")))
	b.WriteString(AddIndent4ln("}"))
	if len(packet.Fields) > 0 {
		b.WriteString(AddIndent4ln(fmt.Sprintf("%s orther_ = (%s) obj;", packet.Name, packet.Name)))
		equals := make([]string, 0, len(packet.Fields))
		for _, f := range packet.Fields {
			equal := fmt.Sprintf("Objects.equals(%s, orther_.%s)",
				strcase.ToLowerCamel(f.Name), strcase.ToLowerCamel(f.Name))
			equals = append(equals, equal)
		}
		b.WriteString(AddIndent4ln(fmt.Sprintf("return %s;", strings.Join(equals, " && "))))
	} else {
		b.WriteString(AddIndent4ln("return true;"))
	}
	b.WriteString("}\n")
	return b.String()
}

// GenerateHashCode gen hasCode method
func (g JavaGenerator) GenerateHashCode(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public int hashCode() {\n")
	if len(packet.Fields) > 0 {
		names := make([]string, 0, len(packet.Fields))
		for _, f := range packet.Fields {
			names = append(names, strcase.ToLowerCamel(f.Name))
		}
		b.WriteString(AddIndent4ln(fmt.Sprintf("return Objects.hash(%s);", strings.Join(names, ", "))))
	} else {
		b.WriteString(AddIndent4ln("return this.hashCode();"))
	}

	b.WriteString("}\n")
	return b.String()
}

// GenerateDecode gen decode method
func (g JavaGenerator) GenerateDecode(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public void decode(ByteBuf byteBuf) {\n")
	for _, f := range packet.Fields {
		b.WriteString(AddIndent4ln(g.GenerateDecodeField(&f)))
	}
	b.WriteString("}\n")
	return b.String()
}

// GenerateDecodeField den decode field
func (g JavaGenerator) GenerateDecodeField(f *model.Field) string {
	switch f.Type {
	case "string", "char[]":
		var b strings.Builder
		fieldLen := g.GetFieldNameLower(f) + "Len"
		lenTyp := javaBasicTypeMap[g.config.StringLenPrefixLenType]
		if g.config.LittleEndian {
			b.WriteString(fmt.Sprintf("%s %s = byteBuf.read%s();\n", lenTyp.BasicType, fieldLen, lenTyp.Le))
		} else {
			b.WriteString(fmt.Sprintf("%s %s = byteBuf.read%s();\n", lenTyp.BasicType, fieldLen, strcase.ToCamel(lenTyp.BasicType)))
		}
		b.WriteString(fmt.Sprintf("if (%s > 0) {\n", fieldLen))
		b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = byteBuf.readCharSequence(%s, StandardCharsets.UTF_8).toString();", g.GetFieldNameLower(f), fieldLen)))
		b.WriteString("}")
		return b.String()
	default:
		if typ, ok := javaBasicTypeMap[f.GetType()]; ok {
			if g.config.LittleEndian {
				return fmt.Sprintf("this.%s = byteBuf.read%s();", g.GetFieldNameLower(f), typ.Le)
			}
			return fmt.Sprintf("this.%s = byteBuf.read%s();", g.GetFieldNameLower(f), strcase.ToCamel(typ.BasicType))
		}
		len, ok := ParseCharArrayType(f.Type)
		if ok {
			return fmt.Sprintf("this.%s = readFixedString(byteBuf, %s);", g.GetFieldNameLower(f), len)
		}
		if f.InerObject != nil {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("if (null == this.%s) {\n", strcase.ToLowerCamel(f.InerObject.Name)))
			b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = new %s();", strcase.ToLowerCamel(f.InerObject.Name), f.InerObject.Name)))
			b.WriteString("}\n")
			b.WriteString(fmt.Sprintf("this.%s.decode(byteBuf);", strcase.ToLowerCamel(f.InerObject.Name)))
			return b.String()
		}

		if f.Type == "match" {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("this.%s = create%s(this.%s);\n", g.GetFieldNameLower(f), g.GetFieldName(f), strcase.ToLowerCamel(f.MatchKey)))
			b.WriteString(fmt.Sprintf("this.%s.decode(byteBuf);", g.GetFieldNameLower(f)))
			return b.String()
		}
		if _, ok := g.binModel.PacketsMap[f.Name]; ok {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("if (null == this.%s) {\n", strcase.ToLowerCamel(f.Name)))
			b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = new %s();", strcase.ToLowerCamel(f.Name), f.Name)))
			b.WriteString("}\n")
			b.WriteString(fmt.Sprintf("this.%s.decode(byteBuf);", strcase.ToLowerCamel(f.Name)))
			return b.String()
		}
		return "--" + f.Type
	}
}

// GenerateEncode gen encode  method
func (g JavaGenerator) GenerateEncode(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public void encode(ByteBuf byteBuf) {\n")
	for _, f := range packet.Fields {
		b.WriteString(AddIndent4ln(g.GenerateEncodeField(packet, &f)))
	}
	b.WriteString("}\n")
	return b.String()
}

// GenerateEncodeField gen encode field
func (g JavaGenerator) GenerateEncodeField(p *model.Packet, f *model.Field) string {
	if f.LengthOfField != "" {
		// auto calculate length field
		var b strings.Builder
		b.WriteString(fmt.Sprintf("ByteBuf %sBuf = null;\n", strcase.ToLowerCamel(f.LengthOfField)))
		b.WriteString(fmt.Sprintf("if (this.%s != null) {\n", strcase.ToLowerCamel(f.LengthOfField)))
		b.WriteString(fmt.Sprintf("%sBuf = Unpooled.buffer();\n", strcase.ToLowerCamel(f.LengthOfField)))
		b.WriteString(fmt.Sprintf("this.%s.encode(%sBuf);\n", strcase.ToLowerCamel(f.LengthOfField), strcase.ToLowerCamel(f.LengthOfField)))
		lenTyp := javaBasicTypeMap[f.GetType()]
		b.WriteString(fmt.Sprintf("this.%s = (%s) %sBuf.readableBytes();\n", strcase.ToLowerCamel(f.Name), lenTyp.BasicType, strcase.ToLowerCamel(f.LengthOfField)))
		b.WriteString("} else {\n")
		b.WriteString(fmt.Sprintf("this.%s = 0;\n", strcase.ToLowerCamel(f.Name)))
		b.WriteString("}\n")

		if g.config.LittleEndian {
			b.WriteString(fmt.Sprintf("byteBuf.write%s(this.%s);\n", lenTyp.Le, strcase.ToLowerCamel(f.Name)))
		} else {
			b.WriteString(fmt.Sprintf("byteBuf.write%s(this.%s);\n", strcase.ToCamel(lenTyp.BasicType), strcase.ToLowerCamel(f.Name)))
		}

		return b.String()
	}
	switch f.GetType() {
	case "string", "char[]":
		var b strings.Builder
		b.WriteString(fmt.Sprintf("if (StringUtil.isNullOrEmpty(this.%s)) {\n", g.GetFieldNameLower(f)))
		b.WriteString(AddIndent4ln("byteBuf.writeShort(0);"))
		b.WriteString("} else {\n")
		b.WriteString(AddIndent4ln(fmt.Sprintf("byte[] bytes = this.%s.getBytes(StandardCharsets.UTF_8);", g.GetFieldNameLower(f))))
		lenTyp := javaBasicTypeMap[g.config.StringLenPrefixLenType]
		if g.config.LittleEndian {
			b.WriteString(AddIndent4ln(fmt.Sprintf("byteBuf.write%s(bytes.length);", lenTyp.Le)))
		} else {
			b.WriteString(AddIndent4ln(fmt.Sprintf("byteBuf.write%s(bytes.length);", strcase.ToCamel(lenTyp.BasicType))))
		}
		b.WriteString(AddIndent4ln("byteBuf.writeBytes(bytes);"))
		b.WriteString("}\n")
		return b.String()
	default:
		if typ, ok := javaBasicTypeMap[f.GetType()]; ok {
			if g.config.LittleEndian {
				return fmt.Sprintf("byteBuf.write%s(this.%s);", typ.Le, g.GetFieldNameLower(f))
			}
			return fmt.Sprintf("byteBuf.write%s(this.%s);", strcase.ToCamel(typ.BasicType), g.GetFieldNameLower(f))
		}
		len, ok := ParseCharArrayType(f.Type)
		if ok {
			return fmt.Sprintf("writeFixedString(byteBuf, this.%s, %s);", g.GetFieldNameLower(f), len)
		}

		if f.Name == p.LengthOfField {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("if (%sBuf != null) {\n", strcase.ToLowerCamel(f.Name)))
			b.WriteString(fmt.Sprintf("byteBuf.writeBytes(%sBuf);\n", strcase.ToLowerCamel(f.Name)))
			b.WriteString(fmt.Sprintf("%sBuf.release();\n", strcase.ToLowerCamel(f.Name)))
			b.WriteString("}\n")
			return b.String()
		}

		if f.InerObject != nil {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("if (null == this.%s) {\n", strcase.ToLowerCamel(f.InerObject.Name)))
			b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = new %s();", strcase.ToLowerCamel(f.InerObject.Name), f.InerObject.Name)))
			b.WriteString("}\n")
			b.WriteString(fmt.Sprintf("this.%s.encode(byteBuf);", strcase.ToLowerCamel(f.InerObject.Name)))
			return b.String()
		}

		if f.Type == "match" {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("if (null != this.%s) {\n", g.GetFieldNameLower(f)))
			b.WriteString(fmt.Sprintf("this.%s.encode(byteBuf);", g.GetFieldNameLower(f)))
			b.WriteString("}")
			return b.String()
		}
		if _, ok := g.binModel.PacketsMap[f.Name]; ok {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("if (null == this.%s) {\n", strcase.ToLowerCamel(f.Name)))
			b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = new %s();", strcase.ToLowerCamel(f.Name), f.Name)))
			b.WriteString("}\n")
			b.WriteString(fmt.Sprintf("this.%s.encode(byteBuf);", strcase.ToLowerCamel(f.Name)))
			return b.String()
		}
		return "//" + f.Type
	}
}

// GenerateJavaTestClassFileForPacket gen test class
func (g JavaGenerator) GenerateJavaTestClassFileForPacket(packet *model.Packet) string {
	var b strings.Builder
	//package
	b.WriteString(fmt.Sprintf("package %s;\n", g.config.JavaPackage))
	//import
	b.WriteString(`import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import org.junit.Test;

import static org.junit.Assert.*;
`)
	b.WriteString("\n")
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("public class %sTest {\n", packet.Name))

	b.WriteString(AddIndent4ln(g.GenerateTestMethod(packet)))
	b.WriteString("}\n")

	return b.String()
}

// GenerateTestMethod gen test method
func (g JavaGenerator) GenerateTestMethod(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Test\n")
	b.WriteString("public void testEncodeDecode() {\n")
	b.WriteString(AddIndent4ln(g.GenerateNewInstance("original", "", packet)))
	b.WriteString(AddIndent4ln("ByteBuf buffer = Unpooled.buffer();"))
	b.WriteString(AddIndent4ln("original.encode(buffer);"))
	b.WriteString(AddIndent4ln(fmt.Sprintf("%s decoded = new %s();", packet.Name, packet.Name)))
	b.WriteString(AddIndent4ln("decoded.decode(buffer);"))
	b.WriteString(AddIndent4ln("assertEquals(original, decoded);"))
	b.WriteString("}")
	return b.String()
}

// GenerateNewInstance new instance for unitest
func (g JavaGenerator) GenerateNewInstance(instanceName string, parent string, packet *model.Packet) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s%s %s = new %s%s();\n", parent, packet.Name, instanceName, parent, packet.Name))
	for _, f := range packet.Fields {
		if f.LengthOfField != "" {
			continue
		} else if _, ok := packet.MatchFields[f.Name]; ok {
			continue
		} else if typ, ok := javaBasicTypeMap[f.GetType()]; ok {
			b.WriteString(fmt.Sprintf("%s.set%s(%s);\n", instanceName, strcase.ToCamel(f.Name), typ.TestValue))
		} else if len, ok := ParseCharArrayType(f.GetType()); ok {
			l, _ := strconv.Atoi(len)
			b.WriteString(fmt.Sprintf("%s.set%s(\"%s\");\n", instanceName, strcase.ToCamel(f.Name), strings.Repeat("1", l)))
		} else if f.GetType() == "string" {
			b.WriteString(fmt.Sprintf("%s.set%s(\"example\");\n", instanceName, strcase.ToCamel(f.Name)))
		} else if f.InerObject != nil {
			inerObjName := strcase.ToLowerCamel(f.InerObject.Name)
			b.WriteString(g.GenerateNewInstance(inerObjName, packet.Name+".", f.InerObject))
			b.WriteString(fmt.Sprintf("%s.set%s(%s);\n", instanceName, strcase.ToCamel(f.Name), inerObjName))
		} else if f.GetType() == "match" {
			//
			key := f.MatchPairs[0].Key
			value := f.MatchPairs[0].Value
			if typ, ok := javaBasicTypeMap[packet.FieldMap[f.MatchKey].GetType()]; ok {
				key = fmt.Sprintf("(%s)%s", typ.BasicType, key)
			}
			b.WriteString(fmt.Sprintf("%s.set%s(%s);\n", instanceName, strcase.ToCamel(f.MatchKey), key))
			if p, ok := g.binModel.PacketsMap[value]; ok {
				inerObjName := strcase.ToLowerCamel(p.Name)
				b.WriteString(g.GenerateNewInstance(inerObjName, "", &p))
				b.WriteString(fmt.Sprintf("%s.set%s(%s);\n", instanceName, strcase.ToCamel(f.Name), inerObjName))
			}
		} else {
			if p, ok := g.binModel.PacketsMap[f.Name]; ok {
				inerObjName := strcase.ToLowerCamel(p.Name)
				b.WriteString(g.GenerateNewInstance(inerObjName, "", &p))
				b.WriteString(fmt.Sprintf("%s.set%s(%s);\n", instanceName, strcase.ToCamel(f.Name), inerObjName))
			} else {
				b.WriteString("--" + f.GetType())
			}
		}

	}
	return b.String()
}
