package lexer

import (
	"okapi/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     -1,
		nextPosition: 0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position++
	l.nextPosition++
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func HandleNextCharToken(l *Lexer, currentToken token.TokenType, nextChar byte, nextToken token.TokenType) token.Token {
	if l.peekChar() == nextChar {
		firstCh := l.ch
		l.readChar()
		secondCh := l.ch
		return token.Token{Type: nextToken, Literal: string(firstCh) + string(secondCh)}
	} else {
		return newToken(currentToken, l.ch)
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '=':
		tok = HandleNextCharToken(l, token.ASSIGN, '=', token.EQ)
	case '!':
		tok = HandleNextCharToken(l, token.BANG, '=', token.NOT_EQ)
	case '<':
		tok = HandleNextCharToken(l, token.LT, '=', token.LTE)
	case '>':
		tok = HandleNextCharToken(l, token.GT, '=', token.GTE)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
