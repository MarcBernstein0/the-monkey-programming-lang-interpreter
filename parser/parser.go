package parser

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

// parse each statment
func (p *Parser) parseStatment() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatment()
	default:
		return nil
	}
}

// read next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectedPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) parseLetStatment() *ast.LetStatment {
	stmt := &ast.LetStatment{
		Token: p.curToken,
	}

	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	// TODO: pasing the expression of the let statement

	// encountering a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parse the program
func (p *Parser) ParseProgram() *ast.Program {
	program := new(ast.Program)
	program.Statements = make([]ast.Statement, 0)

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatment()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}
