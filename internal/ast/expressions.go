package ast

import "github.com/dhruvsaxena1998/rel/internal/lexer"

type NumberExpression struct {
	Value float64
}

func (n *NumberExpression) expression() {}

type StringExpression struct {
	Value string
}

func (n *StringExpression) expression() {}

type SymbolExpression struct {
	Value string
}

func (n *SymbolExpression) expression() {}

type VariableExpression struct {
	Value string
}

func (n *VariableExpression) expression() {}

type BinaryExpression struct {
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (n *BinaryExpression) expression() {}

type PrefixExpression struct {
	Operator lexer.Token
	Right    Expression
}

func (n *PrefixExpression) expression() {}

type AssignmentExpression struct {
	Assigne  Expression
	Operator lexer.Token
	Value    Expression
}

func (n *AssignmentExpression) expression() {}

type ArrayLiteral struct {
	Elements []Expression
}

func (n *ArrayLiteral) expression() {}

type IfBranch struct {
	Condition  Expression
	Consequent Expression
}
type IfExpression struct {
	Branches []IfBranch
	Else     Expression
}

func (n *IfExpression) expression() {}
