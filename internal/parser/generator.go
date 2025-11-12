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
	ListLenPrefixLenType   string         //default u16
	StringLenPrefixLenType string         //default u16
	JavaPackage            string         //default empty
	GoPackage              string         //default empty
	GoModule               string         //default empty
	LittleEndian           bool           //default false
	Padding                *model.Padding //default nil
}

// NewGeneratorConfig new config
func NewGeneratorConfig(options map[string]string) *GeneratorConfig {
	config := &GeneratorConfig{
		// Set defaults first
		ListLenPrefixLenType:   "u16",
		StringLenPrefixLenType: "u16",
		JavaPackage:            "",
		GoPackage:              "",
		GoModule:               "",
		LittleEndian:           false,
		Padding: &model.Padding{
			PadLeft: false,
			PadChar: "' '",
		},
	}

	// Override with provided options
	if val, ok := options[model.ArrayPrefixLenType]; ok {
		config.ListLenPrefixLenType = val
	}
	if val, ok := options[model.StringPrefixLenType]; ok {
		config.StringLenPrefixLenType = val
	}
	if val, ok := options[model.JavaPackage]; ok {
		config.JavaPackage = val
	}
	if val, ok := options[model.GoPackage]; ok {
		config.GoPackage = val
	}
	if val, ok := options[model.LittleEndian]; ok {
		config.LittleEndian = strings.ToLower(val) == "true"
	}
	if val, ok := options[model.GoModule]; ok {
		config.GoModule = val
	}
	fromLeft := false
	padChar := " "
	if val, ok := options[model.FixedStringPadFromLeft]; ok {
		fromLeft = strings.ToLower(val) == "true"
	}
	if val, ok := options[model.FixedStringPadChar]; ok {
		padChar = val
	}
	if fromLeft || padChar != " " {
		config.Padding = &model.Padding{
			PadLeft: fromLeft,
			PadChar: padChar,
		}
	}
	return config
}
