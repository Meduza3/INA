whitespace(' ').
whitespace('\t').
whitespace('\n').
whitespace('\r').

key(read).
key(write).
key(if).
key(then).
key(else).
key(fi).
key(while).
key(do).
key(od).
key(and).
key(or).
key(mod).

sep(';').
sep('+').
sep('-').
sep('*').
sep('/').
sep('(').
sep(')').
sep('<').
sep('>').
sep('=<').
sep('>=').
sep(':=').
sep('=').
sep('/=').

int(String) :-
  atom_number(String, Number),
  integer(Number),
  Number >= 0.

id(ID) :-
  upcase_atom(ID, ID).

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

tokenize(InputStream, Tokens) :- % Tokenize na start
    get_char(InputStream, Char),
    tokenize(Char, InputStream, Tokens). % Tokenize pierwszego chara

tokenize(end_of_file, _, []) :- !. % Jak skonczy sie plik, konczy sie prolog

tokenize(Char, InputStream, Tokens) :- % Pomin whitespace
    whitespace(Char),
    get_char(InputStream, NextChar),
    tokenize(NextChar, InputStream, Tokens).

tokenize(Char, InputStream, [TokensHead|TokensTail]) :-
    (
        sep(Char) ; Char = ':'
    ),
    get_char(InputStream, SecondChar),
    (
        SecondChar = end_of_file ->
        (
            TokensHead = Char,
            TokensTail = []
        );
        (
            atom_chars(CombinedChars, [Char, SecondChar]),
            (
                sep(CombinedChars) ->
                (
                    TokensHead = CombinedChars,
                    get_char(InputStream, ThirdChar),
                    tokenize(ThirdChar, InputStream, TokensTail)
                );
                (
                    TokensHead = Char,
                    tokenize(SecondChar, InputStream, TokensTail)
                )
            )
        )
    ).

tokenize(Char, InputStream, [Word|TokensTail]) :- % Cale slowo jest dodane do listy tokenow
    check_char_and_read_word(Char, NextChar, Chars, InputStream),
    atom_chars(Word, Chars),
    tokenize(NextChar, InputStream, TokensTail).


check_char_and_read_word(end_of_file, _, [], _) :- !.
check_char_and_read_word(LastChar, LastChar, [], _) :-
    (
        whitespace(LastChar) ; sep(LastChar) ; LastChar = ':'
    ),
    !.
check_char_and_read_word(Char, LastChar, [Char|Chars], InputStream) :-
    get_char(InputStream, NextChar),
    check_char_and_read_word(NextChar, LastChar, Chars, InputStream).
  
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

categorize_tokens([], []) :- !.

categorize_tokens([TokensHead|TokensTail], [key(TokensHead)|OutputTail]) :-
    key(TokensHead),
    categorize_tokens(TokensTail, OutputTail).

categorize_tokens([TokensHead|TokensTail], [int(TokensHead)|OutputTail]) :-
    int(TokensHead),
    categorize_tokens(TokensTail, OutputTail).

categorize_tokens([TokensHead|TokensTail], [sep(TokensHead)|OutputTail]) :-
    sep(TokensHead),
    categorize_tokens(TokensTail, OutputTail).


categorize_tokens([TokensHead|TokensTail], [id(TokensHead)|OutputTail]) :-
    id(TokensHead),
    categorize_tokens(TokensTail, OutputTail).

categorize_tokens([TokensHead|TokensTail], [unknown(TokensHead)|OutputTail]) :-
    categorize_tokens(TokensTail, OutputTail).

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

scanner(InputStream, Tokens) :-
    tokenize(InputStream, TokensWithoutCategories),
    categorize_tokens(TokensWithoutCategories, Tokens),
    !.