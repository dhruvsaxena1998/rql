package parser

import (
	"encoding/json"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected string // JSON string of expected JSONLogic output
	}{
		{
			input:    "@age > 18",
			expected: `{">": [{"var": "age"}, 18]}`,
		},
		{
			input:    "@name == 'John'",
			expected: `{"==": [{"var": "name"}, "John"]}`,
		},
		{
			input:    "@age > 18 AND @name == 'John'",
			expected: `{"and":[{">":[{"var": "age"},18]},{"==":[{"var": "name"},"John"]}]}`,
		},
		{
			input:    "@status IN ['active', 'pending']",
			expected: `{"in": [{"var": "status"}, ["active", "pending"]]}`,
		},
		{
			input:    "NOT @isDeleted",
			expected: `{"!": [{"var": "isDeleted"}]}`,
		},
		{
			input:    "@role NOT IN ['admin', 'moderator']",
			expected: `{"!":[{"in":[{"var":"role"},["admin","moderator"]]}]}`,
		},
		{
			input:    "@charge_code IN ['THZ', 'ABC'] AND @country = 'CHINA'",
			expected: `{"and": [{"in": [{"var": "charge_code"}, ["THZ", "ABC"]]}, {"==": [{"var": "country"}, "CHINA"]}]}`,
		},
		{
			input:    "(@age > 18) AND (@score >= 75)",
			expected: `{"and": [{">": [{"var": "age"}, 18]}, {">=": [{"var": "score"}, 75]}]}`,
		},
	}

	for i, tt := range tests {
		lexer := NewLexer(tt.input)
		parser := NewParser(lexer)

		expression := parser.ParseExpression()
		if expression == nil {
			t.Errorf("test[%d] - ParseExpression() returned nil. Errors: %v", i, parser.Errors())
			continue
		}

		jsonLogic, err := Transform(expression)
		if err != nil {
			t.Errorf("test[%d] - Transform() failed: %v", i, err)
			continue
		}

		result, err := json.Marshal(jsonLogic)
		if err != nil {
			t.Errorf("test[%d] - json.Marshal failed: %v", i, err)
			continue
		}

		// Compare JSON strings after normalizing
		var expectedJSON, resultJSON interface{}
		json.Unmarshal([]byte(tt.expected), &expectedJSON)
		json.Unmarshal(result, &resultJSON)

		expectedStr, _ := json.Marshal(expectedJSON)
		resultStr, _ := json.Marshal(resultJSON)

		if string(expectedStr) != string(resultStr) {
			t.Errorf("test[%d] - wrong result. got=%s, want=%s",
				i, string(resultStr), string(expectedStr))
		}
	}
}
