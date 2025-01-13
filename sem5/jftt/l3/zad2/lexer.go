package main

type TokenType int

const (
	EOF TokenType = iota
	NUM
	ADD
	SUB
	MUL
	DIV
	POW
	LPAREN
	RPAREN
	EOL
	INVALID
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for NUL, signifies EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = Token{Type: ADD, Value: "+"}
	case '-':
		tok = Token{Type: SUB, Value: "-"}
	case '*':
		tok = Token{Type: MUL, Value: "*"}
	case '/':
		tok = Token{Type: DIV, Value: "/"}
	case '^':
		tok = Token{Type: POW, Value: "^"}
	case '(':
		tok = Token{Type: LPAREN, Value: "("}
	case ')':
		tok = Token{Type: RPAREN, Value: ")"}
	case '\n':
		tok = Token{Type: EOL, Value: "\n"}
	case '#':
		l.skipComment()
		return l.NextToken()
	case '\\':
		if l.peekChar() == '\n' {
			l.readChar() // skip '\\'
			l.readChar() // skip '\n'
			return l.NextToken()
		} else {
			tok = Token{Type: INVALID, Value: string(l.ch)}
		}
	case 0:
		tok = Token{Type: EOF, Value: ""}
	default:
		if isDigit(l.ch) {
			number := l.readNumber()
			tok = Token{Type: NUM, Value: number}
			return tok
		} else {
			tok = Token{Type: INVALID, Value: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
