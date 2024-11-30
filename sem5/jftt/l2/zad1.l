%{
#include <stdio.h>
#include <string.h>

int yywrap();
int yylex();

// Liczniki statystyk tekstu
int wiersze_total = 1;    // Zaczynamy od 1, bo ostatnia linia
int slowa_total = 0;      // Licznik slow w tekscie
%}

%%
^[\t\n ]*\n         /* Usuwa puste linie */ {}
^[[:blank:]]+       /* Usuwa biale znaki na poczatku linii */ {}
[[:blank:]]+$       /* Usuwa biale znaki na koncu linii */ {}
[[:blank:]]+        /* Zamienia ciagi bialych znakow na pojedyncza spacje */ { printf(" "); }

([^[:blank:]\n])+/[[:blank:]]    { /* Obsługa słów z następującą spacją */
    ECHO;
    slowa_total++;
}

([^[:blank:]\n])+$  { /* Obsługa słów na końcu linii */
    ECHO;
    slowa_total++;
}

\n      { /* Obsługa nowej linii */
    ECHO;
    wiersze_total++;
}
%%

int yywrap() {
    return 1;
}

int main() {
    yylex();
    
    // Wyswietlenie statystyk
    printf("\nLiczba wierszy: %d\nLiczba slow: %d\n",
           wiersze_total, slowa_total);
    return 0;
}