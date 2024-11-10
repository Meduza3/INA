# Marcin Zubrzycki

import Pkg;
Pkg.add("Polynomials")


using Polynomials

p = [1, -210.0, 20615.0,-1256850.0, 53327946.0, -1672280820.0,
    40171771630.0, -756111184500.0, 11310276995381.0, -135585182899530.0, 
    1307535010540395.0, -10142299865511450.0, 63030812099294896.0,
    -311333643161390640.0, 1206647803780373360.0, -3599979517947607200.0, 
    8037811822645051776.0, -12870931245150988800.0, 13803759753640704000.0,      
    -8752948036761600000.0, 2432902008176640000.0]

println("---- Podpunkt (a) - przed naniesieniem zaburzenia: ----\n")
wilkinson_normal = Polynomial(reverse(p))       # postać naturalna
wilkinson_factored = fromroots(1:20)            # postać iloczynowa
wilkinson_roots = roots(wilkinson_normal)

println(rpad("k", 3),
        " & ", rpad("\$z_k\$", 20), 
        " & ", rpad("\$|P(z_k)|\$", 25), 
        " & ", rpad("\$|p(z_k)|\$", 25), 
        " & ", rpad("\$|z_k - k|\$", 25), 
        " \\\\\\hline")

for k in 1:20
    z_k = wilkinson_roots[k]
    println(rpad(k, 3), 
            " & ", rpad(z_k, 20), 
            " & ", rpad(abs(wilkinson_normal(z_k)), 25), 
            " & ", rpad(abs(wilkinson_factored(z_k)), 25), 
            " & ", rpad(abs(z_k - k), 25), 
            " \\\\\\hline")
end

println("\n---- Dodatkowy eksperyment do podpunktu (a): ----\n")
println(rpad("k", 3), 
        " & ", rpad("\$|P(k)|\$", 25), 
        " & ", rpad("\$|p(k)|\$", 10), 
        " \\\\\\hline")

for k in 1:20
    println(rpad(k, 3), 
            " & ", rpad(abs(wilkinson_normal(k)), 25), 
            " & ", rpad(abs(wilkinson_factored(k)), 10), 
            " \\\\\\hline")
end

println("\n---- Eksperyment z zaburzeniem - podpunkt (b): ----\n")
p[2] -= 2.0^-23 
wilkinson_noise = Polynomial(reverse(p)) # postać naturalna z zaburzeniem
new_roots = roots(wilkinson_noise)

println(rpad("k", 3),
        " & ", rpad("\$z_k\$", 50), 
        " & ", rpad("\$|P(z_k)|\$", 25), 
        " \\\\\\hline")

for k in 1:20
    z_k = new_roots[k]
    println(rpad(k, 3), 
            " & ", rpad(z_k, 50),
            " & ", rpad(abs(wilkinson_noise(z_k)), 25),
            " \\\\\\hline")
end