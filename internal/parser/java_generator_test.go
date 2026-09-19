package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generatorSmokeDSL exercises the field kinds every generator must handle:
// basic types, fixed/dynamic strings, repeat, match and object references.
const generatorSmokeDSL = `options {
	JavaPackage = "com.example.messages";
	GoPackage = "messages";
	GoModule = "github.com/example/messages";
}
packet Logon {
	char[10] UserName ` + "`user name`" + `,
	string Password ` + "`password`" + `,
	u64 ClientId,
}
root packet SampleBinary {
	u16 MsgType,
	match MsgType as Body {
		1: Logon,
	},
	u32 Count,
	repeat string Extra,
}`

func TestJavaGeneratorGenerate(t *testing.T) {
	binModel := parseDSL(t, generatorSmokeDSL)

	output, err := NewJavaGenerator(binModel).Generate(binModel)
	require.NoError(t, err)

	assert.Contains(t, output, "main/java/com/example/messages/SampleBinary.java")
	assert.Contains(t, output, "test/java/com/example/messages/SampleBinaryTest.java")
	assert.Contains(t, output, "main/java/com/example/messages/Logon.java")

	rootCode := string(output["main/java/com/example/messages/SampleBinary.java"])
	assert.Contains(t, rootCode, "class SampleBinary")
	assert.Contains(t, rootCode, "public void encode")
	assert.Contains(t, rootCode, "public void decode")
}
