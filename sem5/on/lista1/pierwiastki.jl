using Printf

# Definiujemy zakres k od 1 do 20
k_values = 1:20

x_values = Float64[]
f_values = Float64[]
g_values = Float64[]

for k in k_values
    x = 8.0^(-k)
    f_x = sqrt(x^2 + 1) - 1
    g_x = x^2 / (sqrt(x^2 + 1) + 1)
    
    push!(x_values, x)
    push!(f_values, f_x)
    push!(g_values, g_x)
end


println(@sprintf("%-5s %-22s %-22s %-22s", "k", "x", "f(x)", "g(x)"))
for i in 1:length(k_values)
    println(@sprintf("%-5d %-22.15e %-22.15e %-22.15e", k_values[i], x_values[i], f_values[i], g_values[i]))
end
