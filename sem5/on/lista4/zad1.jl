module lista4
function ilorazyRoznicowe(x::Vector{Float64}, f::Vector{Float64})
  """
  Funkcja oblicza kolejne ilorazy różnicowe węzłów x i wartości funkcji f.
  
  Argumenty:
  - x: wektor węzłów x0, x1, ..., xn (długość n+1)
  - f: wektor wartości funkcji f(x0), f(x1), ..., f(xn) (długość n+1)

  Zwracana wartość:
  - fx: wektor ilorazów różnicowych o długości n+1, gdzie:
    fx[1]   = f[x0]
    fx[2]   = f[x0, x1]
    ...
    fx[n]   = f[x0, ..., x_{n-1}]
    fx[n+1] = f[x0, ..., x_n]
  """

  n = length(x)         # n = n+1 w zadaniu (długość obu wektorów)
  fx = copy(f)          # kopiujemy wektor f, żeby operować „w miejscu”

  # W pętli zewnętrznej „j” oznacza, który rząd ilorazu różnicowego obliczamy.
  # W pętli wewnętrznej dokonujemy kolejnych aktualizacji.
  #
  # Schemat:
  # fx[i] = ( fx[i] - fx[i-1] ) / ( x[i] - x[i-j] )
  #
  # Przechodzimy od końca w dół, żeby nie nadpisać przedwcześnie potrzebnych wartości.

  for j in 1:(n-1)
      for i in n:-1:(j+1)
          fx[i] = (fx[i] - fx[i-1]) / (x[i] - x[i-j])
      end
  end

  return fx
end
end