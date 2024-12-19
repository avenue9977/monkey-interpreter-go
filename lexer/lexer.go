// Package lexer implements Lexer struct used for reading
// "Monkey" source code and outputs the tokens that represent it.
// The Lexer supports the full Unicode range
package lexer

import (
	"github.com/avenue9977/monkey-interpreter/token"
	"unicode"
)

// A Lexer stores the input source code
// and keep track of the current position
// the next position and the currently read char
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         rune // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // We are ath the end of the file
		l.char = 0 // ASCII code for 'NUL' character
	} else {
		l.char = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.char)}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '!':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.char)}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupKeyword(tok.Literal)
			return tok
		} else if unicode.IsDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, char rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return rune(l.input[l.readPosition])
	}
}
