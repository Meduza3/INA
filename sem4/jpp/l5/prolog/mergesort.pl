% Dzieli listę na dwie części
% split(List, Left, Right)
split([], [], []).
split([X], [X], []).
split([X, Y | Rest], [X | L], [Y | R]) :-
    split(Rest, L, R).

% Scala dwie posortowane listy w jedną
% merge(Left, Right, Merged)

merge([], R, R).
merge(L, [], L).
merge([X | L1], [Y | R1], [X | M]) :-
    X =< Y,
    merge(L1, [Y | R1], M).
merge([X | L1], [Y | R1], [Y | M]) :-
    X > Y,
    merge([X | L1], R1, M).

% Sortuje
% mergesort(List, Sorted)
mergesort([], []).
mergesort([X], [X]).
mergesort(List, Sorted) :-
    List = [_, _ | _],  % Lista ma co najmniej dwa elementy
    split(List, Left, Right),
    mergesort(Left, SortedLeft),
    mergesort(Right, SortedRight),
    merge(SortedLeft, SortedRight, Sorted).

