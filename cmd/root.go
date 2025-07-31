package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fin-protoc",
	Short: "A CLI for compiling/formatting packet DSL",
}

// Execute runs the root command.
func Execute() {
	if len(os.Args) == 1 || (len(os.Args) > 1 && !isSubcommand(os.Args[1])) {
		// insert "compile"
		os.Args = append([]string{os.Args[0], "compile"}, os.Args[1:]...)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isSubcommand(arg string) bool {
	for _, c := range rootCmd.Commands() {
		if c.Name() == arg {
			return true
		}
	}
	return false
}
