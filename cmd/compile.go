package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xinchentechnote/fin-protoc/internal/model"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

var (
	luaOutput  string
	rsOutput   string
	goOutput   string
	javaOutput string
	pyOutput   string
	cppOutput  string
)

// Compile DSL to code in various target languages
func Compile(input string, outputs map[string]string) error {
	result, err := parser.ParseFile(input)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	binModel := result.(*model.BinaryModel)
	config := parser.NewGeneratorConfig(binModel.Options)

	generators := []struct {
		lang string
		path string
		gen  func() (map[string][]byte, error)
	}{
		{"Lua", outputs["lua"], func() (map[string][]byte, error) {
			return parser.NewLuaWspGenerator(config, binModel).Generate(binModel)
		}},
		{"Rust", outputs["rust"], func() (map[string][]byte, error) {
			return parser.NewRustGenerator(config, binModel).Generate(binModel)
		}},
		{"Go", outputs["go"], func() (map[string][]byte, error) {
			return parser.NewGoGenerator(config, binModel).Generate(binModel)
		}},
		{"Java", outputs["java"], func() (map[string][]byte, error) {
			return parser.NewJavaGenerator(config, binModel).Generate(binModel)
		}},
		{"Python", outputs["python"], func() (map[string][]byte, error) {
			return parser.NewPythonGenerator(config, binModel).Generate(binModel)
		}},
		{"C++", outputs["cpp"], func() (map[string][]byte, error) {
			return parser.NewCppGenerator(config, binModel).Generate(binModel)
		}},
	}

	for _, g := range generators {
		if g.path == "" {
			continue
		}
		codeMap, err := g.gen()
		if err != nil {
			return fmt.Errorf("failed to generate %s code: %w", g.lang, err)
		}
		if err := parser.WriteCodeToFile(g.path, codeMap); err != nil {
			return fmt.Errorf("write error for %s: %w", g.lang, err)
		}
	}
	return nil
}

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile DSL into multiple target languages",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputs := map[string]string{
			"lua":    luaOutput,
			"rust":   rsOutput,
			"go":     goOutput,
			"java":   javaOutput,
			"python": pyOutput,
			"cpp":    cppOutput,
		}
		return Compile(file, outputs)
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&rsOutput, "rs_output", "r", "", "Rust output path")
	compileCmd.Flags().StringVarP(&luaOutput, "lua_output", "l", "", "Lua output path")
	compileCmd.Flags().StringVarP(&javaOutput, "java_output", "j", "", "Java output path")
	compileCmd.Flags().StringVarP(&goOutput, "go_output", "g", "", "Go output path")
	compileCmd.Flags().StringVarP(&cppOutput, "cpp_output", "c", "", "C++ output path")
	compileCmd.Flags().StringVarP(&pyOutput, "py_output", "p", "", "Python output path")
	rootCmd.AddCommand(compileCmd)
}
