function naturalna(x::Vector{Float64}, fx::Vector{Float64})
  """
  Funkcja oblicza współczynniki a_0, a_1, ..., a_n (w wektorze `a`)
  wielomianu interpolacyjnego w postaci naturalnej, mając podane:
  - węzły x[1] = x0, x[2] = x1, ..., x[n+1] = x_n
  - ilorazy różnicowe fx[1] = c0, fx[2] = c1, ..., fx[n+1] = cn
    (gdzie c_j = f[x0, x1, ..., xj]).

  Zwraca wektor współczynników [a0, a1, ..., an] w czasie O(n^2).

  Parametry:
  - x  :: Vector{Float64}  -- węzły, długość n+1
  - fx :: Vector{Float64}  -- ilorazy różnicowe, długość n+1

  Wynik:
  - a  :: Vector{Float64}  -- współczynniki postaci naturalnej, długość n+1
                             a[1] = a0, a[2] = a1, ..., a[n+1] = an
  """

  n = length(x) - 1   # stopień wielomianu
  # Wektor wynikowy (postaci naturalnej), początkowo same zera
  a = zeros(n+1)

  # p[j] będzie przechowywał współczynniki wielomianu P_j(x):
  #   P_j(x) = sum_{k=0}^j p[j][k] * x^k
  # Tworzymy listę (lub wektor wektorów) p, gdzie p[j] jest długości j+1.
  p = Vector{Vector{Float64}}(undef, n+1)

  # P_0(x) = 1  =>  jedyny współczynnik to 1 przy x^0
  p[1] = [1.0]   # uwaga: p[1] odpowiada P_0, p[2] odpowiada P_1, itd.

  # Wyznaczamy kolejno P_1, P_2, ..., P_n w O(n^2):
  # P_j(x) = (x - x_{j-1}) * P_{j-1}(x)
  for j in 1:n
      # Stary wielomian: P_{j-1}
      oldP = p[j]            # wektor współczynników P_{j-1}, długość j
      newP = zeros(j+1)      # wektor współczynników P_j, długość (j+1)

      # Mnożymy P_{j-1}(x) przez (x - x_{j-1}).
      # Indeks j w Julii odpowiada j w matematyce, ale:
      #   x_{j-1} to x[j], bo x[1] = x0, x[2] = x1, ...
      α = x[j]  # to jest x_{j-1} w "klasycznej" notacji

      # (x - α) * (p0 + p1 x + ... + p_{j-1} x^{j-1})
      #  = p0*x + p1*x^2 + ... + p_{j-1}*x^j
      #    - α*(p0 + p1*x + ... + p_{j-1} x^{j-1})

      # Część mnożenia przez x:
      for k in 1:j
          newP[k+1] += oldP[k]  # przesunięcie współcz. o 1 w górę stopnia
      end

      # Część mnożenia przez (-α):
      for k in 1:j
          newP[k] -= α * oldP[k]
      end

      # Zapisujemy P_j
      p[j+1] = newP
  end

  # Teraz każdy P_j(x) mamy w p[j+1], j=0..n.
  # A współczynniki c_j (ilorazy różnicowe) to fx[j+1].
  # Wielomian w postaci Newtona: sum_{j=0..n} c_j * P_j(x).
  # Dodajemy c_j * p[j+1] do wektora wynikowego a.

  for j in 0:n
      c_j = fx[j+1]      # c_j
      polyPj = p[j+1]    # P_j(x) w postaci współczynników

      # Dodajemy c_j * p_{j,k} do a[k]
      for k in 0:j
          a[k+1] += c_j * polyPj[k+1]
      end
  end

  return a
end
