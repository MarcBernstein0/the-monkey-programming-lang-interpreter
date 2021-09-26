package ast

import "github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"

// every part of the ast has to implement node interface
type Node interface {
	// takes a node and prints the token literal out
	// mainly for debugging purposes
	TokenLiteral() string
}

// statment nodes
type Statement interface {
	Node
	statementNode()
}

// expression nodes
type Expression interface {
	Node
	expresionNode()
}

// root of the ast
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len((p.Statements)) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// represents let statements i.e: let x = 5;
type LetStatment struct {
	Token token.Token // token.LET token
	Name  *Identifier // name of the let variable
	Value Expression  // expression after the assignment
}

func (lst *LetStatment) statementNode() {

}

func (ls *LetStatment) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expresionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
