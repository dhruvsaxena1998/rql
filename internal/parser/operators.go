package parser

import "fmt"

// OperatorPrecedence defines precedence levels for operators
type OperatorPrecedence int

const (
	LOWEST     OperatorPrecedence = iota
	LOGICAL                       // AND, OR
	EQUALS                        // ==, !=, ===, !==
	COMPARISON                    // >, <, >=, <=, IN
	PREFIX                        // NOT, !
	CALL                          // function calls
)

// Map tokens to precedence levels
var precedences = map[TokenType]OperatorPrecedence{
	EQ:         EQUALS,
	STRICT_EQ:  EQUALS,
	NOT_EQ:     EQUALS,
	STRICT_NOT: EQUALS,
	LT:         COMPARISON,
	GT:         COMPARISON,
	LTE:        COMPARISON,
	GTE:        COMPARISON,
	IN:         COMPARISON,
	AND:        LOGICAL,
	OR:         LOGICAL,
}

// Helper functions for token checking
func (p *Parser) peekPrecedence() OperatorPrecedence {
	if precedence, ok := precedences[p.peekToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() OperatorPrecedence {
	if precedence, ok := precedences[p.currentToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

// isComparisonOperator checks if the token type is a comparison operator
func (p *Parser) isComparisonOperator(tokenType TokenType) bool {
	switch tokenType {
	case ASSIGN, EQ, STRICT_EQ, NOT_EQ, STRICT_NOT, GT, LT, GTE, LTE:
		return true
	default:
		return false
	}
}

// Helper methods for token checking and error handling
func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.addError(fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type))
	return false
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) currentTokenIs(t TokenType) bool {
	return p.currentToken.Type == t
}
