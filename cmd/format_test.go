package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

// runFormat executes the format command with the given flags and returns the
// command error (nil means exit code 0 for the CLI).
func runFormat(t *testing.T, args ...string) error {
	t.Helper()
	// reset the shared flag bindings so values cannot leak between tests
	file, dsl = "", ""
	rootCmd.SetArgs(append([]string{"format"}, args...))
	return rootCmd.Execute()
}

func TestFormatFromDslString(t *testing.T) {
	assert.NoError(t, runFormat(t, "-d", "packet A { u16 X, }"))
}

func TestFormatRequiresInput(t *testing.T) {
	err := runFormat(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "please provide a DSL string or a file path")
}

func TestFormatFileNotFound(t *testing.T) {
	err := runFormat(t, "-f", filepath.Join(t.TempDir(), "missing.dsl"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read file")
}

func TestFormatWritesBackToFile(t *testing.T) {
	input := "packet A{u16 X,}"
	path := filepath.Join(t.TempDir(), "in.dsl")
	require.NoError(t, os.WriteFile(path, []byte(input), 0644))

	assert.NoError(t, runFormat(t, "-f", path))

	expected, err := parser.FormatPacketDsl(input)
	require.NoError(t, err)
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, expected, string(data))
	assert.NotEqual(t, input, string(data), "file should have been reformatted")
}

// TestFormatWriteFailureReturnsError guards the write-back error path: a
// failed write must surface as a command error (non-zero exit), not success.
func TestFormatWriteFailureReturnsError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "readonly.dsl")
	require.NoError(t, os.WriteFile(path, []byte("packet A { u16 X, }"), 0444))
	t.Cleanup(func() { _ = os.Chmod(path, 0644) })

	err := runFormat(t, "-f", path)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "write formatted DSL")
}
