# Marcin Zubrzycki

using LinearAlgebra
using Random
using Statistics

# Zaimportowana funkcja ze strony https://cs.pwr.edu.pl/zielinski/lectures/scna/hilb.jl
function hilb(n::Int)
  # Function generates the Hilbert matrix  A of size n,
  #  A (i, j) = 1 / (i + j - 1)
  # Inputs:
  #	n: size of matrix A, n>=1
  #
  #
  # Usage: hilb(10)
  #
  # Pawel Zielinski
          if n < 1
           error("size n should be >= 1")
          end
          return [1 / (i + j - 1) for i in 1:n, j in 1:n]
  end

  function matcond(n::Int, c::Float64)
    # Function generates a random square matrix A of size n with
    # a given condition number c.
    # Inputs:
    #	n: size of matrix A, n>1
    #	c: condition of matrix A, c>= 1.0
    #
    # Usage: matcond(10, 100.0)
    #
    # Pawel Zielinski
            if n < 2
             error("size n should be > 1")
            end
            if c< 1.0
             error("condition number  c of a matrix  should be >= 1.0")
            end
            (U,S,V)=svd(rand(n,n))
            return U*diagm(0 =>[LinRange(1.0,c,n);])*V'
    end
    
    
    
    

  function gauss()
  end

  function invalgo()
  end

  function relative_error(x_approx::Vector{Float64}, x_exact::Vector{Float64})
    return norm(x_approx - x_exact) / norm(x_exact)
end

function experiment(A::Matrix{Float64}, x_exact::Vector{Float64})
  b = A * x_exact

  # Rozwiązanie za pomocą eliminacji Gaussa
  x_gauss = A \ b
  error_gauss = relative_error(x_gauss, x_exact)

  # Rozwiązanie za pomocą odwracania macierzy
  x_inv = inv(A) * b
  error_inv = relative_error(x_inv, x_exact)
  return error_gauss, error_inv
end

# Eksperymenty dla macierzy Hilberta
println("=== Eksperymenty dla macierzy Hilberta ===")
println("n \t Błąd Gauss \t Błąd inv")
for n in 2:15
    A = hilb(n)
    x_exact = ones(n)
    error_gauss, error_inv = experiment(A, x_exact)
    println("$n \t $(round(error_gauss, sigdigits=3)) \t\t $(round(error_inv, sigdigits=3))")
end

# Eksperymenty dla macierzy losowych z zadanym wskaźnikiem uwarunkowania
println("\n=== Eksperymenty dla macierzy losowych ===")
ns = [5, 10, 20]
cs = [1, 10, 1e3, 1e7, 1e12, 1e16]

println("n \t c \t\t Błąd Gauss \t Błąd inv")
for n in ns
    for c in cs
        A = matcond(n, c)
        x_exact = ones(n)
        error_gauss, error_inv = experiment(A, x_exact)
        println("$n \t $(c < 1e5 ? Int(c) : c) \t\t  $(round(error_gauss, sigdigits=3)) \t\t $(round(error_inv, sigdigits=3))")
    end
end