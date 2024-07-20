package lexer_test

import (
  "testing"
  "gopilator/token"
  "gopilator/lexer"
)

func TestNextToken(t *testing.T) {
  input := `PROGRAM IS 
              n, p 
            IN
              READ n;
              REPEAT
                p:=n/2;
                p:=2*p;
                IF n>p THEN
                  WRITE 1;
                ELSE
                  WRITE 0;
                ENDIF
                n:=n/2;
              UNTIL n=0;
            END`

  tests := []struct {
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.PROGRAM, "PROGRAM"},
    {token.IS, "IS"},
    {token.PIDENTIFIER, "n"},
    {token.COMMA, ","},
    {token.PIDENTIFIER, "p"},
    {token.IN, "IN"},
    {token.READ, "READ"},
    {token.PIDENTIFIER, "n"},
    {token.SEMICOLON, ";"},
    {token.REPEAT, "REPEAT"},
    {token.PIDENTIFIER, "p"},
    {token.ASSIGNMENT, ":="},
    {token.PIDENTIFIER, "n"},
    {token.DIVIDE, "/"},
    {token.NUM, "2"},
    {token.SEMICOLON, ";"},
    {token.PIDENTIFIER, "p"},
    {token.ASSIGNMENT, ":="},
    {token.NUM, "2"},
    {token.MULT, "*"},
    {token.PIDENTIFIER, "p"},
    {token.SEMICOLON, ";"},
    {token.IF, "IF"},
    {token.PIDENTIFIER, "n"},
    {token.GREATER, ">"},
    {token.PIDENTIFIER, "p"},
    {token.THEN, "THEN"},
    {token.WRITE, "WRITE"},
    {token.NUM, "1"},
    {token.SEMICOLON, ";"},
    {token.ELSE, "ELSE"},
    {token.WRITE, "WRITE"},
    {token.NUM, "0"},
    {token.SEMICOLON, ";"},
    {token.ENDIF, "ENDIF"},
    {token.PIDENTIFIER, "n"},
    {token.ASSIGNMENT, ":="},
    {token.PIDENTIFIER, "n"},
    {token.DIVIDE, "/"},
    {token.NUM, "2"},
    {token.SEMICOLON, ";"},
    {token.UNTIL, "UNTIL"},
    {token.PIDENTIFIER, "n"},
    {token.EQUALS, "="},
    {token.NUM, "0"},
    {token.SEMICOLON, ";"},
    {token.END, "END"},
    {token.EOF, ""},
  }

  l := lexer.New(input)

  for i, tt := range tests {
    tok := l.NextToken()
    if tok.Type != tt.expectedType {
      t.Fatalf("test[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
    }
    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("test[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
    }
  }
}
