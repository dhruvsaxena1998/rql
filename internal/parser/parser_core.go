package parser

// This file contains the core parser functionality

// parseComparisonExpression handles ==, ===, !=, !==, >, <, >=, <=, IN operations
func (p *Parser) parseComparisonExpression() Expression {
	left := p.parsePrimaryExpression()

	// Check for NOT IN pattern
	if p.isNotInPattern() {
		return p.parseNotInExpression(left)
	}

	// Handle IN operator specially
	if p.currentTokenIs(IN) {
		return p.parseInExpression(left)
	}

	// Handle other comparison operators
	for p.isComparisonOperator(p.currentToken.Type) {
		token := p.currentToken
		p.nextToken()
		right := p.parsePrimaryExpression()
		left = &BinaryExpression{
			Token:    token,
			Left:     left,
			Operator: token.Literal,
			Right:    right,
		}
	}
	return left
}

// isNotInPattern checks if the current tokens form a NOT IN pattern
func (p *Parser) isNotInPattern() bool {
	return p.currentToken.Type == BANG &&
		p.currentToken.Literal == "NOT" &&
		p.peekToken.Type == IN
}

// parseNotInExpression handles the NOT IN special case
func (p *Parser) parseNotInExpression(left Expression) Expression {
	notToken := p.currentToken
	p.nextToken() // consume NOT
	inToken := p.currentToken
	p.nextToken() // consume IN

	// Ensure we have an array after IN
	if !p.currentTokenIs(LBRACKET) {
		p.addError("expected array after NOT IN operator")
		return nil
	}

	// Parse the array literal
	right := p.parseArrayLiteral()
	if right == nil {
		p.addError("failed to parse array after NOT IN operator")
		return nil
	}

	// Create IN expression
	inExpr := &BinaryExpression{
		Token:    inToken,
		Left:     left,
		Operator: inToken.Literal,
		Right:    right,
	}

	// Wrap with NOT
	return &UnaryExpression{
		Token:    notToken,
		Operator: notToken.Literal,
		Right:    inExpr,
	}
}

// parseInExpression handles the IN operator
func (p *Parser) parseInExpression(left Expression) Expression {
	token := p.currentToken
	p.nextToken() // consume IN

	// Ensure we have an array after IN
	if !p.currentTokenIs(LBRACKET) {
		p.addError("expected array after IN operator")
		return nil
	}

	// Parse the array literal
	right := p.parseArrayLiteral()
	if right == nil {
		p.addError("failed to parse array after IN operator")
		return nil
	}

	return &BinaryExpression{
		Token:    token,
		Left:     left,
		Operator: token.Literal,
		Right:    right,
	}
}

// parsePrimaryExpression handles basic expressions like variables, literals, and parenthesized expressions
func (p *Parser) parsePrimaryExpression() Expression {
	switch p.currentToken.Type {
	case LPAREN:
		return p.parseParenthesizedExpression()

	case BANG:
		return p.parseUnaryExpression()

	case VARIABLE:
		return p.parseVariable()

	case NUMBER:
		return p.parseNumberLiteral()

	case STRING:
		return p.parseStringLiteral()

	case LBRACKET:
		return p.parseArrayLiteral()

	case IDENTIFIER:
		if p.peekToken.Type == LPAREN {
			return p.parseFunctionCall()
		}
		return p.parseIdentifier()

	default:
		p.addErrorf("unexpected token: %s", p.currentToken.Type)
		return nil
	}
}

// parseParenthesizedExpression handles expressions in parentheses
func (p *Parser) parseParenthesizedExpression() Expression {
	p.nextToken() // consume '('
	exp := p.ParseExpression()

	if !p.currentTokenIs(RPAREN) {
		p.addError("expected right parenthesis")
		return nil
	}
	p.nextToken() // consume ')'
	return exp
}

// parseUnaryExpression handles unary operations like NOT
func (p *Parser) parseUnaryExpression() Expression {
	token := p.currentToken
	p.nextToken() // consume operator
	right := p.parsePrimaryExpression()

	return &UnaryExpression{
		Token:    token,
		Operator: token.Literal,
		Right:    right,
	}
}

// parseVariable handles variable references like @age
func (p *Parser) parseVariable() Expression {
	variable := &Variable{Token: p.currentToken, Name: p.currentToken.Literal}
	p.nextToken()
	return variable
}

// parseNumberLiteral handles numeric literals
func (p *Parser) parseNumberLiteral() Expression {
	lit := &Literal{Token: p.currentToken, Value: p.currentToken.Literal}
	p.nextToken()
	return lit
}

// parseStringLiteral handles string literals
func (p *Parser) parseStringLiteral() Expression {
	// Remove quotes for the value but keep them in the token
	literal := p.currentToken.Literal
	value := literal
	if len(literal) >= 2 && (literal[0] == '"' || literal[0] == '\'') {
		value = literal[1 : len(literal)-1]
	}
	lit := &Literal{Token: p.currentToken, Value: value}
	p.nextToken()
	return lit
}

// parseIdentifier handles identifiers
func (p *Parser) parseIdentifier() Expression {
	lit := &Literal{Token: p.currentToken, Value: p.currentToken.Literal}
	p.nextToken()
	return lit
}
