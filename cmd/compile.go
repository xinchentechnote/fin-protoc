package main

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
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
			luaGen := parser.NewLuaWspGenerator(config)
			codeMap, err := luaGen.Generate(binModel)
			if nil != err {
				fmt.Println("gen lua err.")
			}
			for name, datas := range codeMap {
				filepath := luaOutput + "/" + name
				f, _ := os.Create(filepath)
				defer f.Close()
				_, _ = f.Write(datas)
				fmt.Println("Generated code for packet:", filepath)
			}
		}
		if rsOutput != "" {
			// initialize generator
			packetsMap := binModel.PacketsMap
			generator := parser.RustGenerator{
				Config:  config,
				Packets: packetsMap,
			}
			// generate code for each packet
			libFilePath := fmt.Sprintf("%s/lib.rs", rsOutput)
			libFile, _ := os.Create(libFilePath)
			for _, pkt := range packetsMap {
				outputFile := fmt.Sprintf("%s/%s%s", rsOutput, strcase.ToSnake(pkt.Name), generator.FileExtension())
				f, _ := os.Create(outputFile)
				defer f.Close()
				code := generator.GenerateCode(pkt)
				_, _ = f.WriteString(code)
				fmt.Println("Generated code for packet:", outputFile)
				_, _ = libFile.WriteString(fmt.Sprintf("pub mod %s;\n", strcase.ToSnake(pkt.Name)))
			}
		}
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&rsOutput, "rs_output", "r", "", "output path")
	compileCmd.Flags().StringVarP(&luaOutput, "lua_output", "l", "", "output path")
	rootCmd.AddCommand(compileCmd)
}
