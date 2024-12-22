function warNewton(x::Vector{Float64}, fx::Vector{Float64}, t::Float64)
  """
  Funkcja oblicza wartość wielomianu interpolacyjnego w postaci Newtona
  w punkcie `t` za pomocą uogólnionego algorytmu Hornera.

  Argumenty:
  - x  : wektor węzłów [x_0, x_1, ..., x_n] (długość n+1)
  - fx : wektor ilorazów różnicowych [f[x_0], f[x_0,x_1], ..., f[x_0,...,x_n]]
         (długość n+1)
  - t  : punkt, w którym obliczamy wartość wielomianu

  Zwracana wartość:
  - nt : wartość wielomianu Newtona w punkcie `t`.
  """

  n = length(x) - 1            # ponieważ x ma n+1 elementów, to "stopień" wynosi n
  nt = fx[n+1]                 # zaczynamy od najwyższego rzędu ilorazu różnicowego

  # Schemat Hornera dla wielomianu w postaci Newtona:
  # N_n(t) = fx[n+1]
  #        + (t - x[n]) * [ fx[n]
  #                        + (t - x[n-1]) * [ fx[n-1]
  #                                          + (t - x[n-2]) * [...]
  #                                          ]
  #                       ]
  
  for i in n:-1:1
      nt = fx[i] + (t - x[i]) * nt
  end

  return nt
end
