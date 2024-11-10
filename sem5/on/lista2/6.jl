# Define the list of cases: (c, x0)
cases = [
    (-2.0, 1.0),
    (-2.0, 2.0),
    (-2.0, 1.99999999999999),
    (-1.0, 1.0),
    (-1.0, -1.0),
    (-1.0, 0.75),
    (-1.0, 0.25)
]

# Number of iterations
num_iterations = 40

# Loop over each case
for (index, (c, x0)) in enumerate(cases)
    println("Case $(index): c = $(c), x0 = $(x0)")
    # Initialize the sequence
    x = Float64[]
    push!(x, x0)
    # Perform iterations
    for n in 1:num_iterations
        x_new = x[end]^2 + c
        push!(x, x_new)
    end
    # Print the sequence
    println("    \\hline
    Case $(index): \$c = $(c), x_0 = $(x0)\$ \\\\
        \\hline")
    for (n, value) in enumerate(x)
        println("$(n - 1) & $(value) \\\\\\hline")
    end
    println("---------------------------------------------------")
end
