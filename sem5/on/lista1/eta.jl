function compute_eta(T)
  x = one(T)
  while x / 2 > zero(T)
      x = x / 2
  end
  return x
end

eta16 = compute_eta(Float16)
eta32 = compute_eta(Float32)
eta64 = compute_eta(Float64)

println("Eta dla Float16: ", eta16)
println("Eta dla Float32: ", eta32)
println("Eta dla Float64: ", eta64)

println("nextfloat(0.0) dla Float16: ", nextfloat(Float16(0.0)))
println("nextfloat(0.0) dla Float32: ", nextfloat(Float32(0.0)))
println("nextfloat(0.0) dla Float64: ", nextfloat(Float64(0.0)))
