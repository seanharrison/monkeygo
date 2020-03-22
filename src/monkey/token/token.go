package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// token types
const (
	ILLEGAL = "ILLEGAL" // illegal/unknown token
	EOF     = "EOF"     // end-of-file
	IDENT   = "IDENT"   // identifier
	INT     = "INT"     // integer

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)