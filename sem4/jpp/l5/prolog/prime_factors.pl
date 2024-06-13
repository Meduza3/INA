% smallest_factor(N, F) - F jest najmniejszym czynnikiem pierwszym N.
smallest_factor(N, F) :- smallest_factor(N, 2, F).

smallest_factor(N, F, F) :- N mod F =:= 0, !.
smallest_factor(N, F, R) :-
    F * F < N, % Jeśli F^2 jest mniejsze niż N, kontynuuj szukanie
    F1 is F + 1,
    smallest_factor(N, F1, R).
smallest_factor(N, F, N) :- F * F >= N. % Jeśli F^2 >= N, to N jest pierwsze

% prime_factors(N, X) - X to lista czynników pierwszych liczby N.
prime_factors(N, X) :-
    N > 1,
    prime_factors(N, [], X).

% prime_factors(N, Acc, X) - pomocniczy predykat z akumulatorem.
prime_factors(1, Acc, Acc).
prime_factors(N, Acc, X) :-
    N > 1,
    smallest_factor(N, F),
    N1 is N // F,
    prime_factors(N1, [F|Acc], X).
