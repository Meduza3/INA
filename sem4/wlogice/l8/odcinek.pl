:- use_module(library(clpfd)).

odcinek(Lista) :-
    length(Lista, 16),
    Lista ins 0..1,
    sum(Lista, #=, 8),
    findall(Lista, ciag_jedynek(Lista), Solutions),
    member(Lista, Solutions).

ciag_jedynek(Lista) :-
    % Znalezienie początku bloku
    append(Prefix, Rest, Lista), % Split Lista w Prefix i Rest
    length(Prefix, Start),  % Start - Miejsce gdzie zaczynaja sie 1
    length(Ones, 8),
    Ones ins 1..1,
    append(Ones, Suffix, Rest), % Split Rest w jedynki i suffix
    Suffix ins 0..0.

% ?- odcinek(X), label(X), writeln(X), fail.
