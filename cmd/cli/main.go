package main

import (
	"fmt"
	"os"

	"github.com/dhruvsaxena1998/rel/cmd/cli/commands"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "rel",
	Short: "REL command line interface",
	Long:  "A CLI tool for working with REL (Rule Expression Language)",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: "0.1.0",
}

func init() {
	RootCommand.AddCommand(commands.TranslateCommand)
}

func main() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
