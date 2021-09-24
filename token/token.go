package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const ( // might be able to do this with iotas

	// errors
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // integer

	// opterators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
