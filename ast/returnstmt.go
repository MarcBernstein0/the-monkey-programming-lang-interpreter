package ast

import "github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"

// represents return statements: i.e return 4;
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression  // the value of the return statement
}

func (rs *ReturnStatement) statementNode() {

}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
