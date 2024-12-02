package token

type TokenType int

const (
	COMMENT TokenType = iota + 57346
	PROBLEM
	ARC
	NUMBER
	PROBLEM_TYPE
	EOL
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
