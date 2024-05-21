:- consult('1.pl'), consult('interpreter.pl').

wykonaj(NazwaPliku) :-
  open(NazwaPliku, read, X),
  scanner(X, Y),
  close(X),
  phrase(program(Z), Y),
  interpreter(Z).