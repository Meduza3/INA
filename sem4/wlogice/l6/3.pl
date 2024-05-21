anbn --> ''.
anbn --> 'a', anbn, 'b'.

anbncn --> licznik(X, 'a'), licznik(X, 'b'), licznik(X, 'c').

licznik(0, _) --> ''.

licznik(s(X), L) -->
  licznik(X, L),
  L.

anbfibn --> licznik(X, 'a'), fibCounter(X, 'b').

fibCounter(0, _) --> ''.
fibCounter(s(0), L) --> L.
fibCounter(s(s(X)), L) --> fibCounter(X, L), fibCOunter(s(X), L).

