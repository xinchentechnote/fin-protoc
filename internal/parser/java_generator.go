package parser

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

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
		code := g.GenerateJavaClassFileForPacket(&pkg, false)
		filePath := "main/java/" + packagePath + "/" + pkg.Name + ".java"
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
import java.util.function.Supplier;
import com.finproto.codec.BinaryCodec;

import io.netty.buffer.ByteBuf;
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
			b.WriteString(AddIndent4ln(g.GenerateCreateMethod(&f)))
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
func (g JavaGenerator) GenerateCreateMethod(f *model.Field) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("private static final Map<Short, Supplier<BinaryCodec>> %sMap = new HashMap<>();", g.GetFieldNameLower(f)))
	b.WriteString("\n")
	b.WriteString("static {\n")
	for _, pair := range f.MatchPairs {
		b.WriteString(fmt.Sprintf("    %sMap.put((short) %s, %s::new);\n", g.GetFieldNameLower(f), pair.Key, pair.Value))
	}
	b.WriteString("}\n")
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("private BinaryCodec create%s(short %s) {\n", g.GetFieldName(f), strcase.ToLowerCamel(f.MatchKey)))
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
	switch f.Type {
	case "string":
		return "String"
	case "i8", "u8", "char":
		return "byte"
	case "i16", "u16", "int16", "uint16":
		return "short"
	case "i32", "u32", "int32", "uint32":
		return "int"
	case "i64", "u64", "int64", "uint64":
		return "long"
	case "f32":
		return "float"
	case "f64":
		return "double"
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
	b.WriteString(AddIndent4ln("return \"\";"))
	b.WriteString("}\n")
	return b.String()
}

// GenerateEquals gen equals method
func (g JavaGenerator) GenerateEquals(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public boolean equals(Object obj) {\n")
	b.WriteString(AddIndent4ln("return true;"))
	b.WriteString("}\n")
	return b.String()
}

// GenerateHashCode gen hasCode method
func (g JavaGenerator) GenerateHashCode(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public int hashCode() {\n")
	b.WriteString(AddIndent4ln("return 0;"))
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
	case "string":
		var b strings.Builder
		fieldLen := g.GetFieldNameLower(f) + "Len"
		b.WriteString(fmt.Sprintf("short %s = byteBuf.readShort();\n", fieldLen))
		b.WriteString(fmt.Sprintf("if (%s > 0) {\n", fieldLen))
		b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = byteBuf.readCharSequence(%s, StandardCharsets.UTF_8).toString();", g.GetFieldNameLower(f), fieldLen)))
		b.WriteString("}")
		return b.String()
	case "i8", "u8", "char":
		return fmt.Sprintf("this.%s = byteBuf.readByte();", g.GetFieldNameLower(f))
	case "i16", "u16", "int16", "uint16":
		return fmt.Sprintf("this.%s = byteBuf.readShort();", g.GetFieldNameLower(f))
	case "i32", "u32", "int32", "uint32":
		return fmt.Sprintf("this.%s = byteBuf.readInt();", g.GetFieldNameLower(f))
	case "i64", "u64", "int64", "uint64":
		return fmt.Sprintf("this.%s = byteBuf.readLong();", g.GetFieldNameLower(f))
	case "f32":
		return fmt.Sprintf("this.%s = byteBuf.readFloat();", g.GetFieldNameLower(f))
	case "f64":
		return fmt.Sprintf("this.%s = byteBuf.readDouble();", g.GetFieldNameLower(f))
	default:
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
		return "--" + f.Type
	}
}

// GenerateEncode gen encode  method
func (g JavaGenerator) GenerateEncode(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public void encode(ByteBuf byteBuf) {\n")
	for _, f := range packet.Fields {
		b.WriteString(AddIndent4ln(g.GenerateEncodeField(&f)))
	}
	b.WriteString("}\n")
	return b.String()
}

// GenerateEncodeField gen encode field
func (g JavaGenerator) GenerateEncodeField(f *model.Field) string {
	switch f.Type {
	case "string":
		var b strings.Builder
		b.WriteString(fmt.Sprintf("if (StringUtil.isNullOrEmpty(this.%s)) {\n", g.GetFieldNameLower(f)))
		b.WriteString(AddIndent4ln("byteBuf.writeShort(0);"))
		b.WriteString("} else {\n")
		b.WriteString(AddIndent4ln(fmt.Sprintf("byte[] bytes = this.%s.getBytes(StandardCharsets.UTF_8);", g.GetFieldNameLower(f))))
		b.WriteString(AddIndent4ln("byteBuf.writeShort(bytes.length);"))
		b.WriteString(AddIndent4ln("byteBuf.writeBytes(bytes);"))
		b.WriteString("}\n")
		return b.String()
	case "i8", "u8", "char":
		return fmt.Sprintf("byteBuf.writeByte(this.%s);", g.GetFieldNameLower(f))
	case "i16", "u16", "int16", "uint16":
		return fmt.Sprintf("byteBuf.writeShort(this.%s);", g.GetFieldNameLower(f))
	case "i32", "u32", "int32", "uint32":
		return fmt.Sprintf("byteBuf.writeInt(this.%s);", g.GetFieldNameLower(f))
	case "i64", "u64", "int64", "uint64":
		return fmt.Sprintf("byteBuf.writeLong(this.%s);", g.GetFieldNameLower(f))
	case "f32":
		return fmt.Sprintf("byteBuf.writeFloat(this.%s);", g.GetFieldNameLower(f))
	case "f64":
		return fmt.Sprintf("byteBuf.writeDouble(this.%s);", g.GetFieldNameLower(f))
	default:
		len, ok := ParseCharArrayType(f.Type)
		if ok {
			return fmt.Sprintf("writeFixedString(byteBuf, this.%s, %s);", g.GetFieldNameLower(f), len)
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
		return "//" + f.Type
	}
}
