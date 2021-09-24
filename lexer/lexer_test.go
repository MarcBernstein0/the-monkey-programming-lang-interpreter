package lexer

import (
	"testing"

	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		toke := l.NextToken()

		// check if the token types match
		if toke.Type != tt.expectedType {
			t.Fatalf("Type test: test[%v] - tokentye wrong, expected=%q, got=%q\n",
				i, tt.expectedType, toke.Type)
		}

		// check if the token literals match
		if toke.Literal != tt.expectedLiteral {
			t.Fatalf("Literal test: tests[%v] - literal wrong. expected=%v, got=%v\n",
				i, tt.expectedLiteral, toke.Literal)
		}
	}

}
