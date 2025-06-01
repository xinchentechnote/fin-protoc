package main

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/xinchentechnote/fin-protoc/internal/model"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

var output string

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile DSL into Rust code",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. initialize generator
		generator := parser.RustGenerator{}

		// 2. parse DSL file
		result, _ := parser.ParseFile(file)
		packets := result.([]model.Packet)
		// 3. generate code for each packet
		for _, pkt := range packets {
			code := generator.GenerateCode(pkt)
			outputFile := fmt.Sprintf("%s/%s%s", output, strcase.ToSnake(pkt.Name), generator.FileExtension())
			f, _ := os.Create(outputFile)
			defer f.Close()
			_, _ = f.WriteString(code)
			code = generator.GenerateTestCode(pkt)
			_, _ = f.WriteString("\n")
			_, _ = f.WriteString(code)
		}
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&output, "output", "o", "", "output path")
	rootCmd.AddCommand(compileCmd)
}
