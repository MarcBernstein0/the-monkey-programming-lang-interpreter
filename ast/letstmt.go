package ast

import "github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"

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
