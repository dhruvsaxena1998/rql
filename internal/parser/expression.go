package parser

import (
	"fmt"
	"strconv"

	"github.com/dhruvsaxena1998/rel/internal/ast"
	"github.com/dhruvsaxena1998/rel/internal/lexer"
)

func parseExpression(p *parser, bp bindingPower) ast.Expression {
	// First, parse the nud
	currentToken := p.currentToken()
	nudFn, exists := nudLU[currentToken.Type]

	if !exists {
		panic(fmt.Sprintf("nud handler expected for token %s\n", lexer.TokenTypeString(currentToken.Type)))
	}

	// Advance and parse ast left hand side
	left := nudFn(p)

	for bpLU[p.currentToken().Type] > bp {
		tokenType := p.currentToken().Type
		ledFn, exists := ledLU[tokenType]
		if !exists {
			panic(fmt.Sprintf("led handler expected for token %s\n", lexer.TokenTypeString(currentToken.Type)))
		}

		left = ledFn(p, left, bpLU[p.currentToken().Type])
	}

	return left
}

func parsePrimaryExpression(p *parser) ast.Expression {
	switch p.currentToken().Type {
	case lexer.NUMBER:
		number, err := strconv.ParseFloat(p.advance().Literal, 64)
		if err != nil {
			p.errors = append(p.errors, err)
			return nil
		}
		return &ast.NumberExpression{Value: number}
	case lexer.STRING:
		return &ast.StringExpression{Value: p.advance().Literal}
	case lexer.IDENTIFIER:
		return &ast.SymbolExpression{Value: p.advance().Literal}
	case lexer.VARIABLE:
		return &ast.VariableExpression{Value: p.advance().Literal}
	case lexer.TRUE:
		p.advance()
		return &ast.BooleanExpression{Value: true}
	case lexer.FALSE:
		p.advance()
		return &ast.BooleanExpression{Value: false}

	default:
		panic(fmt.Sprintf("cannot create primary expression for %s\n", lexer.TokenTypeString(p.currentToken().Type)))
	}
}

func parseBinaryExpression(p *parser, left ast.Expression, bp bindingPower) ast.Expression {
	operatorToken := p.advance()
	right := parseExpression(p, bp)

	return &ast.BinaryExpression{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parsePrefixExpression(p *parser) ast.Expression {
	operatorToken := p.advance()
	right := parseExpression(p, defaultBindingPower)
	return &ast.PrefixExpression{
		Operator: operatorToken,
		Right:    right,
	}
}

func parseGroupExpression(p *parser) ast.Expression {
	p.advance() // Consume the opening parenthesis token
	expression := parseExpression(p, defaultBindingPower)
	p.expect(lexer.RPAREN) // Consume and expect the closing parenthesis token

	return expression
}

func parseAssignmentExpression(p *parser, left ast.Expression, bp bindingPower) ast.Expression {
	operatorToken := p.advance()
	right := parseExpression(p, bp)
	return &ast.AssignmentExpression{
		Assigne:  left,
		Operator: operatorToken,
		Value:    right,
	}
}

func parseArrayExpression(p *parser) ast.Expression {
	var elements = make([]ast.Expression, 0)

	p.expect(lexer.LBRACKET)
	for p.hasTokens() && p.currentToken().Type != lexer.RBRACKET {
		elements = append(elements, parseExpression(p, logical))

		if !p.currentToken().IsOneOfMany(lexer.RBRACKET) {
			p.expect(lexer.COMMA)
		}
	}
	p.expect(lexer.RBRACKET)

	return &ast.ArrayLiteral{
		Elements: elements,
	}
}

func parseIfExpression(p *parser) ast.Expression {
	p.advance() // Consume the if token

	branches := make([]ast.IfBranch, 0)

	// Parse the first condition and branch
	condition := parseExpression(p, defaultBindingPower)
	p.expect(lexer.COLON)
	consequent := parseExpression(p, defaultBindingPower)

	branches = append(branches, ast.IfBranch{
		Condition:  condition,
		Consequent: consequent,
	})

	// Parse additional conditions and branches
	for p.currentToken().Type == lexer.COMMA {
		p.advance() // Consume the comma token

		// check if this is the else token
		if p.currentToken().Type == lexer.ELSE {
			p.advance() // Consume the else token
			p.expect(lexer.COLON)
			elseBranch := parseExpression(p, defaultBindingPower)

			return &ast.IfExpression{
				Branches: branches,
				Else:     elseBranch,
			}
		}

		// Otherwise, it's another condition branch
		condition := parseExpression(p, defaultBindingPower)
		p.expect(lexer.COLON)
		consequent := parseExpression(p, defaultBindingPower)

		branches = append(branches, ast.IfBranch{
			Condition:  condition,
			Consequent: consequent,
		})
	}

	return &ast.IfExpression{
		Branches: branches,
	}
}

func parseBetweenExpression(p *parser, left ast.Expression, bp bindingPower) ast.Expression {
	p.advance() // Consume the between token

	betweenExpression := &ast.BetweenExpression{
		Left:      left,
		Range:     [2]ast.Expression{nil, nil},
		Inclusive: [2]bool{false, false},
	}

	if p.currentToken().IsOneOfMany(lexer.LPAREN, lexer.LBRACKET) {
		if p.currentToken().Type == lexer.LPAREN {
			betweenExpression.Inclusive[0] = true
		}
		p.advance() // Consume the opening token
	} else {
		panic(fmt.Sprintf("expected ( or [ after between, got %s\n", lexer.TokenTypeString(p.currentToken().Type)))
	}

	lowerBound := parseExpression(p, defaultBindingPower)
	p.expect(lexer.COMMA)
	upperBound := parseExpression(p, defaultBindingPower)

	if p.currentToken().IsOneOfMany(lexer.RPAREN, lexer.RBRACKET) {
		if p.currentToken().Type == lexer.RPAREN {
			betweenExpression.Inclusive[1] = true
		}
		p.advance() // Consume the closing token
	} else {
		panic(fmt.Sprintf("expected ) or ] after expression, got %s\n", lexer.TokenTypeString(p.currentToken().Type)))
	}

	betweenExpression.Range = [2]ast.Expression{lowerBound, upperBound}

	return betweenExpression
}
