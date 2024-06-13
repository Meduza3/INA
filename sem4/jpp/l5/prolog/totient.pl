consult("prime_factors.pl").

% unique_factors(List, Unique) - zwraca zbiór elementów z List.
unique_factors(List, Unique) :-
    sort(List, Unique).

% compute_totient(N, Factors, T) - oblicza wartość totient na podstawie unikalnych czynników pierwszych.
compute_totient(N, [], N).
compute_totient(N, [F|Fs], T) :-
    compute_totient(N, Fs, T1),
    T is T1 * (F - 1) // F.

% totient(N, T) - T to wartość funkcji Eulera dla N.
totient(N, T) :-
    N > 0,
    prime_factors(N, Factors),
    unique_factors(Factors, UniqueFactors),
    compute_totient(N, UniqueFactors, T).
