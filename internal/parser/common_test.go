package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAndIndent4ln(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single line",
			input:    "test",
			expected: "    test\n",
		},
		{
			name:     "multi line",
			input:    "line1\nline2",
			expected: "    line1\n    line2\n",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "    \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddIndent4ln(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAddIndent4(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single line",
			input:    "test",
			expected: "    test",
		},
		{
			name:     "multi line",
			input:    "line1\nline2",
			expected: "    line1\n    line2",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "    ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddIndent4(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAddIndent(t *testing.T) {
	tests := []struct {
		name     string
		spaces   int
		input    string
		expected string
	}{
		{
			name:     "2 spaces single line",
			spaces:   2,
			input:    "test",
			expected: "  test",
		},
		{
			name:     "3 spaces multi line",
			spaces:   3,
			input:    "line1\nline2",
			expected: "   line1\n   line2",
		},
		{
			name:     "0 spaces",
			spaces:   0,
			input:    "test",
			expected: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddIndent(tt.input, tt.spaces)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReadFileToString(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	content := "test content"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	tempFile.Close()

	tests := []struct {
		name        string
		filename    string
		expected    string
		expectError bool
	}{
		{
			name:        "existing file",
			filename:    tempFile.Name(),
			expected:    content,
			expectError: false,
		},
		{
			name:        "non-existent file",
			filename:    "nonexistent.txt",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ReadFileToString(tt.filename)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestNewPacketDslParserByFile(t *testing.T) {
	// Create a temporary file with valid DSL content
	tempFile, err := os.CreateTemp("", "testdsl")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	content := "packet TestPacket { u32 field1; }"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	tempFile.Close()

	tests := []struct {
		name        string
		filename    string
		expectError bool
	}{
		{
			name:        "valid DSL file",
			filename:    tempFile.Name(),
			expectError: false,
		},
		{
			name:        "non-existent file",
			filename:    "nonexistent.dsl",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, _, err := NewPacketDslParserByFile(tt.filename)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, parser)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, parser)
			}
		})
	}
}

func TestNewPacketDslParserByContent(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expectError bool
	}{
		{
			name:        "valid DSL content",
			content:     "packet TestPacket { u32 field1; }",
			expectError: false,
		},
		{
			name:        "empty content",
			content:     "",
			expectError: false, // ANTLR can parse empty content
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, _, err := NewPacketDslParserByContent(tt.content)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, parser)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, parser)
			}
		})
	}
}

func TestParseCharArrayType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		ok       bool
	}{
		{
			name:     "valid char array",
			input:    "char[10]",
			expected: "10",
			ok:       true,
		},
		{
			name:     "invalid format",
			input:    "char10]",
			expected: "",
			ok:       false,
		},
		{
			name:     "not a char array",
			input:    "string",
			expected: "",
			ok:       false,
		},
		{
			name:     "empty brackets",
			input:    "char[]",
			expected: "",
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, ok := ParseCharArrayType(tt.input)
			assert.Equal(t, tt.ok, ok)
			if ok {
				assert.Equal(t, tt.expected, size)
			}
		})
	}
}

func TestRenderToString(t *testing.T) {
	tests := []struct {
		name        string
		tmpl        string
		lang        string
		data        interface{}
		expected    string
		expectError bool
	}{
		{
			name:        "simple template",
			tmpl:        "Hello, {{.Name}}!",
			lang:        "test",
			data:        struct{ Name string }{Name: "World"},
			expected:    "Hello, World!",
			expectError: false,
		},
		{
			name:        "invalid template",
			tmpl:        "Hello, {{.InvalidField}}!",
			lang:        "test",
			data:        struct{ Name string }{Name: "World"},
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty template",
			tmpl:        "",
			lang:        "test",
			data:        struct{}{},
			expected:    "",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RenderToString(tt.tmpl, tt.lang, tt.data)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestWriteCodeToFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testoutput")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	codeMap := map[string][]byte{
		"file1.txt": []byte("content1"),
		"file2.txt": []byte("content2"),
	}

	t.Run("write multiple files", func(t *testing.T) {
		WriteCodeToFile(tempDir, codeMap)

		for filename, content := range codeMap {
			filepath := filepath.Join(tempDir, filename)
			data, err := os.ReadFile(filepath)
			assert.NoError(t, err)
			assert.Equal(t, content, data)
		}
	})

	t.Run("invalid directory", func(t *testing.T) {
		// This should not panic and should print error messages
		WriteCodeToFile("/nonexistent/directory", codeMap)
	})
}
