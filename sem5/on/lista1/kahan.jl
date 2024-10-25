# Obliczamy wyrażenie dla Float64
x64 = 3.0 * (4.0 / 3.0 - 1.0) - 1.0
println("Wynik dla Float64: ", x64)
eps64 = eps(Float64)
println("Epsilon maszynowy dla Float64: ", eps64)

# Obliczamy wyrażenie dla Float32
x32 = 3.0f0 * (4.0f0 / 3.0f0 - 1.0f0) - 1.0f0
println("Wynik dla Float32: ", x32)
eps32 = eps(Float32)
println("Epsilon maszynowy dla Float32: ", eps32)

# Obliczamy wyrażenie dla Float16
x16 = 3.0f16 * (4.0f16 / 3.0f16 - 1.0f16) - 1.0f16
println("Wynik dla Float16: ", x16)
eps16 = eps(Float16)
println("Epsilon maszynowy dla Float16: ", eps16)
