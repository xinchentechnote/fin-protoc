package parser

import (
	"strings"

	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// Generator interface for extensibility
type Generator interface {
	Generate(binModel *model.BinaryModel) (map[string][]byte, error)
}

// GeneratorConfig config
type GeneratorConfig struct {
	ListLenPrefixLenType   string //default u16
	StringLenPrefixLenType string //default u16
	LittleEndian           bool   //default false
}

// NewGeneratorConfig new config
func NewGeneratorConfig(options map[string]string) *GeneratorConfig {
	config := &GeneratorConfig{
		// Set defaults first
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u16",
		LittleEndian:           false,
	}

	// Override with provided options
	if val, ok := options["ListLenPrefixLenType"]; ok {
		config.ListLenPrefixLenType = val
	}
	if val, ok := options["StringLenPrefixLenType"]; ok {
		config.StringLenPrefixLenType = val
	}
	if val, ok := options["LittleEndian"]; ok {
		config.LittleEndian = strings.ToLower(val) == "true"
	}

	return config
}
