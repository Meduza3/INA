include("mbisekcji.jl")
include("msiecznych.jl")
include("mstycznych.jl")


# Zadanie 6
f1 = x -> exp(1 - x) - 1
pf1 = x -> -exp(1 - x)

f2 = x -> x * exp(-x)
pf2 = x -> exp(-x) * (1 - x)

delta = 1e-5
epsilon = 1e-5

println("Metoda bisekcji:")
r1_b, v1_b, it1_b, err1_b = mbisekcji(f1, -2.0, 2.0, delta, epsilon)
println("f1: Pierwiastek: $r1_b, Wartość: $v1_b, Iteracje: $it1_b, Błąd: $err1_b")

r2_b, v2_b, it2_b, err2_b = mbisekcji(f2, 0.0, 1.0, delta, epsilon)
println("f2: Pierwiastek: $r2_b, Wartość: $v2_b, Iteracje: $it2_b, Błąd: $err2_b")

println("\nMetoda Newtona:")
r1_n, v1_n, it1_n, err1_n = mstycznych(f1, df1, 1.5, delta, epsilon, 100)
println("f1: Pierwiastek: $r1_n, Wartość: $v1_n, Iteracje: $it1_n, Błąd: $err1_n")

r2_n, v2_n, it2_n, err2_n = mstycznych(f2, df2, 0.5, delta, epsilon, 100)
println("f2: Pierwiastek: $r2_n, Wartość: $v2_n, Iteracje: $it2_n, Błąd: $err2_n")

println("\nMetoda siecznych:")
r1_s, v1_s, it1_s, err1_s = msiecznych(f1, 0.5, 1.5, delta, epsilon, 100)
println("f1: Pierwiastek: $r1_s, Wartość: $v1_s, Iteracje: $it1_s, Błąd: $err1_s")

r2_s, v2_s, it2_s, err2_s = msiecznych(f2, 0.2, 0.8, delta, epsilon, 100)
println("f2: Pierwiastek: $r2_s, Wartość: $v2_s, Iteracje: $it2_s, Błąd: $err2_s")