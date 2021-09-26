// currently only supports ascii, add unicode support in future
package lexer

import (
	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"
)

// Currently reads entire string and stores in struct, change to make it read a file instead and store position in file
// as to not store monkey code in memory
type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	readPos int  // current reading position in input (after current char)
	ch      byte // current char being looked at
}

// create new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

// read next character in program
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

// read entire string until non valid character
// function being passed in checks if given character is valid in monkey
func (l *Lexer) readCharacter(fn func(byte) bool) string {
	position := l.pos
	for fn(l.ch) {
		l.readChar()
	}

	return l.input[position:l.pos]
}

// skip whitespace
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}

// generates the next token in the lexer
func (l *Lexer) NextToken() token.Token {
	var tokeType token.TokenType
	var tokeLiteral string

	l.skipWhitespace()

	// switch case based on what the character is
	switch l.ch {
	case '=':
		tokeType, tokeLiteral = token.ASSIGN, string(l.ch)
	case '+':
		tokeType, tokeLiteral = token.PLUS, string(l.ch)
	case '-':
		tokeType, tokeLiteral = token.MINUS, string(l.ch)
	case '!':
		tokeType, tokeLiteral = token.BANG, string(l.ch)
	case '*':
		tokeType, tokeLiteral = token.ASTERISK, string(l.ch)
	case '/':
		tokeType, tokeLiteral = token.SLASH, string(l.ch)
	case '(':
		tokeType, tokeLiteral = token.LPAREN, string(l.ch)
	case ')':
		tokeType, tokeLiteral = token.RPAREN, string(l.ch)
	case '{':
		tokeType, tokeLiteral = token.LBRACE, string(l.ch)
	case '}':
		tokeType, tokeLiteral = token.RBRACE, string(l.ch)
	case ',':
		tokeType, tokeLiteral = token.COMMA, string(l.ch)
	case ';':
		tokeType, tokeLiteral = token.SEMICOLON, string(l.ch)
	case '<':
		tokeType, tokeLiteral = token.LT, string(l.ch)
	case '>':
		tokeType, tokeLiteral = token.GT, string(l.ch)
	case 0:
		tokeType, tokeLiteral = token.EOF, "" // zero is empty character as '' is an illegal rune literal
	default:
		// return in default as to not stop at the space then skip the character by accident
		if isLetter(l.ch) {
			tokeLiteral = l.readCharacter(isLetter)
			tokeType = token.LookupIdent(tokeLiteral)

		} else if isDigit(l.ch) {
			tokeType = token.INT
			tokeLiteral = l.readCharacter(isDigit)
		} else {
			tokeType, tokeLiteral = token.ILLEGAL, string(l.ch)

		}
		toke := newToken(tokeType, tokeLiteral)

		return toke

	}

	toke := newToken(tokeType, tokeLiteral)

	l.readChar()
	return toke
}

// generates new token struct
func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
	}
}

// checks if character is valid for identifiers and keywords
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'

}

// checks if character is valid integer
// TODO: add support for floats, hex notation and/or binary
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
