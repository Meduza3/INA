# Definiujemy funkcję f(x) = |x| (w Julii: abs(x)):
f3(x) = abs(x)

# Testy:
rysujNnfx(f3, -1.0, 1.0, 5)
rysujNnfx(f3, -1.0, 1.0, 10)
rysujNnfx(f3, -1.0, 1.0, 15)

# Definiujemy funkcję f(x) = 1/(1 + x^2):
f4(x) = 1 / (1 + x^2)

# Testy:
rysujNnfx(f4, -5.0, 5.0, 5)
rysujNnfx(f4, -5.0, 5.0, 10)
rysujNnfx(f4, -5.0, 5.0, 15)
