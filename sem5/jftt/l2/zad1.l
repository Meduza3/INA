%{
  #include<stdio.h>
  #include<string.h>

  int yywrap();
  int yylex();
  int lines_count = 1;
  int words_count = 0;
%}

%%
(?# usuwanie pustych linii)
^[[:blank:]\\]*\n {};

(?# usuwanie pustego na początku i na końcu linii)
^[[:blank:]]+ {};
[[:blank:]]+$ {};

(?# zwijanie do jednej spacji)
[[:blank:]]+ {printf("");};

([^[:blank:]\n])+/[[:blank:]] {ECHO; words_count++;};
([^[:blank:]\n])+$ {ECHO; words_count++;};
\n {ECHO; lines_count++;};

%%
int yywrap() {
  return 1;
}
int main() {
  yylex();

  printf("\nlines count: $d\nwords count: %d\n", lines_count, words_count);
}