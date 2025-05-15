package cmd

import (
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile [dsl]",
	Short: "Compile DSL into target format (e.g., JSON, Go struct)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//TODO
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
}
