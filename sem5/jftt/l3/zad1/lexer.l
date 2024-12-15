%option noyywrap

%{
    #include "parser.tab.hpp"
    #include <string.h>
%}

%%

^#(.|\\\n)*\n   { /* komentarze */ }
\\\n            { /* next line */ }
[[:blank:]]+    { /* spacje itd */ }
\n              { return EOL; }
[[:digit:]]+    { yylval = atoi(yytext); return NUM; }
"+"             { return ADD; }
"-"             { return SUB; }
"*"             { return MUL; }
"/"             { return DIV; }
"^"             { return POW; }
"("             { return LPAREN; }
")"             { return RPAREN; }
.               { return INVALID; }

%%