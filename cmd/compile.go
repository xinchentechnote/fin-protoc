package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xinchentechnote/fin-protoc/internal/model"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

var luaOutput string
var rsOutput string

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile DSL into Rust code",
	Run: func(cmd *cobra.Command, args []string) {

		// parse DSL file
		result, _ := parser.ParseFile(file)
		binModel := result.(*model.BinaryModel)
		config := parser.NewGeneratorConfig(binModel.Options)
		if luaOutput != "" {
			// generate code for lua base wireshatk plugin
			luaGen := parser.NewLuaWspGenerator(config, binModel)
			codeMap, err := luaGen.Generate(binModel)
			if nil != err {
				fmt.Println("gen lua code err.")
			}
			parser.WriteCodeToFile(luaOutput, codeMap)
		}
		if rsOutput != "" {
			// initialize generator
			rustGen := parser.NewRustGenerator(config, binModel)
			codeMap, err := rustGen.Generate(binModel)
			if nil != err {
				fmt.Println("gen rust code err.")
			}
			parser.WriteCodeToFile(rsOutput, codeMap)
		}
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&rsOutput, "rs_output", "r", "", "output path")
	compileCmd.Flags().StringVarP(&luaOutput, "lua_output", "l", "", "output path")
	rootCmd.AddCommand(compileCmd)
}
