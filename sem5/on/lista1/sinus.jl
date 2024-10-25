using Printf

f(x) = sin(x) + cos(3x)
f_prime_exact = cos(1) - 3sin(3)

n_values = 0:54
h_values = 2.0 .^ (-n_values)
approx_derivatives = zeros(length(n_values))
errors = zeros(length(n_values))

for (i, n) in enumerate(n_values)
    h = 2.0 ^ (-n)
    f1 = f(1.0)
    f2 = f(1.0 + h)
    approx_derivative = (f2 - f1) / h
    error = abs(f_prime_exact - approx_derivative)
    approx_derivatives[i] = approx_derivative
    errors[i] = error
    @printf("n = %2d, h = %.5e, Approx Derivative = %.10f, Error = %.5e\n",
            n, h, approx_derivative, error)
end
