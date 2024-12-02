%top{

package lexer

import (
    "io"
    "os"
    "log"
    "dijkstra/dimacs/token"
    "dijkstra/dimacs/parser"
)


// return type of yy.Lex()
type YYtype = token.Token


// type of Scanner.Context
type YYcontext interface{}


var yytext string

}

%%

[\t ]+          ;   // Ignore whitespace

\n              ;

c[^\n]*         ; // Ignore comments
                

p               {
                    return token.NewToken(token.PROBLEM, "p")
                }

a               {
                    return token.NewToken(token.ARC, "a")
                }

sp              {
                    return token.NewToken(token.PROBLEM_TYPE, "sp")
                }

[0-9]+          {
                    return token.NewToken(token.NUMBER, string(yytext))
                }

.|\n            ;   // Ignore unrecognized characters

%%

func yywrap() int {
    return 1
}

func (s *Scanner) Error(error string)  {
    return
}