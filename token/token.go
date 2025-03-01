package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special tokens
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + literals
	IDENT
	INT

	// Operators
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT
	GTE
	LTE
	EQ
	NOT_EQ

	// Delimiters
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

// TokenType to string mapping
var tokenStrings = map[TokenType]string{
	ILLEGAL:   "ILLEGAL",
	EOF:       "EOF",
	IDENT:     "IDENT",
	INT:       "INT",
	ASSIGN:    "=",
	PLUS:      "+",
	MINUS:     "-",
	BANG:      "!",
	ASTERISK:  "*",
	SLASH:     "/",
	LT:        "<",
	GT:        ">",
	GTE:       ">=",
	LTE:       "<=",
	EQ:        "==",
	NOT_EQ:    "!=",
	COMMA:     ",",
	SEMICOLON: ";",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",
	FUNCTION:  "FUNCTION",
	LET:       "LET",
	TRUE:      "TRUE",
	FALSE:     "FALSE",
	IF:        "IF",
	ELSE:      "ELSE",
	RETURN:    "RETURN",
}

// Convert TokenType to string
func (t TokenType) String() string {
	if str, ok := tokenStrings[t]; ok {
		return str
	}
	return "UNKNOWN"
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, isIdent := keywords[ident]; isIdent {
		return tok
	}
	return IDENT
}
