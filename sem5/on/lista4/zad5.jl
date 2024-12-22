import Pkg
Pkg.add("Plots")
using Plots

# Definiujemy funkcję f(x) = e^x (w Julii: exp(x)):
f1(x) = exp(x)

# Testy:
rysujNnfx(f1, 0.0, 1.0, 5)
rysujNnfx(f1, 0.0, 1.0, 10)
rysujNnfx(f1, 0.0, 1.0, 15)

# Definiujemy funkcję f(x) = x^2 * sin(x):
f2(x) = x^2 * sin(x)

# Testy:
rysujNnfx(f2, -1.0, 1.0, 5)
rysujNnfx(f2, -1.0, 1.0, 10)
rysujNnfx(f2, -1.0, 1.0, 15)