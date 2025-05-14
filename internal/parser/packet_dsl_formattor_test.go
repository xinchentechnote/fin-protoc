package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

func TestFormatPacketDsl(t *testing.T) {
	formattedDsl := parser.FormatPacketDsl("packet SzseBinary {field1}")
	assert.Equal(t, "packet SzseBinary {field1}", formattedDsl)
}
