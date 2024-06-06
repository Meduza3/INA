anbn --> ''.
anbn --> 'a', anbn, 'b'.

anbncn --> licznik(X, 'a'), licznik(X, 'b'), licznik(X, 'c').

licznik(0, _) --> ''.

licznik(s(X), L) -->
  licznik(X, L),
  L.

anbfibn --> licznik(X, 'a'), fibo_licznik(X, 'b').

fibo_licznik(0, _) --> ''.
fibo_licznik(s(0), L) --> L.
fibo_licznik(s(s(X)), L) --> fibo_licznik(X, L), fibo_licznik(s(X), L).