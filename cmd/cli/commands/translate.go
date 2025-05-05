package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	fileInput   string
	inlineInput string
	outFile     string
	prettyPrint bool
)

func inputAsString() (string, error) {
	var input string

	if inlineInput != "" {
		input = inlineInput
	} else {
		data, err := os.ReadFile(fileInput)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %w", fileInput, err)
		}
		input = string(data)
	}

	return input, nil
}

var TranslateCommand = &cobra.Command{
	Use:   "translate [flags]",
	Short: "Translate REL to JSONLogic",
	RunE: func(cmd *cobra.Command, args []string) error {
		if (fileInput == "" && inlineInput == "") || (fileInput != "" && inlineInput != "") {
			return fmt.Errorf("you must specify exactly one of --file or --inline")
		}

		input, err := inputAsString()
		if err != nil {
			return err
		}

		fmt.Printf("Translating REL to JSONLogic: %s\n", input)
		return nil
	},
}

func init() {
	TranslateCommand.Flags().StringVarP(&fileInput, "file", "f", "", "Path to REL input file")
	TranslateCommand.Flags().StringVarP(&inlineInput, "inline", "i", "", "Provide inline REL expression")
	TranslateCommand.Flags().StringVarP(&outFile, "out", "o", "", "Output file path (defaults to stdout)")
	TranslateCommand.Flags().BoolVarP(&prettyPrint, "pretty", "p", false, "Pretty-print JSON output")

	TranslateCommand.MarkFlagsMutuallyExclusive("file", "inline")
}
