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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT  = "<"
	LEQ = "<="
	GT  = ">"
	GEQ = ">="

	EQ     = "=="
	NOT_EQ = "!="

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// check if ident is a keyword or not
func LookupIdent(ident string) TokenType {
	if toke, ok := keywords[ident]; ok {
		return toke
	}
	return IDENT
}
