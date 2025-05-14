package parser_test

import (
	"fmt"
	"testing"

	// Update this to the correct import path
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

// TestVisitPacketDefinition tests the VisitPacketDefinition method of the visitor.
func TestVisitPacketDefinition(t *testing.T) {
	fileName := "testdata/test_packet.dsl"
	astData, err := parser.ParseFile(fileName)
	if err != nil {
		t.Fatalf("Failed to parse file: %v", err)
	}

	fmt.Println("AST Data:", astData)

}
