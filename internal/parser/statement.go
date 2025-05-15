package parser

import (
	"fmt"

	"github.com/dhruvsaxena1998/rel/internal/ast"
	"github.com/dhruvsaxena1998/rel/internal/lexer"
)

func parseStatement(p *parser) ast.Statement {
	stmtFn, exists := statementLU[p.currentToken().Type]

	if exists {
		return stmtFn(p)
	}

	expression := parseExpression(p, defaultBindingPower)

	return &ast.ExpressionStatement{
		Expression: expression,
	}
}

func parseVarDeclarationStatement(p *parser) ast.Statement {
	isConstant := p.advance().Type == lexer.CONST
	variableName := p.expectError(lexer.IDENTIFIER, fmt.Errorf("expected identifier after %v", p.currentToken().Type)).Literal

	p.expect(lexer.ASSIGN)
	assignedValue := parseExpression(p, assignment)

	return &ast.VarDeclarationStatement{
		IsConstant:    isConstant,
		Identifier:    variableName,
		AssignedValue: assignedValue,
	}
}
