package main

import (
	"fmt"
	"os"

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

		// 1. parse DSL file
		result, _ := parser.ParseFile(file)
		binModel := result.(*model.BinaryModel)
		// packetsMap := binModel.Packets

		// 3. generate code for each packet
		luaGen := parser.NewLuaWspGenerator(parser.GeneratorConfig{
			ListLenPrefix:   "u16",
			StringLenPrefix: "u16",
		})
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

		// 2. initialize generator
		// generator := parser.RustGenerator{
		// 	ListLenPrefix:   "u16",
		// 	StringLenPrefix: "u16",
		// 	Packets:         packetsMap,
		// }
		// libFilePath := fmt.Sprintf("%s/lib.rs", output)
		// libFile, _ := os.Create(libFilePath)
		// for _, pkt := range packetsMap {
		// 	outputFile := fmt.Sprintf("%s/%s%s", output, strcase.ToSnake(pkt.Name), generator.FileExtension())
		// 	f, _ := os.Create(outputFile)
		// 	defer f.Close()
		// 	code := generator.GenerateCode(pkt)
		// 	_, _ = f.WriteString(code)
		// 	fmt.Println("Generated code for packet:", outputFile)
		// 	_, _ = libFile.WriteString(fmt.Sprintf("pub mod %s;\n", strcase.ToSnake(pkt.Name)))
		// }
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&rsOutput, "rs_output", "r", "", "output path")
	compileCmd.Flags().StringVarP(&luaOutput, "lua_output", "l", "", "output path")
	rootCmd.AddCommand(compileCmd)
}
