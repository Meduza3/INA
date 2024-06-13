% sieve(Numbers, Primes) - zwraca listę liczb pierwszych z listy Numbers.
sieve([], []).
sieve([First|Rest], [First|Primes]) :-
    cross_out(Rest, First, NewRest),
    sieve(NewRest, Primes).

% cross_out(Numbers, N, Result) - usuwa wielokrotności liczby N z listy Numbers.
cross_out([], _, []).
cross_out([H|T], N, R) :-
    (H mod N =:= 0 -> 
        cross_out(T, N, R); 
        R = [H|RT], cross_out(T, N, RT)).

% list_from_to(Low, High, List) - tworzy listę liczb od Low do High.
list_from_to(Low, High, []) :- Low > High.
list_from_to(Low, High, [Low|Tail]) :-
    Low =< High,
    Next is Low + 1,
    list_from_to(Next, High, Tail).

% primes(N, Primes) - zwraca listę liczb pierwszych od 2 do N.
primes(N, Primes) :-
    N >= 2,
    list_from_to(2, N, Numbers),
    sieve(Numbers, Primes).
