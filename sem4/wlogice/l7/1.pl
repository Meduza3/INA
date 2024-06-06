merge([], IN2, IN2).
merge(IN1, [], IN1).

merge([H1|R1], [H2|R2], [H1|P]) :-
    H1 =< H2,
    merge(R1, [H2|R2], P).

merge([H1|R1], [H2|R2], [H2|P]) :-
    H1 > H2,
    merge([H1|R1], R2, P).