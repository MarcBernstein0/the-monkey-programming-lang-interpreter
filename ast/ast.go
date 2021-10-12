package ast

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
