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
	RunE: func(cmd *cobra.Command, args []string) error {
		var inputDsl string

		// Check if dsl string is provided
		if dsl != "" {
			inputDsl = dsl
		} else if file != "" {
			// If file path is provided, read from the file
			data, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("read file: %w", err)
			}
			inputDsl = string(data)
		} else {
			return fmt.Errorf("please provide a DSL string or a file path")
		}

		// Now format the DSL string
		result, err := parser.FormatPacketDsl(inputDsl)
		if err != nil {
			return fmt.Errorf("format DSL: %w", err)
		}
		if file != "" {
			if err := os.WriteFile(file, []byte(result), 0644); err != nil {
				return fmt.Errorf("write formatted DSL: %w", err)
			}
		} else {
			fmt.Println(result)
		}
		return nil
	},
}

func init() {
	formatCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the DSL file")
	formatCmd.Flags().StringVarP(&dsl, "dsl", "d", "", "DSL string to format")
	rootCmd.AddCommand(formatCmd)
}
