package parser

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

const (
	// Special Tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers
	VARIABLE   TokenType = "VARIABLE"
	IDENTIFIER TokenType = "IDENTIFIER"
	NUMBER     TokenType = "NUMBER"
	STRING     TokenType = "STRING"

	// Delimiters
	ASSIGN    TokenType = "="
	SEMICOLON TokenType = ";"
	COMMA     TokenType = ","

	// Brackets
	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"

	// Operators
	EQ         TokenType = "=="
	STRICT_EQ  TokenType = "==="
	NOT_EQ     TokenType = "!="
	STRICT_NOT TokenType = "!=="
	BANG       TokenType = "!"
	GT         TokenType = ">"
	LT         TokenType = "<"
	GTE        TokenType = ">="
	LTE        TokenType = "<="

	// Keywords
	AND TokenType = "AND"
	OR  TokenType = "OR"
	IN  TokenType = "IN"
	LOG TokenType = "LOG"
)

var keywords = map[string]TokenType{
	"AND": AND,
	"OR":  OR,
	"IN":  IN,
	"NOT": BANG,
	"LOG": LOG,
}

func LookupIdentifier(ident string) TokenType {
	upper := strings.ToUpper(ident)
	if tok, ok := keywords[upper]; ok {
		return tok
	}
	return IDENTIFIER
}
