# Definicja wektorów x i y
x = [2.718281828, -3.141592654,  1.414213562, 0.5772156649, 0.3010299957]
y = [1486.2497,   878366.9879,  -22.37492,   4773714.647,  0.000185049]

# Wartość dokładna iloczynu skalarnego
S_exact = -1.00657107000000e-11

# Funkcja do obliczania błędów
function compute_errors(S_computed, S_exact)
    abs_error = abs(S_computed - S_exact)
    rel_error = abs_error / abs(S_exact)
    return abs_error, rel_error
end

# Algorytm (a): Obliczanie iloczynu skalarnego w przód
function scalar_product_forward(x, y)
    S = 0.0
    for i in 1:length(x)
        S += x[i] * y[i]
    end
    return S
end

# Algorytm (b): Obliczanie iloczynu skalarnego w tył
function scalar_product_backward(x, y)
    S = 0.0
    for i in length(x):-1:1
        S += x[i] * y[i]
    end
    return S
end

# Algorytm (c): Od największego do najmniejszego
function scalar_product_sorted_desc(x, y)
    products = x .* y
    positive_products = sort(filter(p -> p > 0, products); rev=true)
    negative_products = sort(filter(p -> p < 0, products))
    S = sum(positive_products) + sum(negative_products)
    return S
end

# Algorytm (d): Od najmniejszego do największego
function scalar_product_sorted_asc(x, y)
    products = x .* y
    positive_products = sort(filter(p -> p > 0, products))
    negative_products = sort(filter(p -> p < 0, products); rev=true)
    S = sum(positive_products) + sum(negative_products)
    return S
end

# Obliczanie iloczynów skalarnych
S_a = scalar_product_forward(x, y)
S_b = scalar_product_backward(x, y)
S_c = scalar_product_sorted_desc(x, y)
S_d = scalar_product_sorted_asc(x, y)

# Obliczanie błędów
E_a, RE_a = compute_errors(S_a, S_exact)
E_b, RE_b = compute_errors(S_b, S_exact)
E_c, RE_c = compute_errors(S_c, S_exact)
E_d, RE_d = compute_errors(S_d, S_exact)

# Wyświetlanie wyników
using Printf

println(@sprintf("Iloczyn skalarny a (algorytm w przód): S = %.16e", S_a))
println(@sprintf("Błąd bezwzględny: E = %.16e", E_a))
println(@sprintf("Błąd względny: RE = %.2e\n", RE_a))

println(@sprintf("Iloczyn skalarny b (algorytm w tył): S = %.16e", S_b))
println(@sprintf("Błąd bezwzględny: E = %.16e", E_b))
println(@sprintf("Błąd względny: RE = %.2e\n", RE_b))

println(@sprintf("Iloczyn skalarny c (od największego do najmniejszego): S = %.16e", S_c))
println(@sprintf("Błąd bezwzględny: E = %.16e", E_c))
println(@sprintf("Błąd względny: RE = %.2e\n", RE_c))

println(@sprintf("Iloczyn skalarny d (od najmniejszego do największego): S = %.16e", S_d))
println(@sprintf("Błąd bezwzględny: E = %.16e", E_d))
println(@sprintf("Błąd względny: RE = %.2e\n", RE_d))

println(@sprintf("Wartość dokładna: S_exact = %.16e", S_exact))
