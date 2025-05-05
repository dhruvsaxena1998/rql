package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dhruvsaxena1998/rel/internal/parser"
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

		// Create lexer and parser
		lexer := parser.NewLexer(input)
		p := parser.NewParser(lexer)

		// Parse expression
		expression := p.ParseExpression()
		if expression == nil {
			return fmt.Errorf("parsing error: %v", p.Errors())
		}

		// Transform to JSONLogic
		jsonLogic, err := parser.Transform(expression)
		if err != nil {
			return fmt.Errorf("transform error: %v", err)
		}

		// Write output
		var out *os.File
		if outFile != "" {
			var err error
			out, err = os.Create(outFile)
			if err != nil {
				return fmt.Errorf("failed to create output file: %v", err)
			}
			defer out.Close()
		} else {
			out = os.Stdout
		}

		// Create encoder and disable HTML escaping
		enc := json.NewEncoder(out)
		enc.SetEscapeHTML(false)
		if prettyPrint {
			enc.SetIndent("", "  ")
		}

		// Encode JSONLogic directly
		if err := enc.Encode(jsonLogic); err != nil {
			return fmt.Errorf("JSON encoding error: %v", err)
		}

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
