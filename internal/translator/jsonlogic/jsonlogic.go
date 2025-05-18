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

// translateExpression converts an AST expression node into its corresponding JSON Logic representation.
// It handles binary, prefix, between, variable, literal, and array expressions, panicking if the expression type is unsupported.
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

// translateBinaryExpression converts a binary AST expression into its JSON Logic representation based on the operator type.
// For comparison and equality operators, it returns a map with the operator as key and the translated operands as a two-element array.
// For logical "and" and "or" operators, it flattens nested conditions of the same type and returns a map with the operator as key and the combined conditions as an array.
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

// translatePrefixExpression converts a prefix logical expression (such as NOT or double NOT) into its JSON Logic representation.
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

// translateBetweenExpression converts a "between" AST expression into a JSON Logic "and" condition with appropriate comparison operators for the lower and upper bounds.
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

// combineConditions flattens nested logical conditions of the same operator into a single slice.
// It merges any sub-expressions that use the same logical operator, producing a flat list of conditions for use in JSON Logic structures.
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
