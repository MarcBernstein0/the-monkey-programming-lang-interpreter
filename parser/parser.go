package paser

import (
	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/ast"
	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/lexer"
	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// create a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	// Read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// read next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// parse the program
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
