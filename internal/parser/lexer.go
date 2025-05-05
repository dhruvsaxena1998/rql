package parser

import (
	"strings"
	"unicode"
)

type Lexer struct {
	input        string // the string being scanned
	position     int    // current position in input (points to current char)
	nextPosition int    // next position in input (after current char)
	ch           rune   // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

// readChar reads the next character and advances our positions in the input.
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.nextPosition])
	}
	l.position = l.nextPosition
	l.nextPosition++
}

// peekChar returns the next character without advancing the position.
func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0 // ASCII for NULL, signifies EOF
	}
	return l.input[l.nextPosition]
}

// skipWhitespaceAndComments advances past whitespace and '//' comments.
func (l *Lexer) skipWhitespaceAndComments() {
	for {
		if unicode.IsSpace(l.ch) {
			l.readChar()
		} else if l.ch == '/' && l.peekChar() == '/' {
			// single-line comment
			for l.ch != '\n' && l.ch != 0 {
				l.readChar()
			}
		} else if l.ch == '/' && l.peekChar() == '*' {
			// block comment
			l.readChar()
			l.readChar()
			for !(l.ch == '*' && l.peekChar() == '/') && l.ch != 0 {
				l.readChar()
			}
			// consume '*/'
			l.readChar()
			l.readChar()
		} else {
			break
		}
	}
}

// readIdentifier reads in an identifier (or keyword).
func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	raw := l.input[start:l.position]
	// normalize for case-insensitive keywords
	return raw
}

// readVariable reads in a variable prefixed by '@'.
func (l *Lexer) readVariable() string {
	start := l.position
	l.readChar() // consume '@'
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.position]
}

// readNumber reads in a contiguous sequence of digits.
func (l *Lexer) readNumber() string {
	start := l.position
	dotCount := 0
	for isDigit(l.ch) || (l.ch == '.' && dotCount == 0) {
		if l.ch == '.' {
			dotCount++
		}
		l.readChar()
	}
	return l.input[start:l.position]
}

// readString reads a quoted string literal, including delimiters.
func (l *Lexer) readString() string {
	l.readChar() // skip opening '"'

	var result strings.Builder
	for l.ch != '"' && l.ch != 0 {
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case '"':
				result.WriteRune('"')
			case 'n':
				result.WriteRune('\n')
			case 't':
				result.WriteRune('\t')
			case '\\':
				result.WriteRune('\\')
			default:
				result.WriteRune(l.ch)
			}
		} else {
			result.WriteRune(l.ch)
		}
		l.readChar()
	}
	// skip closing '"'
	l.readChar()
	return result.String()
}

func (l *Lexer) nextToken() Token {
	var token Token

	l.skipWhitespaceAndComments()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			if l.peekChar() == '=' {
				// '===' strict equality
				l.readChar()
				token = NewToken(STRICT_EQ, "===")
			} else {
				token = NewToken(EQ, "==")
			}
		} else {
			token = NewToken(ASSIGN, "=")
		}

	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			if l.peekChar() == '=' {
				// '!==' strict not equal
				l.readChar()
				token = NewToken(STRICT_NOT, "!==")
			} else {
				token = NewToken(NOT_EQ, "!=")
			}
		} else {
			token = NewToken(BANG, "!")
		}

	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			token = NewToken(GTE, ">=")
		} else {
			token = NewToken(GT, ">")
		}

	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			token = NewToken(LTE, "<=")
		} else {
			token = NewToken(LT, "<")
		}

	case ';':
		token = NewToken(SEMICOLON, ";")

	case ',':
		token = NewToken(COMMA, ",")

	case '(':
		token = NewToken(LPAREN, "(")

	case ')':
		token = NewToken(RPAREN, ")")

	case '[':
		token = NewToken(LBRACKET, "[")

	case ']':
		token = NewToken(RBRACKET, "]")

	case '{':
		token = NewToken(LBRACE, "{")

	case '}':
		token = NewToken(RBRACE, "}")

	case '"', '\'':
		literal := l.readString()
		token = Token{Type: STRING, Literal: literal}

	case '@':
		literal := l.readVariable()
		token = Token{Type: VARIABLE, Literal: literal}

	case 0:
		token = Token{Type: EOF, Literal: ""}

	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokenType := LookupIdentifier(literal)
			token = NewToken(tokenType, literal)
			return token
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			token = NewToken(NUMBER, literal)
			return token
		} else {
			token = NewToken(ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return token
}
