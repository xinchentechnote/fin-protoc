package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xinchentechnote/fin-protoc/internal/model"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

var luaOutput string
var rsOutput string

// Compile DSL to rust, lua base wireshark plugin and so on.
func Compile(input string, luaPath string, rsPath string) {
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
}

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile DSL into Rust code",
	Run: func(cmd *cobra.Command, args []string) {
		Compile(file, luaOutput, rsOutput)
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&rsOutput, "rs_output", "r", "", "output path")
	compileCmd.Flags().StringVarP(&luaOutput, "lua_output", "l", "", "output path")
	rootCmd.AddCommand(compileCmd)
}
