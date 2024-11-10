p0 = Float32(0.01)
r = Float32(3.0)
n_iter = 40

# Funkcja do obliczania iteracji bez modyfikacji
function logistic_map(p0, r, n_iter)
    p = zeros(Float32, n_iter+1)
    p[1] = p0
    for i in 1:n_iter
        p[i+1] = p[i] + r * p[i] * (1 - p[i])
    end
    return p
end

# Funkcja do obliczania iteracji z obcięciem po 10-tej iteracji
function logistic_map_truncated(p0, r, n_iter, truncate_at)
    p = zeros(Float32, n_iter+1)
    p[1] = p0
    for i in 1:n_iter
        p_value = p[i] + r * p[i] * (1 - p[i])
        if i == truncate_at
            # Obcięcie wyniku po trzecim miejscu po przecinku
            p_value = floor(p_value * 1000) / 1000
        end
        p[i+1] = p_value
    end
    return p
end

# Wykonanie obliczeń
p_original = logistic_map(p0, r, n_iter)
p_float64 = logistic_map(Float64(0.01), Float64(3.0), 40)
p_truncated = logistic_map_truncated(p0, r, n_iter, 10)

# Porównanie wyników
println("Iteracja & Float32 & Float32 z obcięciem & Float64")
for i in 1:n_iter+1
    println("$(i-1) & $(p_original[i]) & $(p_truncated[i]) & $(p_float64[i]) \\\\ \\hline")
end