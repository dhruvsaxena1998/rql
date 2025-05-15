package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *Lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type Lexer struct {
	Tokens   []Token
	patterns []regexPattern
	source   string
	position int
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	for !lex.atEOF() {
		matched := false
		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remaining())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("unexpected character: %q\n %v ", lex.at(), lex.remaining()))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

func defaultHandler(tokenType TokenType, literal string) regexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
		lex.advanceN(len(literal))
		lex.push(NewToken(tokenType, literal))
	}
}

func numberHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remaining())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

func stringHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remaining())
	stringLiteral := lex.remaining()[match[0]+1 : match[1]-1]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

func symbolHandler(lex *Lexer, regex *regexp.Regexp) {
	value := regex.FindString(lex.remaining())

	if tokenType, exists := ReservedKeywords[value]; exists {
		lex.push(NewToken(tokenType, value))
	} else {
		lex.push(NewToken(IDENTIFIER, value))
	}

	lex.advanceN(len(value))
}
func variableHandler(lex *Lexer, regex *regexp.Regexp) {
	value := regex.FindString(lex.remaining())
	lex.push(NewToken(VARIABLE, value))
	lex.advanceN(len(value))
}

func skipHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remaining())
	lex.advanceN(match[1])
}

func (lex *Lexer) advanceN(n int) {
	lex.position += n
}

func (lex *Lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *Lexer) at() byte {
	return lex.source[lex.position]
}

func (lex *Lexer) remaining() string {
	return lex.source[lex.position:]
}

func (lex *Lexer) atEOF() bool {
	return lex.position >= len(lex.source)
}

func createLexer(source string) *Lexer {
	return &Lexer{
		source:   source,
		position: 0,
		Tokens:   make([]Token, 0),
		patterns: []regexPattern{
			// Skip whitespaces and tabs
			{regex: regexp.MustCompile(`[\s]+`), handler: skipHandler},
			// Skip comments
			{regex: regexp.MustCompile(`\/\/.* `), handler: skipHandler},

			{regex: regexp.MustCompile(`"[^"]*"`), handler: stringHandler},
			{regex: regexp.MustCompile(`'[^']*'`), handler: stringHandler},
			{regex: regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), handler: numberHandler},
			{regex: regexp.MustCompile(`@[a-zA-Z][a-zA-Z0-9_]*`), handler: variableHandler},
			{regex: regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_]*`), handler: symbolHandler},

			{regex: regexp.MustCompile(`\[`), handler: defaultHandler(LBRACKET, "[")},
			{regex: regexp.MustCompile(`\]`), handler: defaultHandler(RBRACKET, "]")},
			{regex: regexp.MustCompile(`\{`), handler: defaultHandler(LBRACE, "{")},
			{regex: regexp.MustCompile(`\}`), handler: defaultHandler(RBRACE, "}")},
			{regex: regexp.MustCompile(`\(`), handler: defaultHandler(LPAREN, "(")},
			{regex: regexp.MustCompile(`\)`), handler: defaultHandler(RPAREN, ")")},

			{regex: regexp.MustCompile(`===`), handler: defaultHandler(STRICT_EQ, "===")},
			{regex: regexp.MustCompile(`==`), handler: defaultHandler(EQ, "==")},
			{regex: regexp.MustCompile(`!==`), handler: defaultHandler(NOT_STRICT_EQ, "!==")},
			{regex: regexp.MustCompile(`!=`), handler: defaultHandler(NOT_EQ, "!=")},
			{regex: regexp.MustCompile(`=`), handler: defaultHandler(ASSIGN, "=")},
			{regex: regexp.MustCompile(`!`), handler: defaultHandler(NOT, "!")},
			{regex: regexp.MustCompile(`>=`), handler: defaultHandler(GTE, ">=")},
			{regex: regexp.MustCompile(`<=`), handler: defaultHandler(LTE, "<=")},
			{regex: regexp.MustCompile(`>`), handler: defaultHandler(GT, ">")},
			{regex: regexp.MustCompile(`<`), handler: defaultHandler(LT, "<")},

			{regex: regexp.MustCompile(`:`), handler: defaultHandler(COLON, ":")},
			{regex: regexp.MustCompile(`;`), handler: defaultHandler(SEMICOLON, ";")},
			{regex: regexp.MustCompile(`,`), handler: defaultHandler(COMMA, ",")},
		},
	}
}
