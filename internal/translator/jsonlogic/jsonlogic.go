package translator

import (
	"fmt"

	"github.com/dhruvsaxena1998/rel/internal/ast"
	"github.com/dhruvsaxena1998/rel/internal/lexer"
)

func TranslateToJSONLogic(blockStatement ast.BlockStatement) map[string]any {
	body := blockStatement.Body.(*ast.LineStatement)

	if len(body.Expressions) == 0 {
		return map[string]any{}
	}

	if len(body.Expressions) == 1 {
		return translateExpression(body.Expressions[0]).(map[string]any)
	}

	return map[string]any{}
}

func translateExpression(expr ast.Expression) any {
	switch e := expr.(type) {
	case *ast.BinaryExpression:
		return translateBinaryExpression(e.Left, e.Operator, e.Right)

	case *ast.PrefixExpression:
		return translatePrefixExpression(e.Operator, e.Right)

	case *ast.BetweenExpression:
		return translateBetweenExpression(e.Left, e.Inclusive, e.Range)

	case *ast.VariableExpression:
		return map[string]any{"var": e.Value[1:]}

	case *ast.NumberExpression:
		return e.Value

	case *ast.StringExpression:
		return e.Value

	case *ast.SymbolExpression:
		return e.Value

	case *ast.BooleanExpression:
		return e.Value

	case *ast.ArrayLiteral:
		elements := make([]any, len(e.Elements))
		for i, element := range e.Elements {
			elements[i] = translateExpression(element)
		}
		return elements
	default:
		panic(fmt.Sprintf("cannot translate expression: %T", expr))
	}
}

func translateBinaryExpression(left ast.Expression, operator lexer.Token, right ast.Expression) any {

	result := make(map[string]interface{})

	switch operator.Type {
	case lexer.EQ, lexer.NOT_EQ, lexer.GT, lexer.GTE, lexer.LT, lexer.LTE, lexer.STRICT_EQ, lexer.NOT_STRICT_EQ, lexer.IN:
		result[operator.Literal] = []any{translateExpression(left), translateExpression(right)}

	case lexer.AND, lexer.OR:
		leftExpr := translateExpression(left)
		rightExpr := translateExpression(right)

		conditions := combineConditions(operator.Literal, leftExpr, rightExpr)
		result[operator.Literal] = conditions
	}

	return result
}

func translatePrefixExpression(operator lexer.Token, right ast.Expression) any {
	result := make(map[string]any)

	switch operator.Type {
	case lexer.NOT:
		result["!"] = translateExpression(right)
	case lexer.NOT_NOT:
		result["!!"] = translateExpression(right)
	}

	return result
}

func translateBetweenExpression(left ast.Expression, inclusive [2]bool, right [2]ast.Expression) any {
	result := make(map[string]any)

	lowerOperator := ">"
	if inclusive[0] {
		lowerOperator = ">="
	}

	upperOperator := "<"
	if inclusive[1] {
		upperOperator = "<="
	}

	result["and"] = []any{
		map[string]any{lowerOperator: []any{translateExpression(left), translateExpression(right[0])}},
		map[string]any{upperOperator: []any{translateExpression(left), translateExpression(right[1])}},
	}

	return result
}

func combineConditions(operator string, expressions ...any) []any {
	conditions := make([]any, 0)

	for _, expr := range expressions {
		// check if expr is a logical operation of the same type
		if logicalExpr, ok := expr.(map[string]any); ok {
			if logicalConditions, exists := logicalExpr[operator]; exists {
				if conditionArray, isArray := logicalConditions.([]any); isArray {
					conditions = append(conditions, conditionArray...)
				} else {
					conditions = append(conditions, logicalConditions)
				}
				continue
			}
		}

		conditions = append(conditions, expr)
	}

	return conditions
}
