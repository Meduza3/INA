

function mbisekcji(f, a::Float64, b::Float64, delta::Float64, epsilon::Float64)
  u = f(a)
  v = f(b)
  it = 0

  if sign(u) == sign(v) 
    return (NaN, NaN, it, 1)
  end

  r = (a + b) / 2.0

  while (b - a) / 2.0 > delta && abs(f(r)) > epsilon
    it += 1
    if f(a) * f(r) < 0
      b = r
    else
      a = r
    end
    r = (a + b) / 2.0
  end

  v = f(r)
  err = 0

  return (r, v, it, err)
end