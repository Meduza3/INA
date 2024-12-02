// lexer_wrapper.go
package lexer

import (
	"dijkstra/dimacs/parser" // Adjust the import path as necessary
	"dijkstra/dimacs/token"
	"fmt"
)

// LexerWrapper adapts the flexgo-generated lexer to the parser's expected YyLexer interface
type LexerWrapper struct {
	lexer *Scanner
}

// NewLexerWrapper initializes the LexerWrapper with the given lexer
func NewLexerWrapper(lexer *Scanner) *LexerWrapper {
	return &LexerWrapper{
		lexer: lexer,
	}
}

// Lex implements the parser.YyLexer interface's Lex method
func (lw *LexerWrapper) Lex(lval *parser.YySymType) int {
	tok := lw.lexer.Lex() // Calls the flexgo-generated Lex() method, returning token.Token

	switch tok.Type {
	case token.PROBLEM:
		return int(token.PROBLEM)
	case token.ARC:
		return int(token.ARC)
	case token.PROBLEM_TYPE:
		return int(token.PROBLEM_TYPE)
	case token.NUMBER:
		return int(token.NUMBER)
	case token.EOL:
		return int(token.EOL)
	default:
		return 0 // Undefined token
	}
}

// Error implements the parser.YyLexer interface's Error method
func (lw *LexerWrapper) Error(s string) {
	fmt.Printf("Parse error: %s\n", s)
}
