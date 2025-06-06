package parser_test

import (
	"strings"
	"testing"

	"github.com/xinchentechnote/fin-protoc/internal/model"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

func TestGenerateCode(t *testing.T) {
	packet := model.Packet{
		Name: "Logon",
		Fields: []model.Field{
			{Name: "Username", Type: "String"},
			{Name: "Password", Type: "String"},
			{Name: "SeqNum", Type: "u32"},
		},
	}

	gen := parser.RustGenerator{}
	code := gen.GenerateCode(packet)

	if !strings.Contains(code, "pub struct Logon") {
		t.Errorf("expected struct definition, got: %s", code)
	}
	if !strings.Contains(code, "pub username: String") {
		t.Errorf("missing username field")
	}
	if !strings.Contains(code, "buf.put_u32(self.seq_num);") {
		t.Errorf("missing encode for seq_num")
	}
	if !strings.Contains(code, "let seq_num = buf.get_u32();") {
		t.Errorf("missing decode for seq_num")
	}
}

func TestGenerateTestCode(t *testing.T) {
	packet := model.Packet{
		Name: "Heartbeat",
		Fields: []model.Field{
			{Name: "TestReqId", Type: "String"},
			{Name: "SeqNum", Type: "u64"},
		},
	}

	gen := parser.RustGenerator{}
	testCode := gen.GenerateTestCode(packet)

	if !strings.Contains(testCode, "#[test]") {
		t.Error("expected #[test] attribute in generated test")
	}
	if !strings.Contains(testCode, "let original = Heartbeat") {
		t.Error("expected original instance construction")
	}
	if !strings.Contains(testCode, "assert_eq!(original, decoded);") {
		t.Error("expected assert_eq in test")
	}
}
