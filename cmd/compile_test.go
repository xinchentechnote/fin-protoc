package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const compileTestDSL = `options {
	GoPackage = "messages";
	GoModule = "github.com/example/messages";
}
packet Logon {
	u32 ClientId,
}
root packet RootPacket {
	u16 MsgType,
	match MsgType as Body {
		1: Logon,
	},
}`

func writeDslFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.dsl")
	require.NoError(t, os.WriteFile(path, []byte(content), 0644))
	return path
}

func TestCompileGeneratesGoCode(t *testing.T) {
	outDir := t.TempDir()
	err := Compile(writeDslFile(t, compileTestDSL), map[string]string{"go": outDir})
	require.NoError(t, err)

	data, err := os.ReadFile(filepath.Join(outDir, "root_packet.go"))
	require.NoError(t, err)
	assert.Contains(t, string(data), "type RootPacket struct")

	_, err = os.Stat(filepath.Join(outDir, "root_packet_test.go"))
	assert.NoError(t, err)
}

func TestCompileNoOutputsIsNoop(t *testing.T) {
	err := Compile(writeDslFile(t, compileTestDSL), map[string]string{})
	assert.NoError(t, err)
}

func TestCompileFileNotFound(t *testing.T) {
	err := Compile(filepath.Join(t.TempDir(), "missing.dsl"), map[string]string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse file")
}

func TestCompileSyntaxError(t *testing.T) {
	err := Compile(writeDslFile(t, "packet A { u16 "), map[string]string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse file")
}

func TestCompileSemanticError(t *testing.T) {
	dsl := `packet Dup {}
packet Dup {}`
	err := Compile(writeDslFile(t, dsl), map[string]string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "found 1 syntax errors")
}
