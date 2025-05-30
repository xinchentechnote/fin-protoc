package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xinchentechnote/fin-protoc/internal/parser"
)

var file string
var dsl string

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format packet DSL string",
	Args:  cobra.NoArgs, // No arguments expected in the command
	Run: func(cmd *cobra.Command, args []string) {
		var inputDsl string
		var err error

		// Check if dsl string is provided
		if dsl != "" {
			inputDsl = dsl
		} else if file != "" {
			// If file path is provided, read from the file
			data, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("Error reading file:", err)
				os.Exit(1)
			}
			inputDsl = string(data)
		} else {
			fmt.Println("Please provide a DSL string or a file path")
			os.Exit(1)
		}

		// Now format the DSL string
		result, err := parser.FormatPacketDsl(inputDsl)
		if err != nil {
			fmt.Println("Error formatting DSL:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	},
}

func init() {
	formatCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	formatCmd.Flags().StringVarP(&dsl, "dsl", "d", "", "DSL string to format")
	rootCmd.AddCommand(formatCmd)
}
