// currently only supports ascii, add unicode support in future
package lexer

import (
	"fmt"

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

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var toke token.Token

	// switch case based on what the character is
	switch l.ch {
	case '=':
		toke = newToken(token.ASSIGN, l.ch)
	case '+':
		toke = newToken(token.PLUS, l.ch)
	case '(':
		toke = newToken(token.LPAREN, l.ch)
	case ')':
		toke = newToken(token.RPAREN, l.ch)
	case '{':
		toke = newToken(token.LBRACE, l.ch)
	case '}':
		toke = newToken(token.RBRACE, l.ch)
	case ',':
		toke = newToken(token.COMMA, l.ch)
	case ';':
		toke = newToken(token.SEMICOLON, l.ch)
	case 0:
		toke = newToken(token.EOF, 0) // zero is empty character as '' is an illegal rune literal
	}
	fmt.Printf("Token: %+v\n", toke)

	l.readChar()
	return toke
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}
