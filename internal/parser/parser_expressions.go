package parser

// This file contains the expression parsing logic for the parser

// ParseExpression parses a complete expression
func (p *Parser) ParseExpression() Expression {
	return p.parseOrExpression()
}

// parseOrExpression handles OR operations
func (p *Parser) parseOrExpression() Expression {
	left := p.parseAndExpression()

	for p.currentTokenIs(OR) {
		token := p.currentToken
		p.nextToken()
		right := p.parseAndExpression()
		left = &BinaryExpression{
			Token:    token,
			Left:     left,
			Operator: token.Literal,
			Right:    right,
		}
	}
	return left
}

// parseAndExpression handles AND operations
func (p *Parser) parseAndExpression() Expression {
	left := p.parseComparisonExpression()

	for p.currentTokenIs(AND) {
		token := p.currentToken
		p.nextToken()
		right := p.parseComparisonExpression()
		left = &BinaryExpression{
			Token:    token,
			Left:     left,
			Operator: token.Literal,
			Right:    right,
		}
	}
	return left
}
