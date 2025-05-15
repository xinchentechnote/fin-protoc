package parser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

func TestFormatPacketDsl(t *testing.T) {
	dsl, _ := parser.ReadFileToString("testdata/need_format.dsl")
	expectedDsl, _ := parser.ReadFileToString("testdata/formatted.dsl")
	formattedDsl, err := parser.FormatPacketDsl(dsl)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, strings.ReplaceAll(expectedDsl, "\r\n", "\n"), formattedDsl)
}

func TestFormatPacketDsl1(t *testing.T) {
	formattedDsl, err := parser.FormatPacketDsl("packet MyPacket {  Int64 name   `description` }")
	assert.NoError(t, err)
	fmt.Println(formattedDsl)
}
