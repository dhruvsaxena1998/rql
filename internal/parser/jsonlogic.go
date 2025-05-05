package parser

import (
	"fmt"
	"strconv"
)

// JSONLogic represents a JSON Logic compatible structure
type JSONLogic map[string]interface{}

// Transform converts an AST Expression into a JSONLogic structure
func Transform(node Expression) (interface{}, error) {
	if node == nil {
		return nil, fmt.Errorf("cannot transform nil node")
	}

	switch n := node.(type) {
	case *BinaryExpression:
		return transformBinaryExpression(n)
	case *UnaryExpression:
		return transformUnaryExpression(n)
	case *Variable:
		return transformVariable(n)
	case *Literal:
		return transformLiteral(n)
	case *ArrayLiteral:
		return transformArrayLiteral(n)
	case *FunctionCall:
		return transformFunctionCall(n)
	default:
		return nil, fmt.Errorf("unsupported node type: %T", n)
	}
}

// transformBinaryExpression handles binary operations (AND, OR, comparisons)
func transformBinaryExpression(be *BinaryExpression) (JSONLogic, error) {
	left, err := Transform(be.Left)
	if err != nil {
		return nil, err
	}

	right, err := Transform(be.Right)
	if err != nil {
		return nil, err
	}

	// Map operators to JSONLogic format
	switch be.Operator {
	case "AND":
		return JSONLogic{"and": []interface{}{left, right}}, nil
	case "OR":
		return JSONLogic{"or": []interface{}{left, right}}, nil
	case "=", "==", "===":
		return JSONLogic{"==": []interface{}{left, right}}, nil
	case "!=", "!==":
		return JSONLogic{"!=": []interface{}{left, right}}, nil
	case ">":
		return JSONLogic{">": []interface{}{left, right}}, nil
	case "<":
		return JSONLogic{"<": []interface{}{left, right}}, nil
	case ">=":
		return JSONLogic{">=": []interface{}{left, right}}, nil
	case "<=":
		return JSONLogic{"<=": []interface{}{left, right}}, nil
	case "IN":
		return JSONLogic{"in": []interface{}{left, right}}, nil
	default:
		return nil, fmt.Errorf("unsupported binary operator: %s", be.Operator)
	}
}

// transformUnaryExpression handles unary operations (NOT, !)
func transformUnaryExpression(ue *UnaryExpression) (JSONLogic, error) {
	right, err := Transform(ue.Right)
	if err != nil {
		return nil, err
	}

	// Both '!' and 'NOT' are mapped to the same JSONLogic operator
	// Wrap the operand in an array as per JSONLogic format
	return JSONLogic{"!": []interface{}{right}}, nil
}

// transformVariable handles variable references
func transformVariable(v *Variable) (JSONLogic, error) {
	// Remove the @ prefix for JSONLogic
	name := v.Name[1:]
	return JSONLogic{"var": name}, nil
}

// transformLiteral handles literal values (numbers, strings)
func transformLiteral(l *Literal) (interface{}, error) {
	// For numbers, convert string to float64
	if l.Token.Type == NUMBER {
		if num, err := strconv.ParseFloat(l.Value.(string), 64); err == nil {
			return num, nil
		}
	}

	// For strings and other literals, use as-is
	return l.Value, nil
}

// transformArrayLiteral handles array literals
func transformArrayLiteral(al *ArrayLiteral) (interface{}, error) {
	elements := make([]interface{}, len(al.Elements))
	for i, elem := range al.Elements {
		transformed, err := Transform(elem)
		if err != nil {
			return nil, err
		}
		elements[i] = transformed
	}

	return elements, nil
}

// transformFunctionCall handles function calls (LOG)
func transformFunctionCall(fc *FunctionCall) (JSONLogic, error) {
	args := make([]interface{}, len(fc.Arguments))
	for i, arg := range fc.Arguments {
		transformed, err := Transform(arg)
		if err != nil {
			return nil, err
		}
		args[i] = transformed
	}

	switch fc.Function {
	case "LOG":
		return JSONLogic{"log": args}, nil
	default:
		return nil, fmt.Errorf("unsupported function: %s", fc.Function)
	}
}
