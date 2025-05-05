package parser_test

import (
	"testing"

	"github.com/dhruvsaxena1998/rel/internal/parser"
)

func TestLookupKeyword(t *testing.T) {
	tests := []struct {
		input    string
		expected parser.TokenType
	}{
		{"AND", parser.AND},
		{"OR", parser.OR},
		{"IN", parser.IN},
		{"NOT", parser.BANG},
		{"LOG", parser.LOG},
	}

	for _, tt := range tests {
		if token := parser.LookupIdentifier(tt.input); token != tt.expected {
			t.Fatalf("LookupIdentifier(%q) = %v, expected %v", tt.input, token, tt.expected)
		}
	}
}

func TestLookupIdentifiers(t *testing.T) {
	tests := []string{"foo", "exp1", "exp2", "hello", ""}
	for _, input := range tests {
		if tok := parser.LookupIdentifier(input); tok != parser.IDENTIFIER {
			t.Fatalf("LookupIdent(%q) = %q; want IDENT", input, tok)
		}
	}
}
