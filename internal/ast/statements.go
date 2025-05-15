package ast

type BlockStatement struct {
	Body Statement
}

func (b *BlockStatement) statement() {}

type LineStatement struct {
	Expressions []Expression
}

func (l *LineStatement) statement() {}

type ExpressionStatement struct {
	Expression Expression
}

func (e *ExpressionStatement) statement() {}

type VarDeclarationStatement struct {
	Identifier    string
	IsConstant    bool
	AssignedValue Expression
}

func (e *VarDeclarationStatement) statement() {}
