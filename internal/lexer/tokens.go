package lexer

import (
	"fmt"
	"slices"
)

type TokenType int
type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF TokenType = iota
	ILLEGAL
	NULL
	TRUE
	FALSE
	IDENTIFIER
	VARIABLE

	NUMBER
	STRING

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	ASSIGN
	EQ
	NOT
	NOT_NOT
	STRICT_EQ
	NOT_EQ
	NOT_STRICT_EQ
	NOT_IN

	GT
	LT
	GTE
	LTE
	BETWEEN

	COLON
	SEMICOLON
	COMMA

	// Reserved keywords
	LET
	CONST
	IN
	AND
	OR
	IF
	ELSE
)

var ReservedKeywords map[string]TokenType = map[string]TokenType{
	"true":    TRUE,
	"false":   FALSE,
	"null":    NULL,
	"let":     LET,
	"const":   CONST,
	"if":      IF,
	"else":    ELSE,
	"not":     NOT,
	"not not": NOT_NOT,
	"not in":  NOT_IN,
	"in":      IN,
	"and":     AND,
	"or":      OR,
	"between": BETWEEN,
}

func (token Token) IsOneOfMany(expectedTokens ...TokenType) bool {
	return slices.Contains(expectedTokens, token.Type)
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func TokenTypeString(tokentype TokenType) string {
	switch tokentype {
	case EOF:
		return "eof"
	case ILLEGAL:
		return "illegal"
	case NULL:
		return "null"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case IDENTIFIER:
		return "identifier"
	case VARIABLE:
		return "variable"

	case NUMBER:
		return "number"
	case STRING:
		return "string"

	case LPAREN:
		return "lparen"
	case RPAREN:
		return "rparen"
	case LBRACE:
		return "lbrace"
	case RBRACE:
		return "rbrace"
	case LBRACKET:
		return "lbracket"
	case RBRACKET:
		return "rbracket"

	case ASSIGN:
		return "assign"
	case EQ:
		return "eq"
	case NOT:
		return "not"
	case NOT_NOT:
		return "not not"
	case NOT_IN:
		return "not in"
	case STRICT_EQ:
		return "strict_eq"
	case NOT_EQ:
		return "not_eq"
	case NOT_STRICT_EQ:
		return "not_strict_eq"

	case GT:
		return "gt"
	case LT:
		return "lt"
	case GTE:
		return "gte"
	case LTE:
		return "lte"
	case BETWEEN:
		return "between"

	case COLON:
		return "colon"
	case SEMICOLON:
		return "semicolon"
	case COMMA:
		return "comma"

	case LET:
		return "let"
	case CONST:
		return "const"
	case IN:
		return "in"
	case AND:
		return "and"
	case OR:
		return "or"
	case IF:
		return "if"
	case ELSE:
		return "else"

	default:
		return "unknown"
	}
}

func (t Token) Debug() {
	if t.IsOneOfMany(IDENTIFIER, NUMBER, STRING, VARIABLE) {
		fmt.Printf("%s (%s)\n", TokenTypeString(t.Type), t.Literal)
	} else {
		fmt.Printf("%s ()\n", TokenTypeString(t.Type))
	}
}
