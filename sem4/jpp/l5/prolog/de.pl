% extended_gcd(A, B, X, Y, G) - G jest największym wspólnym dzielnikiem A i B,
% a X i Y są współczynnikami, dla których AX + BY = G.
extended_gcd(0, B, 0, 1, B).
extended_gcd(A, 0, 1, 0, A).
extended_gcd(A, B, X, Y, G) :-
    B > 0,
    Q is A // B,
    R is A mod B,
    extended_gcd(B, R, X1, Y1, G),
    X is Y1,
    Y is X1 - Q * Y1.

% de(A, B, X, Y, Z) - znajduje rozwiązanie równania diofantycznego AX + BY = Z,
% gdzie Z jest największym wspólnym dzielnikiem A i B.
de(A, B, X, Y, Z) :-
    extended_gcd(A, B, X0, Y0, Z),  % Znajdź współczynniki i GCD
    X is X0,
    Y is Y0.
