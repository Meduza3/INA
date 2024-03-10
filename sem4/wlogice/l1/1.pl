ojciec(X, Y). /* X jest ojcem Y */
matka(X, Y). /* X jest matką Y */
mezczyzna(X). /* X jest mężczyzną */
kobieta(X). /* X jest kobietą */
rodzic(X, Y). /* X jest rodzicem Y */

jest_matka(X) :- matka(X, _).
jest_ojcem(X) :- ojciec(X, _).
jest_synem(X) :- rodzic(Y, X), mezczyzna(X).
siostra(X, Y) :- kobieta(X), rodzic(Z, Y), rodzic(Z, Y), X \= Y .
dziadek(X, Y) :- rodzic(Z, X), ojciec(Y, Z).
rodzenstwo(X, Y) :- rodzic(Z, Y), rodzic(Z, Y), X \= Y.
