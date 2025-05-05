package parser

// This file contains the literal parsing logic for the parser

// parseArrayLiteral handles array literals like [1, 2, 3]
func (p *Parser) parseArrayLiteral() Expression {
	// Create array literal node
	array := &ArrayLiteral{Token: p.currentToken}
	p.nextToken() // consume [

	array.Elements = []Expression{}

	// Handle empty array
	if p.currentTokenIs(RBRACKET) {
		p.nextToken() // consume ]
		return array
	}

	// Parse first element
	elem := p.parsePrimaryExpression()
	if elem == nil {
		p.addError("invalid array element")
		return nil
	}
	array.Elements = append(array.Elements, elem)

	// Parse remaining elements
	for p.currentTokenIs(COMMA) {
		p.nextToken() // consume comma

		// Check if we have a trailing comma followed by closing bracket
		if p.currentTokenIs(RBRACKET) {
			break
		}

		elem = p.parsePrimaryExpression()
		if elem == nil {
			p.addError("invalid array element after comma")
			return nil
		}
		array.Elements = append(array.Elements, elem)
	}

	if !p.currentTokenIs(RBRACKET) {
		p.addErrorf("expected right bracket, got %s", p.currentToken.Type)
		return nil
	}
	p.nextToken() // consume ]
	return array
}

// parseFunctionCall handles function calls like LOG(x)
func (p *Parser) parseFunctionCall() Expression {
	fc := &FunctionCall{
		Token:    p.currentToken,
		Function: p.currentToken.Literal,
	}
	p.nextToken() // move to '('
	p.nextToken() // move past '('

	fc.Arguments = []Expression{}
	for !p.currentTokenIs(RPAREN) {
		arg := p.ParseExpression()
		if arg != nil {
			fc.Arguments = append(fc.Arguments, arg)
		}

		if p.currentTokenIs(COMMA) {
			p.nextToken()
		}
	}

	if !p.currentTokenIs(RPAREN) {
		p.addError("expected right parenthesis")
		return nil
	}
	p.nextToken()
	return fc
}
