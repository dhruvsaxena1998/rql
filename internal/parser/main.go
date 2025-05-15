package parser

import (
	"fmt"

	"github.com/dhruvsaxena1998/rel/internal/ast"
	"github.com/dhruvsaxena1998/rel/internal/lexer"
)

type parser struct {
	tokens   []lexer.Token
	position int
	errors   []error
}

func CreateParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	return &parser{
		tokens:   tokens,
		position: 0,
		errors:   make([]error, 0),
	}
}

func Parse(tokens []lexer.Token) ast.BlockStatement {
	Body := &ast.LineStatement{}
	p := CreateParser(tokens)

	currentLineExpressions := make([]ast.Expression, 0)

	for p.hasTokens() {
		if p.currentToken().Type == lexer.SEMICOLON {
			if len(currentLineExpressions) > 0 {
				Body = &ast.LineStatement{
					Expressions: currentLineExpressions,
				}

				// reset for next line
				currentLineExpressions = make([]ast.Expression, 0)
			}
			p.advance()
		} else {
			statement := parseStatement(p)
			if expression, ok := statement.(*ast.ExpressionStatement); ok {
				currentLineExpressions = append(currentLineExpressions, expression.Expression)
			}
		}
	}

	if len(currentLineExpressions) > 0 {
		Body = &ast.LineStatement{
			Expressions: currentLineExpressions,
		}
	}

	return ast.BlockStatement{
		Body: Body,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.position]
}

func (p *parser) advance() lexer.Token {
	currentToken := p.currentToken()
	p.position++
	return currentToken
}

func (p *parser) hasTokens() bool {
	return p.position < len(p.tokens) && p.currentToken().Type != lexer.EOF
}

func (p *parser) expectOneOf(expectedTokenTypes ...lexer.TokenType) lexer.Token {
	token := p.currentToken()
	expectedTokenTypeStrings := make([]string, 0)
	for _, expectedTokenType := range expectedTokenTypes {
		expectedTokenTypeStrings = append(expectedTokenTypeStrings, lexer.TokenTypeString(expectedTokenType))
		if token.Type == expectedTokenType {
			return p.advance()
		}
	}

	panic(fmt.Errorf("expected one of %s, got %s", expectedTokenTypeStrings, lexer.TokenTypeString(token.Type)))
}

func (p *parser) expect(expectedTokenType lexer.TokenType) lexer.Token {
	return p.expectError(expectedTokenType, nil)
}

func (p *parser) expectError(expectedTokenType lexer.TokenType, err error) lexer.Token {
	token := p.currentToken()
	if token.Type != expectedTokenType {
		if err == nil {
			err = fmt.Errorf("expected %s, got %s", lexer.TokenTypeString(expectedTokenType), lexer.TokenTypeString(token.Type))
			p.errors = append(p.errors, err)
		}

		panic(err)
	}

	return p.advance()
}
