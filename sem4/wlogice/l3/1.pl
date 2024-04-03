suma_liczb([], 0, 0).
suma_liczb([H | T], Suma, Liczba) :-
    suma_liczb(T, ResztaSumy, ResztaLiczby),
    Suma is H + ResztaSumy,
    Liczba is ResztaLiczby + 1.


srednia(Lista, Srednia) :-
    suma_liczb(Lista, Suma, Liczba),
    Srednia is Suma / Liczba.

suma_kwadratow_odchylen([], _, 0).
suma_kwadratow_odchylen([H | T], Srednia, Suma) :-
    suma_kwadratow_odchylen(T, Srednia, ResztaSumy),
    Odchylenie is H - Srednia,
    Suma is ResztaSumy + Odchylenie * Odchylenie.

wariancja(Lista, Wariancja) :-
    srednia(Lista, Srednia),
    suma_kwadratow_odchylen(Lista, Srednia, SumaKwadratow),
    length(Lista, N),
    Wariancja is SumaKwadratow / (N).