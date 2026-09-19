package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCppGeneratorGenerate(t *testing.T) {
	binModel := parseDSL(t, generatorSmokeDSL)

	output, err := NewCppGenerator(binModel).Generate(binModel)
	require.NoError(t, err)

	assert.Contains(t, output, "include/sample_binary.hpp")
	assert.Contains(t, output, "test/sample_binary_test.cpp")

	hppCode := string(output["include/sample_binary.hpp"])
	assert.Contains(t, hppCode, "struct SampleBinary : public codec::BinaryCodec")
	assert.Contains(t, hppCode, "void encode")
	assert.Contains(t, hppCode, "void decode")
}
