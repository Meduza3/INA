include("mbisekcji.jl")
include("msiecznych.jl")
include("mstycznych.jl")

f = x -> sin(x) - (0.5x)^2


x0, x1 = 1.0, 2.0
delta = 0.5e-5
epsilon = 0.5e-5
maxit = 100
r, v, it, err = msiecznych(f, x0, x1, delta, epsilon, maxit)

println("METODA SIECZNYCH:")
if err == 0
  println("Przybliżenie pierwiastka: $r")
  println("Wartość funkcji w przybliżeniu: $v")
  println("Liczba iteracji: $it")
else
  println("Błąd: Funkcja nie zmienia znaku w podanym przedziale.")
end

a, b = 1.5, 2.0
delta = 0.5e-5
epsilon = 0.5e-5

r, v, it, err = mbisekcji(f, a, b, delta, epsilon)

println("METODA BISEKCJI:")
if err == 0
  println("Przybliżenie pierwiastka: $r")
  println("Wartość funkcji w przybliżeniu: $v")
  println("Liczba iteracji: $it")
else
  println("Błąd: Funkcja nie zmienia znaku w podanym przedziale.")
end


pf = x -> cos(x) - x
x_0 = 1.5
delta = 0.5e-5
epsilon = 0.5e-5

r, v, it, err = mstycznych(f, pf, x_0, delta, epsilon, 100)

println("METODA STYCZNYCH:")
if err == 0
  println("Przybliżenie pierwiastka: $r")
  println("Wartość funkcji w przybliżeniu: $v")
  println("Liczba iteracji: $it")
else
  println("Błąd: Funkcja nie zmienia znaku w podanym przedziale.")
end