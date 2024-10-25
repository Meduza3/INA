function compute_macheps(T)
  oneT = one(T)
  macheps = oneT
  prev_macheps = macheps
  while oneT + macheps > oneT
      prev_macheps = macheps
      macheps = macheps / T(2)
  end
  return prev_macheps
end

println("Wyliczony macheps dla Float16: ", compute_macheps(Float16), " eps(Float16): ", eps(Float16))
println("Wyliczony macheps dla Float32: ", compute_macheps(Float32), " eps(Float32): ", eps(Float32))
println("Wyliczony macheps dla Float64: ", compute_macheps(Float64), " eps(Float64): ", eps(Float64))


println("Wartości z float.h:")
println("FLT_EPSILON: ", 1.19209290e-07)
println("DBL_EPSILON: ", 2.2204460492503131e-16)