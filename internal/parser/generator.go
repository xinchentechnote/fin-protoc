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
	}

	// Override with provided options
	if val, ok := options["RepeatPreFixSizeType"]; ok {
		config.ListLenPrefixLenType = val
	}
	if val, ok := options["StringPreFixLenType"]; ok {
		config.StringLenPrefixLenType = val
	}
	if val, ok := options["JavaPackage"]; ok {
		config.JavaPackage = val
	}
	if val, ok := options["GoPackage"]; ok {
		config.GoPackage = val
	}
	if val, ok := options["LittleEndian"]; ok {
		config.LittleEndian = strings.ToLower(val) == "true"
	}
	if val, ok := options["GoModule"]; ok {
		config.GoModule = val
	}
	fromLeft := false
	padChar := " "
	if val, ok := options["FixStringPadFromLeft"]; ok {
		fromLeft = strings.ToLower(val) == "true"
	}
	if val, ok := options["FixStringPadChar"]; ok {
		padChar = val
	}
	if fromLeft || padChar != " " {
		config.Padding = &model.Padding{
			FromLeft: fromLeft,
			PadChar:  padChar,
		}
	}
	return config
}
