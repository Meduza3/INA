% Zadanie 1

srodkowy(X, L) :-
    L = [_|_],
    length(L, Len),
    divmod(Len, 2, Q, R),
    R =:= 1,
    N is Q + R,
    nth1(N, L, X).

% Zadanie 2

jednokrotnie(X, L) :-
    select(X, L, L2),
    \+ member(X, L2).

dwukrotnie(X, L) :-
    select(X, L, L2),
    select(X, L2, L3),
    \+ member(X, L3).


% Zadanie 3

arc(a, b).
arc(b, a).
arc(b, c).
arc(c, d).

osiagalny(X, Y) :- osiagalny(X, Y, []).

osiagalny(X, X, _).
osiagalny(X, Y, Visited) :-
    arc(X, Z),
    \+ member(Z, Visited),
    osiagalny(Z, Y, [X|Visited]).
