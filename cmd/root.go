package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fin-protoc",
	Short: "A CLI for compiling/formatting packet DSL",
	// Runtime errors are printed to stderr by cobra itself; suppress the
	// usage dump so a failed run only shows the error.
	SilenceUsage: true,
}

// Execute runs the root command.
func Execute() {
	fmt.Println(len(os.Args))
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "--help")
	} else if len(os.Args) > 1 && !isSubcommand(os.Args[1]) {
		// insert "compile"
		os.Args = append([]string{os.Args[0], "compile"}, os.Args[1:]...)
	}
	// cobra already reports the error on stderr; only translate it to the
	// process exit code here.
	if err := rootCmd.Execute(); err != nil {
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
