package translator_test

import (
	"testing"

	"github.com/dhruvsaxena1998/rel/internal/lexer"
	"github.com/dhruvsaxena1998/rel/internal/parser"
	translator "github.com/dhruvsaxena1998/rel/internal/translator/jsonlogic"
)

func BenchmarkTranslateToJSONLogic(b *testing.B) {

	inputs := []string{
		"1 == 1",
		"1 < 2 and 3 > 2",
		"1 == 0 or (5 < 10 and 5 > 0)",
		"@price < 100",
		"@fruit in ['apple', 'banana', 'orange']",
		"@temp < 30 and @humidity > 50 and (@weather == 'sunny' or @weather == 'cloudy')",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		input := inputs[i%len(inputs)]
		translator.TranslateToJSONLogic(
			parser.Parse(
				lexer.Tokenize(input),
			),
		)
	}
}
