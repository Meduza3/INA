is_prime(2).
is_prime(N) :-
    N > 2,
    \+ has_divisor(N, 2).

has_divisor(N, D) :-
    N mod D =:= 0.
has_divisor(N, D) :-
    D * D < N,
    D2 is D + 1,
    has_divisor(N, D2).

prime(LO, HI, N) :- between(LO, HI, N), is_prime(N).
