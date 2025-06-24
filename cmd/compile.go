package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xinchentechnote/fin-protoc/internal/model"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

var luaOutput string
var rsOutput string
var goOutput string
var javaOutput string
var pyOutput string
var cppOutput string

// Compile DSL to rust, lua base wireshark plugin and so on.
func Compile(input string, luaPath string, rsPath string, goOutput string, javaOutput string, pyOutput string, cppOutput string) {
	// parse DSL file
	result, _ := parser.ParseFile(file)
	binModel := result.(*model.BinaryModel)
	config := parser.NewGeneratorConfig(binModel.Options)
	if luaPath != "" {
		// generate code for lua base wireshatk plugin
		luaGen := parser.NewLuaWspGenerator(config, binModel)
		codeMap, err := luaGen.Generate(binModel)
		if nil != err {
			fmt.Println("gen lua code err.")
		}
		parser.WriteCodeToFile(rsPath, codeMap)
	}
	if rsPath != "" {
		// initialize generator
		rustGen := parser.NewRustGenerator(config, binModel)
		codeMap, err := rustGen.Generate(binModel)
		if nil != err {
			fmt.Println("gen rust code err.")
		}
		parser.WriteCodeToFile(rsPath, codeMap)
	}
	if javaOutput != "" {
		// generate code for java
		javaGen := parser.NewJavaGenerator(config, binModel)
		codeMap, err := javaGen.Generate(binModel)
		if nil != err {
			fmt.Println("gen java code err.")
		}
		parser.WriteCodeToFile(javaOutput, codeMap)
	}
}

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile DSL into Rust code",
	Run: func(cmd *cobra.Command, args []string) {
		Compile(file, luaOutput, rsOutput, goOutput, javaOutput, pyOutput, cppOutput)
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&rsOutput, "rs_output", "r", "", "rust output path")
	compileCmd.Flags().StringVarP(&luaOutput, "lua_output", "l", "", "lua output path")
	compileCmd.Flags().StringVarP(&javaOutput, "java_output", "j", "", "java output path")
	compileCmd.Flags().StringVarP(&goOutput, "go_output", "g", "", "go output path")
	compileCmd.Flags().StringVarP(&cppOutput, "cpp_output", "c", "", "cpp output path")
	compileCmd.Flags().StringVarP(&pyOutput, "py_output", "p", "", "python output path")
	rootCmd.AddCommand(compileCmd)
}
