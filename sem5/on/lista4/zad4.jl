using Plots

"""
    rysujNnfx(f, a::Float64, b::Float64, n::Int)

Interpoluje funkcję `f` w przedziale [a, b] za pomocą wielomianu 
interpolacyjnego Newtona stopnia n i rysuje wykres oryginalnej funkcji 
oraz wielomianu.

Argumenty:
- f  : anonimowa funkcja (lub dowolna funkcja jednowymiarowa f(x))
- a  : lewa granica przedziału [a, b]
- b  : prawa granica przedziału [a, b]
- n  : stopień wielomianu interpolacyjnego

Wykorzystuje funkcje:
- `ilorazyRoznicowe(x, fvals)`  (z zad. 1),
- `warNewton(x, fx, t)`        (z zad. 2).

Uwaga:
- Należy najpierw zdefiniować/zaimportować `ilorazyRoznicowe` i `warNewton`.
- Funkcja nie zwraca bezpośrednio obiektu, a od razu rysuje wykres (side effect).
"""
function rysujNnfx(f::Function, a::Float64, b::Float64, n::Int)
    # 1. Wyznaczamy węzły równoodległe:
    h = (b - a) / n
    x = [a + k*h for k in 0:n]                 # wektor węzłów x0, x1, ..., xn

    # 2. Obliczamy wartości f w węzłach:
    fvals = [f(x[k+1]) for k in 0:n]           # f(x0), f(x1), ..., f(xn)

    # 3. Obliczamy ilorazy różnicowe (c0, c1, ..., cn):
    fx = ilorazyRoznicowe(x, fvals)

    # 4. Przygotowujemy siatkę punktów do rysowania (gęstsza niż węzły):
    T = range(a, b, length=300)               # 300 punktów w [a, b]
    # 5. Obliczamy wartości oryginalnej funkcji oraz wielomianu
    fT = [f(t) for t in T]
    pT = [warNewton(x, fx, t) for t in T]

    # 6. Rysujemy: najpierw oryginalną funkcję f, potem wielomian:
    #    używamy Plots.jl (funkcja plot! dodaje serię do istniejącego wykresu)
    plot(
        T, fT,
        label = "f(x)",
        linewidth = 2,
        title = "Interpolacja funkcji f(x) wielomianem Newtona stopnia $n",
        xlabel = "x", ylabel = "y",
        legend = :top
    )
    plot!(
        T, pT,
        label = "Wielomian interpolacyjny Nₙ(x)",
        linewidth = 2,
        linestyle = :dash
    )

    # 7. Opcjonalnie można dodać węzły jako punkty
    scatter!(x, fvals, label="Węzły interpolacji", markersize=4, color=:black)

    # 8. Dzięki Plots wystarczy wywołać powyższe polecenia i wykres pojawi się w sesji REPL/Jupyter.
end
