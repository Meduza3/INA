function mstycznych(f, pf, x0::Float64, delta::Float64, epsilon::Float64, maxit::Int)
  x_prev = x0
  err = 1
  tolerance = 1e-10

  for it = 1:maxit
    f_val = f(x_prev)
    pf_val = pf(x_prev)

    if abs(pf_val) < tolerance
      r = x_prev
      v = f_val
      err = 2
      return (r, v, it, err)
    end

    x_next = x_prev - f_val / pf_val

    if abs(x_next - x_prev) < delta && abs(f(x_next)) < epsilon
      r = x_next
      v = f(x_next)
      err = 0
      return (r, v, it, err)
    end

    x_prev = x_next
  end

  r = x_prev
  v = f(x_prev)
  err = 1
  return (r, v, maxit, err)
end
