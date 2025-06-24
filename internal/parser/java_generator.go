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
		code := g.GenerateJavaClassFileForPacket(&pkg)
		filePath := "main/java/" + packagePath + "/" + pkg.Name + ".java"
		output[filePath] = []byte(code)
	}
	return output, nil
}

// GenerateJavaClassFileForPacket gen java class file
func (g JavaGenerator) GenerateJavaClassFileForPacket(packet *model.Packet) string {
	var b strings.Builder
	//package
	b.WriteString(fmt.Sprintf("package %s;\n", g.config.JavaPackage))
	//import
	b.WriteString(`import java.nio.charset.StandardCharsets;

import com.finproto.codec.BinaryCodec;

import io.netty.buffer.ByteBuf;
import io.netty.util.internal.StringUtil;`)
	b.WriteString("\n")
	b.WriteString("\n")
	//class
	b.WriteString(fmt.Sprintf("public class %s implements BinaryCodec {\n", packet.Name))
	//field
	for _, f := range packet.Fields {
		b.WriteString(AddIndent4ln(fmt.Sprintf("private %s %s;", g.GetType(f.Type), strcase.ToLowerCamel(f.Name))))
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

	b.WriteString("\n")
	//hashcode and equals
	b.WriteString(AddIndent4ln(g.GenerateHashCode(packet)))
	b.WriteString(AddIndent4ln(g.GenerateEquals(packet)))

	b.WriteString("\n")
	//tostring
	b.WriteString(AddIndent4ln(g.GenerateToString(packet)))

	b.WriteString("\n")
	b.WriteString("}")
	return b.String()
}

// GetType get java type
func (g JavaGenerator) GetType(typ string) string {
	switch typ {
	case "string":
		return "String"
	case "i8", "u8":
		return "byte"
	case "i16", "u16":
		return "short"
	case "i32", "u32":
		return "int"
	case "i64", "u64":
		return "long"
	case "f32":
		return "float"
	case "f64":
		return "double"
	default:
		return typ
	}
}

// GenerateGetterAndSetter gen getter and setter method
func (g JavaGenerator) GenerateGetterAndSetter(f *model.Field) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("public %s get%s() {\n", g.GetType(f.Type), f.Name))
	b.WriteString(AddIndent4ln(fmt.Sprintf("return this.%s;", strcase.ToLowerCamel(f.Name))))
	b.WriteString("}\n")
	b.WriteString("\n")
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("public void set%s(%s %s) {\n", f.Name, g.GetType(f.Type), strcase.ToLowerCamel(f.Name)))
	b.WriteString(AddIndent4ln(fmt.Sprintf("this.%s = %s;", strcase.ToLowerCamel(f.Name), strcase.ToLowerCamel(f.Name))))
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
	b.WriteString("}\n")
	return b.String()
}

// GenerateEncode gen encode  method
func (g JavaGenerator) GenerateEncode(packet *model.Packet) string {
	var b strings.Builder
	b.WriteString("@Override\n")
	b.WriteString("public void encode(ByteBuf byteBuf) {\n")
	b.WriteString("}\n")
	return b.String()
}
