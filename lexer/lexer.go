package lexer

import "github.com/avenue9977/monkey-interpreter/token"

// The Lexer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // We are ath the end of the file
		l.char = 0 // ASCII code for 'NUL' character
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	//nextToken := l.input[l.readPosition]
	//return token.Token{l.input[l.readPosition]}
	return token.Token{}
}
