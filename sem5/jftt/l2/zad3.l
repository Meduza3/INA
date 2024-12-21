%option noyywrap

%{
    #include <stdbool.h>

    static bool preserveDocumentation = false;

    static void handleOutput(const char* text, bool isDocComment);
%}

/* Token definitions for various comment types and string literals */
STRING_LITERAL         \"([^\"\\]|\\.)*\"
PREPROCESSOR          ^[ \t]*#.*$

/* Doxygen documentation patterns */
DOXYGEN_INLINE       \/\/[\/!][^\n]*\n
DOXYGEN_MULTILINE    \/\*[\*!]([^*]|\*[^/])*\*\/

/* Standard C++ comments */
REGULAR_INLINE       \/\/[^\/!][^\n]*\n
REGULAR_MULTILINE    \/\*[^*!]([^*]|\*[^/])*\*\/

%%
{STRING_LITERAL}      { ECHO; }                     /* Preserve string contents */
{PREPROCESSOR}        { ECHO; }                     /* Keep preprocessor directives */

{DOXYGEN_INLINE}      { handleOutput(yytext, true); }
{DOXYGEN_MULTILINE}   { handleOutput(yytext, true); }

{REGULAR_INLINE}      { handleOutput(yytext, false); }
{REGULAR_MULTILINE}   { handleOutput(yytext, false); }

.|\n                  { ECHO; }                     /* Pass through other characters */
`
%%

static void handleOutput(const char* text, bool isDocComment) {
    if (isDocComment && preserveDocumentation) {
        fprintf(yyout, "%s", text);
    } else {
        fprintf(yyout, " ");  /* Replace comment with single space */
    }
}

int main(int argc, char* argv[]) {
    /* Process command line arguments */
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--preserve-docs") == 0 || 
            strcmp(argv[i], "-d") == 0) {
            preserveDocumentation = true;
            break;
        }
    }
    
    yylex();
    return 0;
}