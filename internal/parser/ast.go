package parser

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

// Represents a variable reference like @age
type Variable struct {
	Token Token // The '@' token
	Name  string
}

func (v *Variable) expressionNode()      {}
func (v *Variable) TokenLiteral() string { return v.Token.Literal }

// Represents literals like numbers and strings
type Literal struct {
	Token Token
	Value interface{}
}

func (l *Literal) expressionNode()      {}
func (l *Literal) TokenLiteral() string { return l.Token.Literal }

// Represents binary operations like AND, OR, ==, >, <
type BinaryExpression struct {
	Token    Token // The operator token, e.g. AND
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Token.Literal }

// Represents unary operations like NOT
type UnaryExpression struct {
	Token    Token // The operator token, e.g. NOT
	Operator string
	Right    Expression
}

func (ue *UnaryExpression) expressionNode()      {}
func (ue *UnaryExpression) TokenLiteral() string { return ue.Token.Literal }

// Represents array/list literals for IN operator
type ArrayLiteral struct {
	Token    Token // The '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

// Represents function calls like LOG()
type FunctionCall struct {
	Token     Token // The function name token
	Function  string
	Arguments []Expression
}

func (fc *FunctionCall) expressionNode()      {}
func (fc *FunctionCall) TokenLiteral() string { return fc.Token.Literal }
