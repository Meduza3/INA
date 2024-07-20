:- use_module(library(clpfd)).

plecak(Wartości, Wielkości, Pojemność, Zmienne, Wartość) :-
    same_length(Wartości, Wielkości),
    same_length(Wartości, Zmienne),
    Zmienne ins 0..1,
    scalar_product(Wielkości, Zmienne, #=<, Pojemność), % Zapakowanie nie przekracza Pojemności
    scalar_product(Wartości, Zmienne, #=, Wartość), % Wylicz ile wartości się zmieściło
    once(labeling([max(Wartość)], Zmienne)). % wyprintuj zmienne z najwyzsza wartosc

% ?- plecak([10,7,1,3,20,2],[9,12,2,15,7,5], 40, X).
