package parser

type Parser struct {
	lexer        *Lexer
	currentToken Token
	peekToken    Token
	errors       []string
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	// Read two tokens to initialize currentToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}

// ParseExpression parses a complete expression
func (p *Parser) ParseExpression() Expression {
	return p.parseOrExpression()
}

// parseOrExpression handles OR operations
func (p *Parser) parseOrExpression() Expression {
	left := p.parseAndExpression()

	for p.currentToken.Type == OR {
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

	for p.currentToken.Type == AND {
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

// parseComparisonExpression handles ==, ===, !=, !==, >, <, >=, <=, IN operations
func (p *Parser) parseComparisonExpression() Expression {
	left := p.parsePrimaryExpression()

	// Check for NOT IN pattern
	if p.currentToken.Type == BANG && p.currentToken.Literal == "NOT" && p.peekToken.Type == IN {
		notToken := p.currentToken
		p.nextToken() // consume NOT
		inToken := p.currentToken
		p.nextToken() // consume IN

		// Ensure we have an array after IN
		if p.currentToken.Type != LBRACKET {
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

	// Handle IN operator specially
	if p.currentToken.Type == IN {
		token := p.currentToken
		p.nextToken() // consume IN

		// Ensure we have an array after IN
		if p.currentToken.Type != LBRACKET {
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

// parsePrimaryExpression handles basic expressions like variables, literals, and parenthesized expressions
func (p *Parser) parsePrimaryExpression() Expression {
	switch p.currentToken.Type {
	case LPAREN:
		p.nextToken()
		exp := p.ParseExpression()
		if p.currentToken.Type != RPAREN {
			p.addError("expected right parenthesis")
			return nil
		}
		p.nextToken()
		return exp

	case BANG:
		token := p.currentToken
		p.nextToken()
		right := p.parsePrimaryExpression()
		return &UnaryExpression{
			Token:    token,
			Operator: token.Literal,
			Right:    right,
		}

	case VARIABLE:
		variable := &Variable{Token: p.currentToken, Name: p.currentToken.Literal}
		p.nextToken()
		return variable

	case NUMBER:
		lit := &Literal{Token: p.currentToken, Value: p.currentToken.Literal}
		p.nextToken()
		return lit

	case STRING:
		// Remove quotes for the value but keep them in the token
		literal := p.currentToken.Literal
		value := literal
		if len(literal) >= 2 && (literal[0] == '"' || literal[0] == '\'') {
			value = literal[1 : len(literal)-1]
		}
		lit := &Literal{Token: p.currentToken, Value: value}
		p.nextToken()
		return lit

	case LBRACKET:
		return p.parseArrayLiteral()

	case IDENTIFIER:
		if p.peekToken.Type == LPAREN {
			return p.parseFunctionCall()
		}
		lit := &Literal{Token: p.currentToken, Value: p.currentToken.Literal}
		p.nextToken()
		return lit

	default:
		p.addError("unexpected token: " + string(p.currentToken.Type))
		return nil
	}
}

// parseArrayLiteral handles array literals like [1, 2, 3]
func (p *Parser) parseArrayLiteral() Expression {
	// Create array literal node
	array := &ArrayLiteral{Token: p.currentToken}
	p.nextToken() // consume [

	array.Elements = []Expression{}

	// Handle empty array
	if p.currentToken.Type == RBRACKET {
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
	for p.currentToken.Type == COMMA {
		p.nextToken() // consume comma

		// Check if we have a trailing comma followed by closing bracket
		if p.currentToken.Type == RBRACKET {
			break
		}

		elem = p.parsePrimaryExpression()
		if elem == nil {
			p.addError("invalid array element after comma")
			return nil
		}
		array.Elements = append(array.Elements, elem)
	}

	if p.currentToken.Type != RBRACKET {
		p.addError("expected right bracket, got " + string(p.currentToken.Type))
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
	for p.currentToken.Type != RPAREN {
		arg := p.ParseExpression()
		if arg != nil {
			fc.Arguments = append(fc.Arguments, arg)
		}

		if p.currentToken.Type == COMMA {
			p.nextToken()
		}
	}

	if p.currentToken.Type != RPAREN {
		p.addError("expected right parenthesis")
		return nil
	}
	p.nextToken()
	return fc
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
