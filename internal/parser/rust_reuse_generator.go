package parser

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

func (g RustGenerator) generateDynamicStringIntoHelperCode() string {
	return `#[inline]
fn get_string_field_into<L: LengthPrefix>(
    buf: &mut Bytes,
    value: &mut String,
    little_endian: bool,
) -> Option<()> {
    let len = L::read(buf, little_endian)?;
    if buf.remaining() < len {
        return None;
    }
    let bytes = buf.copy_to_bytes(len);
    let text = std::str::from_utf8(&bytes).ok()?;
    value.clear();
    value.push_str(text);
    Some(())
}
`
}

func (g RustGenerator) hasReusableDynamicString(pkt *model.Packet) bool {
	if g.canGenerateBorrowed(pkt) {
		for _, f := range pkt.Fields {
			if _, ok := f.Attr.(*model.DynamicStringFieldAttribute); ok {
				return true
			}
		}
	}
	for _, f := range pkt.Fields {
		if object, ok := f.Attr.(*model.ObjectFieldAttribute); ok && object.IsIner {
			if g.hasReusableDynamicString(object.RefPacket) {
				return true
			}
		}
	}
	return false
}

func (g RustGenerator) hasReusableFixedString(pkt *model.Packet) bool {
	if g.canGenerateBorrowed(pkt) {
		for _, f := range pkt.Fields {
			if _, ok := f.Attr.(*model.FixedStringFieldAttribute); ok {
				return true
			}
		}
	}
	for _, f := range pkt.Fields {
		if object, ok := f.Attr.(*model.ObjectFieldAttribute); ok && object.IsIner {
			if g.hasReusableFixedString(object.RefPacket) {
				return true
			}
		}
	}
	return false
}

func (g RustGenerator) generateDecodeIntoCode(pkt *model.Packet) string {
	if !g.canGenerateBorrowed(pkt) {
		return ""
	}
	structName := strcase.ToCamel(pkt.Name)
	var b strings.Builder
	b.WriteString(fmt.Sprintf("impl %s {\n", structName))
	b.WriteString(AddIndent4ln("/// Decodes into an existing value, reusing owned string capacity."))
	b.WriteString(AddIndent4ln("///"))
	b.WriteString(AddIndent4ln("/// If decoding fails, `out` may have been partially updated. Only use it"))
	b.WriteString(AddIndent4ln("/// after this method returns `Some(())`."))
	b.WriteString(AddIndent4ln("#[inline]"))
	if len(pkt.Fields) == 0 {
		b.WriteString(AddIndent4ln("pub fn decode_into(_buf: &mut Bytes, _out: &mut Self) -> Option<()> {"))
	} else {
		b.WriteString(AddIndent4ln("pub fn decode_into(buf: &mut Bytes, out: &mut Self) -> Option<()> {"))
	}
	for _, f := range pkt.Fields {
		b.WriteString(AddIndent4ln(AddIndent4(g.decodeIntoField(structName, f))))
	}
	b.WriteString(AddIndent4ln(AddIndent4("Some(())")))
	b.WriteString(AddIndent4ln("}"))
	b.WriteString("}\n")
	return b.String()
}

func (g RustGenerator) decodeIntoField(parentName string, f *model.Field) string {
	name := g.GetFieldName(f)
	padding := g.GetPadding(f)
	switch attr := f.Attr.(type) {
	case *model.FixedStringFieldAttribute:
		return fmt.Sprintf("get_fixed_string_field_into(buf, &mut out.%s, %d, b%s, %t)?;", name, attr.Length, padding.PadChar, padding.PadLeft)
	case *model.DynamicStringFieldAttribute:
		return fmt.Sprintf("get_string_field_into::<%s>(buf, &mut out.%s, %t)?;", g.GetConfig().StringLenPrefixLenType, name, g.GetConfig().LittleEndian)
	case *model.MatchFieldAttribute:
		return g.decodeIntoMatchField(parentName, f, attr)
	case *model.ObjectFieldAttribute:
		return fmt.Sprintf("%s::decode_into(buf, &mut out.%s)?;", strcase.ToCamel(attr.RefPacket.Name), name)
	}

	switch f.GetType() {
	case "char":
		return fmt.Sprintf("out.%s = get_char(buf)?;", name)
	case "u8", "i8":
		return fmt.Sprintf("out.%s = buf.get_%s();", name, f.GetType())
	case "u16", "u32", "u64", "i16", "i32", "i64", "f32", "f64":
		if g.GetConfig().LittleEndian {
			return fmt.Sprintf("out.%s = buf.get_%s_le();", name, f.GetType())
		}
		return fmt.Sprintf("out.%s = buf.get_%s();", name, f.GetType())
	default:
		return fmt.Sprintf("// unknown type for decode_into: %s", f.GetType())
	}
}

func (g RustGenerator) decodeIntoMatchField(parentName string, f *model.Field, match *model.MatchFieldAttribute) string {
	matchKey := "out." + g.GetFieldName(match.MatchKeyField)
	if len(match.MatchPairs) > 0 && g.HasQuotes(match.MatchPairs[0].Key) {
		matchKey += ".as_str()"
	}
	fieldName := g.GetFieldName(f)
	enumName := parentName + f.Name + "Enum"
	var b strings.Builder
	b.WriteString(fmt.Sprintf("match %s {\n", matchKey))
	seen := make(map[string]struct{})
	for _, pair := range match.MatchPairs {
		if _, exists := seen[pair.Key]; exists {
			continue
		}
		seen[pair.Key] = struct{}{}
		childName := strcase.ToCamel(pair.Value)
		b.WriteString(AddIndent4ln(fmt.Sprintf("%s => match &mut out.%s {", pair.Key, fieldName)))
		b.WriteString(AddIndent4ln(AddIndent4(fmt.Sprintf("%s::%s(value) => %s::decode_into(buf, value)?,", enumName, pair.Value, childName))))
		b.WriteString(AddIndent4ln(AddIndent4(fmt.Sprintf("slot => *slot = %s::%s(%s::decode(buf)?),", enumName, pair.Value, childName))))
		b.WriteString(AddIndent4ln("},"))
	}
	b.WriteString(AddIndent4ln("_ => return None,"))
	b.WriteString("}")
	return b.String()
}
