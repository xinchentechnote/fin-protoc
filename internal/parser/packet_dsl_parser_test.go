package parser_test

import (
	"fmt"
	"os"
	"testing"

	// Update this to the correct import path
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

// TestVisitPacketDefinition tests the VisitPacketDefinition method of the visitor.
func TestVisitPacketDefinition(t *testing.T) {
	// Mock input data (simulating the contents of a protocol file)
	input := `packet SzseBinary {field1}`

	// Create a temporary file for testing
	fileName := "test_protocol.txt"

	// Write the input string to the test file
	err := os.WriteFile(fileName, []byte(input), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Call ParseFile to parse the file content and get the AST
	astData, err := parser.ParseFile(fileName)
	if err != nil {
		t.Fatalf("Failed to parse file: %v", err)
	}

	fmt.Println("AST Data:", astData)

	// Clean up the test file
	err = os.Remove(fileName)
	if err != nil {
		t.Fatalf("Failed to remove test file: %v", err)
	}
}
