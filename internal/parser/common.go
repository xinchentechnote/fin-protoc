package parser

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	filepathpkg "path/filepath"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/xinchentechnote/fin-protoc/internal/grammar"
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// SyntaxErrorListener implements antlr.ErrorListener to collect syntax errors
type SyntaxErrorListener struct {
	Errors []model.SyntaxError
}

// NewSyntaxErrorListener creates a new SyntaxErrorListener
func NewSyntaxErrorListener() *SyntaxErrorListener {
	return &SyntaxErrorListener{
		Errors: make([]model.SyntaxError, 0),
	}
}

// ReportAmbiguity implements antlr.ErrorListener
func (s *SyntaxErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Typically you might want to ignore this in production parsers
	// but you could log it if needed
}

// ReportAttemptingFullContext implements antlr.ErrorListener
func (s *SyntaxErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Typically you might want to ignore this in production parsers
}

// ReportContextSensitivity implements antlr.ErrorListener
func (s *SyntaxErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	// Typically you might want to ignore this in production parsers
}

// SyntaxError implements antlr.ErrorListener
func (s *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	s.Errors = append(s.Errors, model.SyntaxError{
		Line:            line,
		Column:          column,
		Msg:             msg,
		OffendingSymbol: offendingSymbol,
	})
}

// HasErrors returns true if any syntax errors were collected
func (s *SyntaxErrorListener) HasErrors() bool {
	return len(s.Errors) > 0
}

// Error returns a formatted string containing all syntax errors
func (s *SyntaxErrorListener) Error() string {
	if !s.HasErrors() {
		return ""
	}

	var errStr string
	for _, err := range s.Errors {
		errStr += fmt.Sprintf("line %d:%d %s\n", err.Line, err.Column, err.Msg)
	}
	return errStr
}

// AddIndent4ln adds 4-space indent and a newline (similar to fmt.Println)
func AddIndent4ln(s string) string {
	return AddIndent4(s) + "\n"
}

// AddIndent4 add 4 Indent bydefault
func AddIndent4(s string) string {
	return AddIndent(s, 4)
}

// AddIndent add Indent
func AddIndent(s string, spaces int) string {
	indent := strings.Repeat(" ", spaces)
	return indent + strings.ReplaceAll(s, "\n", "\n"+indent)
}

// ReadFileToString reads the content of a file and returns it as a string.
func ReadFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NewPacketDslParserByFile creates a new PacketDslParser instance from a file.
func NewPacketDslParserByFile(filename string) (*gen.PacketDslParser, *antlr.CommonTokenStream, error) {
	data, err := ReadFileToString(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read file: %v", err)
	}
	return NewPacketDslParserByContent(data)
}

// NewPacketDslParserByContent creates a new PacketDslParser instance from a string.
func NewPacketDslParserByContent(data string) (*gen.PacketDslParser, *antlr.CommonTokenStream, error) {
	input := antlr.NewInputStream(string(data))

	lexer := gen.NewPacketDslLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := gen.NewPacketDslParser(stream)
	return parser, stream, nil
}

// ParseCharArrayType parse char[\d]
func ParseCharArrayType(fieldType string) (size string, ok bool) {
	pattern := `^char\[(\d+)\]$`
	matches := regexp.MustCompile(pattern).FindStringSubmatch(fieldType)
	if len(matches) == 2 {
		return matches[1], true
	}
	return "", false
}

// RenderToString render tmpl
func RenderToString(tmpl string, lang string, data interface{}) (string, error) {
	t := template.Must(template.New(lang).Parse(tmpl))

	var buf bytes.Buffer
	err := t.Execute(&buf, data) // Directly pass data instead of wrapping it
	if err != nil {
		return "", fmt.Errorf("template execution failed: %w", err)
	}

	return buf.String(), nil
}

// WriteCodeToFile write code to file
func WriteCodeToFile(path string, codeMap map[string][]byte) error {
	for name, datas := range codeMap {
		filepath := path + "/" + name
		dir := filepathpkg.Dir(filepath)
		err := os.MkdirAll(dir, 0755)
		if nil != err {
			return fmt.Errorf("create directory fail:%w", err)
		}
		f, err := os.Create(filepath)
		if nil != err {
			return fmt.Errorf("create file fail:%w", err)
		}
		defer f.Close()
		_, err = f.Write(datas)
		if nil != err {
			return fmt.Errorf("write file fail:%w", err)
		}
		fmt.Println("Generated code for packet:", filepath)
	}
	return nil
}
