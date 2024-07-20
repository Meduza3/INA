package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	PROCEDURE = "PROCEDURE"
	IS        = "IS"
	IN        = "IN"
	END       = "END"

	PROGRAM = "PROGRAM"

	ASSIGNMENT = ":="
	SEMICOLON  = ";"
	IF         = "IF"
	THEN       = "THEN"
	ELSE       = "ELSE"
	ENDIF      = "ENDIF"

	WHILE    = "WHILE"
	DO       = "DO"
	ENDWHILE = "ENDWHILE"

	REPEAT = "REPEAT"
	UNTIL  = "UNTIL"

	READ  = "READ"
	WRITE = "WRITE"

	PIDENTIFIER = "PIDENTIFIER"
	NUM         = "NUM"

	LPARENT  = "["
	RPARENT  = "]"
	LBRACKET = "("
	RBRACKET = ")"
	COMMA    = ","

	T = "T"

	PLUS    = "+"
	MINUS   = "-"
	MULT    = "*"
	DIVIDE  = "/"
	PERCENT = "%"

	EQUALS    = "="
	NOTEQUALS = "!="
	GREATER   = ">"
	SMALLER   = "<"
	GEQ       = ">="
	LEQ       = "<="
)
