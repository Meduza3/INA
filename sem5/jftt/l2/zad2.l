%option noyywrap

%{
%}

%%

\"[^\"]*\"  { ECHO; }
--.*        {}

%%

int main() {
    yylex();
}