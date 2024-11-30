include("mbisekcji.jl")

# Funkcja opisująca różnicę między y = 3x i y = e^x
f = x -> 3x - exp(x)
a, b = 0.0, 1.0
delta = 1e-4
epsilon = 1e-4

r, v, it, err = mbisekcji(f, a, b, delta, epsilon)

if err == 0
    println("Przybliżenie pierwiastka: $r")
    println("Wartość funkcji w przybliżeniu: $v")
    println("Liczba iteracji: $it")
else
    println("Błąd: Funkcja nie zmienia znaku w podanym przedziale.")
end