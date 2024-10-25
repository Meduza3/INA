function znajdź_min_x()
  x = 1.0
  while x < 2.0
      x = nextfloat(x)
      y = x * (1.0 / x)
      if y != 1.0
          println("Znaleziono x = $x, dla którego x * (1 / x) ≠ 1")
          println("Wartość x * (1 / x) = $y")
          break
      end
  end
end

znajdź_min_x()
