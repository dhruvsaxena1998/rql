package parser

import (
	"github.com/dhruvsaxena1998/rel/internal/ast"
	"github.com/dhruvsaxena1998/rel/internal/lexer"
)

type bindingPower int

const (
	defaultBindingPower bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

// statementHandler represents a function that parses a statement
type statementHandler func(p *parser) ast.Statement

// nudHandler represents a function that parses a prefix expression (null denotation)
type nudHandler func(p *parser) ast.Expression

// ledHandler represents a function that parses an infix expression (left denotation)
type ledHandler func(p *parser, left ast.Expression, bp bindingPower) ast.Expression

// statementLookup maps token types to their statement handlers
type statementLookup map[lexer.TokenType]statementHandler

// nudLookup maps token types to their null denotation handlers
type nudLookup map[lexer.TokenType]nudHandler

// ledLookup maps token types to their left denotation handlers
type ledLookup map[lexer.TokenType]ledHandler

// bindingPowerLookup maps token types to their binding power levels
type bindingPowerLookup map[lexer.TokenType]bindingPower

var statementLU = statementLookup{}
var nudLU = nudLookup{}
var ledLU = ledLookup{}
var bpLU = bindingPowerLookup{}

// led registers a left denotation handler with its binding power
func registerLedHandler(tokenType lexer.TokenType, bp bindingPower, handler ledHandler) {
	bpLU[tokenType] = bp
	ledLU[tokenType] = handler
}

// nud registers a null denotation handler
func registerNudHandler(tokenType lexer.TokenType, handler nudHandler) {
	nudLU[tokenType] = handler
}

// stmt registers a statement handler
func registerStatementHandler(tokenType lexer.TokenType, handler statementHandler) {
	bpLU[tokenType] = defaultBindingPower
	statementLU[tokenType] = handler
}

func createTokenLookups() {
	// Unary
	registerLedHandler(lexer.ASSIGN, assignment, parseAssignmentExpression)
	registerNudHandler(lexer.NOT, parsePrefixExpression)
	registerNudHandler(lexer.NOT_NOT, parsePrefixExpression)

	// Logical
	registerLedHandler(lexer.AND, logical, parseBinaryExpression)
	registerLedHandler(lexer.OR, logical, parseBinaryExpression)

	// Relational
	registerLedHandler(lexer.LT, relational, parseBinaryExpression)
	registerLedHandler(lexer.LTE, relational, parseBinaryExpression)
	registerLedHandler(lexer.GT, relational, parseBinaryExpression)
	registerLedHandler(lexer.GTE, relational, parseBinaryExpression)
	registerLedHandler(lexer.EQ, relational, parseBinaryExpression)
	registerLedHandler(lexer.NOT_EQ, relational, parseBinaryExpression)
	registerLedHandler(lexer.STRICT_EQ, relational, parseBinaryExpression)
	registerLedHandler(lexer.NOT_STRICT_EQ, relational, parseBinaryExpression)
	registerLedHandler(lexer.IN, relational, parseBinaryExpression)

	// Between
	registerLedHandler(lexer.BETWEEN, relational, parseBetweenExpression)

	// Additive & Multiplicative

	// Literals and Symbols
	registerNudHandler(lexer.NUMBER, parsePrimaryExpression)
	registerNudHandler(lexer.STRING, parsePrimaryExpression)
	registerNudHandler(lexer.IDENTIFIER, parsePrimaryExpression)
	registerNudHandler(lexer.VARIABLE, parsePrimaryExpression)

	// Bool
	registerNudHandler(lexer.TRUE, parsePrimaryExpression)
	registerNudHandler(lexer.FALSE, parsePrimaryExpression)

	// If Expression
	registerNudHandler(lexer.IF, parseIfExpression)

	// Grouping
	registerNudHandler(lexer.LPAREN, parseGroupExpression)

	// Member
	registerNudHandler(lexer.LBRACKET, parseArrayExpression)

	// Statements
	registerStatementHandler(lexer.CONST, parseVarDeclarationStatement)
	registerStatementHandler(lexer.LET, parseVarDeclarationStatement)
}
