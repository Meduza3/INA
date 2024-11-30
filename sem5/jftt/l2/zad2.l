%option noyywrap

%{
%}

%%

\"[^\"]*\"  /* Zachowuje literały napisów */ { ECHO; }
--.*        /* Usuwa komentarze jednoliniowe */ {}
.               /* Zachowuje wszystkie inne znaki */ { ECHO; }

%%

int main() {
    yylex();
    return 0;
}