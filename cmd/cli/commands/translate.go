package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	inlineInput string
	outFile     string
	prettyPrint bool
)

func translateRelToJsonLogic(args []string) error {
	var input string
	if inlineInput != "" {
		input = inlineInput
	} else if len(args) == 1 {
		data, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("failed to read file: %s: %w", args[0], err)
		}

		input = string(data)
	} else {
		return fmt.Errorf("input file or --inline expression must be provided")
	}

	fmt.Println(input)
	fmt.Printf("Output: %q", outFile)

	return nil

}

var TranslateCommand = &cobra.Command{
	Use:   "translate [file] [flags]",
	Short: "Translate REL to JSONLogic",
	RunE: func(cmd *cobra.Command, args []string) error {
		return translateRelToJsonLogic(args)
	},
}

func init() {
	TranslateCommand.Flags().StringVarP(&inlineInput, "inline", "i", "", "Provide inline REL expression (reads from flag instead of file)")
	TranslateCommand.Flags().StringVarP(&outFile, "out", "o", "", "Output file path (defaults to stdout)")
	TranslateCommand.Flags().BoolVarP(&prettyPrint, "pretty", "p", false, "Pretty-print JSON output")
}
