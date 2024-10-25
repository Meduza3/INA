for k in 0:5
  x = 1.0 + k * 2.0^-52
  println("x = ", x, ", bitstring = ", bitstring(x))
end
