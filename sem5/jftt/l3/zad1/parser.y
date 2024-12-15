%{
#include "operations.hpp"
#include <iostream>
#include <string>

constexpr long long Characteristic = 1234577;

int yylex();
void yyerror(const std::string& message);

std::string rpnResult = "";
%}

%token NUM ADD SUB MUL DIV POW LPAREN RPAREN EOL INVALID

%left ADD SUB
%left MUL DIV
%nonassoc POW
%precedence NEG

%%

input:
    %empty
    | input line
    ;

line:
    expression EOL  { 
                        std::cout << rpnResult << '\n';
                        std::cout << "Wynik: " << $1 << '\n';

                        rpnResult = "";
                    }
    | error EOL     { 
                        std::cout << "Błąd.\n";
                        yyerrok;
                        rpnResult = "";
                    }
    ;

expression:
    number                                   { rpnResult += std::to_string($1) + " "; $$ = $1; }
    | expression ADD expression              { rpnResult += "+ "; $$ = Add($1, $3, Characteristic); }
    | expression SUB expression              { rpnResult += "- "; $$ = Sub($1, $3, Characteristic); }
    | expression MUL expression              { rpnResult += "* "; $$ = Mul($1, $3, Characteristic); }
    | expression DIV expression              { 
                                                 if ($3 == 0) {
                                                     yyerror;
                                                 } else {
                                                     rpnResult += "/ ";
                                                     $$ = Div($1, $3, Characteristic);
                                                 }
                                             }
    | expression POW exponent                { rpnResult += "^ "; $$ = Pow($1, $3, Characteristic); }
    | LPAREN expression RPAREN               { $$ = $2; }
    | SUB LPAREN expression RPAREN %prec NEG { rpnResult += "~ "; $$ = Sub(0, $3, Characteristic); }
    ;

exponent:
    exponent_number                        { rpnResult += std::to_string($1) + " "; $$ = $1; }
    | exponent ADD exponent                { rpnResult += "+ "; $$ = Add($1, $3, Characteristic - 1); }
    | exponent SUB exponent                { rpnResult += "- "; $$ = Sub($1, $3, Characteristic - 1); }
    | exponent MUL exponent                { rpnResult += "* "; $$ = Mul($1, $3, Characteristic - 1); }
    | exponent DIV exponent                { 
                                               if ($3 == 0) {
                                                   yyerror;
                                               } else {
                                                   rpnResult += "/ ";
                                                   $$ = Div($1, $3, Characteristic - 1);
                                               }
                                           }
    | LPAREN exponent RPAREN               { $$ = $2; }
    | SUB LPAREN exponent RPAREN %prec NEG { rpnResult += "~ "; $$ = Sub(0, $3, Characteristic - 1); }
    ;

number:
    NUM                    { $$ = $1 % Characteristic; }
    | SUB number %prec NEG { $$ = Sub(0, $2, Characteristic); }
    ;

exponent_number:
    NUM                             { $$ = $1 % (Characteristic - 1); }
    | SUB exponent_number %prec NEG { $$ = Sub(0, $2, Characteristic - 1); }
    ;

%%

void yyerror(const std::string& message) {
}

int main() {
    yyparse();
}
