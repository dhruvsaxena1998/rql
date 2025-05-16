package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dhruvsaxena1998/rel/internal/lexer"
	"github.com/dhruvsaxena1998/rel/internal/parser"
	translator "github.com/dhruvsaxena1998/rel/internal/translator/jsonlogic"
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
)

var (
	prettyPrint bool
	debug       bool
	inlineInput string
	inputFile   string
	outFile     string
)

var TranslateCommand = &cobra.Command{
	Use:   "translate",
	Short: "Translate a REL",
	Long:  "Translate a REL expression to a JSON expression",
	RunE: func(cmd *cobra.Command, args []string) error {
		return handleTranslateCommand(cmd, args)
	},
}

func init() {
	TranslateCommand.Flags().BoolVarP(
		&prettyPrint,
		"pretty",
		"p",
		false,
		"pretty print the output",
	)

	TranslateCommand.Flags().BoolVar(
		&debug,
		"debug",
		false,
		"log debug statements",
	)

	TranslateCommand.Flags().StringVarP(
		&inlineInput,
		"inline",
		"i",
		"",
		"inline input",
	)
	TranslateCommand.Flags().StringVarP(
		&inputFile,
		"file",
		"f",
		"",
		"file path",
	)

	TranslateCommand.MarkFlagsMutuallyExclusive("inline", "file")

	TranslateCommand.Flags().StringVarP(
		&outFile,
		"output",
		"o",
		"",
		"output file path",
	)
}

func handleTranslateCommand(cmd *cobra.Command, args []string) error {
	input, err := handleInput()
	if err != nil {
		return err
	}

	tokens := lexer.Tokenize(input)
	ast := parser.Parse(tokens)

	if debug {
		litter.Dump(tokens)
		litter.Dump(ast)
	}
	jsonlogic := translator.TranslateToJSONLogic(ast)

	err = handleOutput(jsonlogic)
	if err != nil {
		return err
	}

	return nil
}

// handleInput retrieves the input content from either an inline parameter or a file.
// It validates that exactly one input source is provided and reads file content when necessary.
//
// Returns:
//   - string: The input content as a string
//   - error: Returns nil on success, or an error if validation fails or file reading fails
func handleInput() (string, error) {
	if inputFile == "" && inlineInput == "" || inputFile != "" && inlineInput != "" {
		return "", fmt.Errorf("input file or inline input must be provided")
	}

	if inlineInput != "" {
		return inlineInput, nil
	}

	// Read the content of the file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read file : %v", err)
	}

	return string(content), nil
}

// handleOutput manages the output of the translate command by writing the result to either a file or stdout.
// It handles JSON encoding with options for pretty printing and HTML escaping.
//
// Parameters:
//   - output interface{}: The data to be encoded as JSON
//
// Returns:
//   - error: Returns nil on success, or an error if file creation or JSON encoding fails
func handleOutput(output any) error {
	// Initialize file pointer for output destination
	var out *os.File

	// If outFile flag is set, create and write to file, otherwise use stdout
	if outFile != "" {
		var err error
		out, err = os.Create(outFile)
		if err != nil {
			return fmt.Errorf("failed to create output file : %v", err)
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}

	// Create a new JSON encoder for the output destination
	enc := json.NewEncoder(out)
	// Disable HTML escaping in the output JSON
	enc.SetEscapeHTML(false)

	// Enable pretty printing with 2-space indentation if the pretty flag is set
	if prettyPrint {
		enc.SetIndent("", "  ")
	}

	// Encode the output data as JSON and handle any encoding errors
	if err := enc.Encode(output); err != nil {
		return fmt.Errorf("JSON encoding error : %v", err)
	}

	return nil
}
