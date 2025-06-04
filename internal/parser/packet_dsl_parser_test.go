package parser

import (
	"os"
	"reflect"
	"testing"

	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// writeTempFile helper writes content to a temporary file and returns its path
func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "test-*.dsl")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	return f.Name()
}

func TestParseSimplePacket(t *testing.T) {
	// Example DSL: one root packet with two simple fields
	dsl := `
root packet Logon {
	string username
	u16 password
}`
	path := writeTempFile(t, dsl)
	defer os.Remove(path)

	data, err := ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile returned error: %v", err)
	}
	// Expect []model.Packet with one entry
	packets := data.(*model.BinaryModel).Packets
	if len(packets) != 1 {
		t.Fatalf("expected 1 packet, got %d", len(packets))
	}
	pkt := packets["Logon"]
	if pkt.Name != "Logon" {
		t.Errorf("expected packet name 'Logon', got '%s'", pkt.Name)
	}
	if !pkt.IsRoot {
		t.Errorf("expected IsRoot=true")
	}
	// Check fields
	expectedFields := []model.Field{
		{Name: "username", Type: "string", IsRepeat: false},
		{Name: "password", Type: "u16", IsRepeat: false},
	}
	if !reflect.DeepEqual(pkt.Fields, expectedFields) {
		t.Errorf("fields mismatch. expected %+v, got %+v", expectedFields, pkt.Fields)
	}
}

func TestParseNestedAndMatchField(t *testing.T) {
	// DSL contains nested object and match field
	dsl := `
packet Parent {
	u16 MsgType,
	nestedChild {
		childField
	}
	match MsgType {
		0:OptionA,
		1:OptionB
	}
}`
	path := writeTempFile(t, dsl)
	defer os.Remove(path)

	data, err := ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile returned error: %v", err)
	}
	packets := data.(*model.BinaryModel).Packets
	if len(packets) != 1 {
		t.Fatalf("expected 1 packet, got %d", len(packets))
	}
	pkt := packets["Parent"]
	if pkt.Name != "Parent" {
		t.Errorf("expected packet name 'Parent', got '%s'", pkt.Name)
	}
	// Should have two fields: nestedChild and matchType
	if len(pkt.Fields) != 3 {
		t.Fatalf("expected 2 fields, got %d", len(pkt.Fields))
	}
	// Verify nestedChild
	nested := pkt.Fields[1]
	if nested.Name != "nestedChild" {
		t.Errorf("expected first field 'nestedChild', got '%s'", nested.Name)
	}
	if nested.Type != "" {
		t.Errorf("expected nestedChild Type=empty, got '%s'", nested.Type)
	}
	if nested.InerObject.Name != "nestedChild" {
		t.Errorf("expected nested InerObject.Name 'nestedChild', got '%s'", nested.InerObject.Name)
	}
	if len(nested.InerObject.Fields) != 1 || nested.InerObject.Fields[0].Name != "childField" {
		t.Errorf("nested child fields incorrect: %+v", nested.InerObject.Fields)
	}
	// Verify matchType
	matchField := pkt.Fields[2]
	if matchField.Name != "MsgType" {
		t.Errorf("expected second field 'matchType', got '%s'", matchField.Name)
	}
	if len(matchField.MatchPairs) != 2 {
		t.Errorf("expected 2 match pairs, got %d", len(matchField.MatchPairs))
	}
	if matchField.MatchPairs[0].Key != "0" || matchField.MatchPairs[0].Value != "OptionA" {
		t.Errorf("first match pair incorrect: %+v", matchField.MatchPairs[0])
	}
	if matchField.MatchPairs[1].Key != "1" || matchField.MatchPairs[1].Value != "OptionB" {
		t.Errorf("second match pair incorrect: %+v", matchField.MatchPairs[1])
	}
}
