function msiecznych(f, x0::Float64, x1::Float64, delta::Float64, epsilon::Float64, maxit::Int)
  r = x1
  v = f(x1)
  it = 0
  err = 0

  for i in 1:maxit
      it += 1
      f_x0 = f(x0)
      f_x1 = f(x1)

      if abs(f_x1) < epsilon
          return (x1, f_x1, it, err)
      end

      if abs(x1 - x0) < delta
          return (x1, f_x1, it, err)
      end

      if abs(f_x1 - f_x0) < epsilon
          return (r, v, it, 1)
      end

      x_new = x1 - f_x1 * (x1 - x0) / (f_x1 - f_x0)
      x0, x1 = x1, x_new
  end

  return (x1, f(x1), it, 1)  # Error if maxit reached
end
