function compute_max(T)
  x = one(T)
  prev_x = x
  while !isinf(x)
      prev_x = x
      x = x * T(2.0)
  end
  x = prev_x
  delta = x / T(2.0)
  while delta != zero(T)
      if isinf(x + delta)
          delta /= T(2.0)
      else
          x += delta
          delta /= T(2.0)
      end
  end
  return x
end

max16 = compute_max(Float16)
max32 = compute_max(Float32)
max64 = compute_max(Float64)

println("MAX dla Float16: ", max16)
println("floatmax(Float16): ", floatmax(Float16))
println()

println("MAX dla Float32: ", max32)
println("floatmax(Float32): ", floatmax(Float32))
println()

println("MAX dla Float64: ", max64)
println("floatmax(Float64): ", floatmax(Float64))
