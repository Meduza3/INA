inversion_count([], _, 0).
inversion_count([H | T], List, Count) :-
    findall(X, (member(X, List), X < H), Smaller),
    length(Smaller, CountSmaller),
    inversion_count(T, List, CountRest),
    Count is CountSmaller + CountRest.

even(0).
even(X) :- 
    X > 0, 
    Y is X -2, 
    even(Y).

odd(X) :- \+ even(X).

even_permutation(Xs, Ys) :-
    permutation(Xs, Ys),
    inversion_count(Ys, Xs, Count),
    even(Count).

odd_permutation(Xs, Ys) :-
    permutation(Xs, Ys),
    inversion_count(Ys, Xs, Count),
    odd(Count).

