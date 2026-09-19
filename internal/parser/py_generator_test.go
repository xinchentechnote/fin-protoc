package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPythonGeneratorGenerate(t *testing.T) {
	binModel := parseDSL(t, generatorSmokeDSL)

	output, err := NewPythonGenerator(binModel).Generate(binModel)
	require.NoError(t, err)

	assert.Contains(t, output, "sample_binary.py")
	assert.Contains(t, output, "sample_binary_test.py")

	rootCode := string(output["sample_binary.py"])
	assert.Contains(t, rootCode, "class SampleBinary")
	assert.Contains(t, rootCode, "def encode")
	assert.Contains(t, rootCode, "def decode")
}
