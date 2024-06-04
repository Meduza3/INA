hetmany(N, P) :-
  numlist(1, N, L),
  permutation(L, P),
  dobra(P).

dobra(P) :- 
         \+ zla(P).

zla(P) :-
  append(_, [Wi | L1], P),
  append(L2, [Wj | _], L1),
  length(L2, K),
  abs(Wi - Wj) =:= K + 1.



print_line(CurrentRowIndex, MaxIndex) :-
    CurrentRowIndex >= MaxIndex,
    write('+-----+'),
    !.

print_line(CurrentRowIndex, MaxIndex) :-
    write('+-----'),
    print_line(CurrentRowIndex + 1, MaxIndex).


print_inner_segment_black(CurrentRowIndex, MaxIndex, [HetmanPos|_], Level) :-
    CurrentRowIndex >= MaxIndex,
    (
        Level =:= HetmanPos ->
        write('|:###:|') ; write('|:::::|')
    ),
    !.

print_inner_segment_black(CurrentRowIndex, MaxIndex, [HetmanPos|Reszta], Level) :-
    (
        Level =:= HetmanPos ->
        write('|:###:') ; write('|:::::')
    ),
    print_inner_segment_white(CurrentRowIndex + 1, MaxIndex, Reszta, Level).

print_inner_segment_white(CurrentRowIndex, MaxIndex, [HetmanPos|_], Level) :-
    CurrentRowIndex >= MaxIndex,
    (
        Level =:= HetmanPos ->
        write('| ### |') ; write('|     |')
    ),
    !.

print_inner_segment_white(CurrentRowIndex, MaxIndex, [HetmanPos|Reszta], Level) :-
    (
        Level =:= HetmanPos ->
        write('| ### ') ; write('|     ')
    ),
    print_inner_segment_black(CurrentRowIndex + 1, MaxIndex, Reszta, Level).


%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

print_rows_black(CurrentRowIndex, BoardSize, _) :-
    CurrentRowIndex > BoardSize,
    print_line(1, BoardSize), nl,
    !.

print_rows_black(CurrentRowIndex, BoardSize, HetmanPoss) :-
    print_line(1, BoardSize), nl,
    print_inner_segment_black(1, BoardSize, HetmanPoss, BoardSize - CurrentRowIndex + 1), nl,
    print_inner_segment_black(1, BoardSize, HetmanPoss, BoardSize - CurrentRowIndex + 1), nl,
    print_rows_white(CurrentRowIndex + 1, BoardSize, HetmanPoss).

print_rows_white(CurrentRowIndex, BoardSize, _) :-
    CurrentRowIndex > BoardSize,
    print_line(1, BoardSize), nl,
    !.

print_rows_white(CurrentRowIndex, BoardSize, HetmanPoss) :-
    print_line(1, BoardSize), nl,
    print_inner_segment_white(1, BoardSize, HetmanPoss, BoardSize - CurrentRowIndex + 1), nl,
    print_inner_segment_white(1, BoardSize, HetmanPoss, BoardSize - CurrentRowIndex + 1), nl,
    print_rows_black(CurrentRowIndex + 1, BoardSize, HetmanPoss).


%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

board(HetmanPoss) :-
    length(HetmanPoss, BoardSize),
    (
        BoardSize mod 2 =:= 0 ->
        print_rows_white(1, BoardSize, HetmanPoss) ; print_rows_black(1, BoardSize, HetmanPoss)
    ),
    !.