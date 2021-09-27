package parser

import (
	"testing"

	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/ast"
	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/lexer"
	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil\n")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d\n", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != token.LET {
		t.Errorf("s.TokenLiteral not 'let' - got = %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatment)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got = %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got = %s", name, letStmt.Name.Value)
		return false
	}

	return true
}
