package lexer

import (
	"fmt"
	"testing"

	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"
)

// run basic lexor test
func TestNextTokenBasic(t *testing.T) {
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
		// fmt.Printf("Token: %+v\n", toke)

		// check if the token types match
		if toke.Type != tt.expectedType {
			t.Fatalf("Type test: test[%v] - tokentye wrong, expected = %q, got = %q\n",
				i, tt.expectedType, toke.Type)
		}

		// check if the token literals match
		if toke.Literal != tt.expectedLiteral {
			t.Fatalf("Literal test: tests[%v] - literal wrong. expected = %v, got = %v\n",
				i, tt.expectedLiteral, toke.Literal)
		}
	}

}

// run test on basic, valid monkey code
func TestNextTokenBasicCode(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x,y){
		x + y;
	};
	let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		toke := l.NextToken()
		fmt.Printf("Token: %+v\n", toke)

		// check if the token types match
		if toke.Type != tt.expectedType {
			t.Fatalf("Type test: test[%v] - tokentye wrong, expected = %q, got = %q\n",
				i, tt.expectedType, toke.Type)
		}

		// check if the token literals match
		if toke.Literal != tt.expectedLiteral {
			t.Fatalf("Literal test: tests[%v] - literal wrong. expected = %v, got = %v\n",
				i, tt.expectedLiteral, toke.Literal)
		}
	}

}

// run test on basic, valid monkey code
func TestNextTokenCodeWithOperators(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	let add = fn(x, y){
		x + y;
	};
	let result = add(five, ten);

	!-/*5;

	5 < 10 > 5;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		toke := l.NextToken()
		fmt.Printf("Token: %+v\n", toke)

		// check if the token types match
		if toke.Type != tt.expectedType {
			t.Fatalf("Type test: test[%v] - tokentye wrong, expected = %q, got = %q\n",
				i, tt.expectedType, toke.Type)
		}

		// check if the token literals match
		if toke.Literal != tt.expectedLiteral {
			t.Fatalf("Literal test: tests[%v] - literal wrong. expected = %v, got = %v\n",
				i, tt.expectedLiteral, toke.Literal)
		}
	}

}

func TestNextTokenCodeWithIfStatements(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	let add = fn(x, y){
		x + y;
	};
	let result = add(five, ten);

	!-/*5;

	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		toke := l.NextToken()
		fmt.Printf("Token: %+v\n", toke)

		// check if the token types match
		if toke.Type != tt.expectedType {
			t.Fatalf("Type test: test[%v] - tokentye wrong, expected = %q, got = %q\n",
				i, tt.expectedType, toke.Type)
		}

		// check if the token literals match
		if toke.Literal != tt.expectedLiteral {
			t.Fatalf("Literal test: tests[%v] - literal wrong. expected = %v, got = %v\n",
				i, tt.expectedLiteral, toke.Literal)
		}
	}

}

func TestNextTokenCodeWithTokensComposedofTwoChars(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	let add = fn(x, y){
		x + y;
	};
	let result = add(five, ten);

	!-/*5;

	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	10 >= 9;
	10 <= 10;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.GEQ, ">="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.LEQ, "<="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		toke := l.NextToken()
		fmt.Printf("Token: %+v\n", toke)

		// check if the token types match
		if toke.Type != tt.expectedType {
			t.Fatalf("Type test: test[%v] - tokentye wrong, expected = %q, got = %q\n",
				i, tt.expectedType, toke.Type)
		}

		// check if the token literals match
		if toke.Literal != tt.expectedLiteral {
			t.Fatalf("Literal test: tests[%v] - literal wrong. expected = %v, got = %v\n",
				i, tt.expectedLiteral, toke.Literal)
		}
	}

}
