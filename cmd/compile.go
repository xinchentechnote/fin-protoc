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

		// 1. parse DSL file
		result, _ := parser.ParseFile(file)
		packets := result.([]model.Packet)
		packetsMap := make(map[string]model.Packet)
		for _, pkt := range packets {
			packetsMap[pkt.Name] = pkt
		}
		// 2. initialize generator
		generator := parser.RustGenerator{
			Packets: packetsMap,
		}

		// 3. generate code for each packet
		libFilePath := fmt.Sprintf("%s/lib.rs", output)
		libFile, _ := os.Create(libFilePath)
		for _, pkt := range packets {
			outputFile := fmt.Sprintf("%s/%s%s", output, strcase.ToSnake(pkt.Name), generator.FileExtension())
			f, _ := os.Create(outputFile)
			defer f.Close()
			code := generator.GenerateCode(pkt)
			_, _ = f.WriteString(code)
			fmt.Println("Generated code for packet:", outputFile)
			_, _ = libFile.WriteString(fmt.Sprintf("pub mod %s;\n", strcase.ToSnake(pkt.Name)))
		}
	},
}

func init() {
	compileCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	compileCmd.Flags().StringVarP(&output, "output", "o", "", "output path")
	rootCmd.AddCommand(compileCmd)
}
