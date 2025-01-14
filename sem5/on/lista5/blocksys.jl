module blocksys
using SparseArrays
using LinearAlgebra

export solve_block_tridiag!, blockify, join_blocks, solve_block_tridiag_pivot!

function join_blocks(xblocks)
  v = length(xblocks)            # = 4
  l = length(xblocks[1])         # = 4
  x = Vector{Float64}(undef, v * l)
  for k in 1:v
    rng = (k-1)*l+1:k*l
    x[rng] = xblocks[k]
  end
  return x
end

"""
    solve_block_tridiag!(Ablocks, Bblocks, Cblocks, bblocks)

Rozwiązuje blokowo trójdiagonalny układ A x = b metodą Gaussa **bez pivotowania**.

  Argumenty:
  - `Ablocks`  - wektor długości v, gdzie `Ablocks[k]` = A_k (macierz l×l),
  - `Bblocks`  - wektor długości v-1, gdzie `Bblocks[k]` = B_{k+1},
  - `Cblocks`  - wektor długości v-1, gdzie `Cblocks[k]` = C_k,
  - `bblocks`  - wektor długości v, gdzie `bblocks[k]` = b_k (wektor l).

Zwraca:
- wektor blokowy `xblocks` (długości v), gdzie `xblocks[k]` to rozwiązanie x_k (wektor l).
"""
function solve_block_tridiag!(Ablocks, Bblocks, Cblocks, bblocks)
  v = length(Ablocks)
  @assert length(Bblocks) == v - 1
  @assert length(Cblocks) == v - 1
  @assert length(bblocks) == v

  # -------------------------------------------------
  # 1) Eliminacja "w przód" (forward elimination)
  # -------------------------------------------------
  for k in 1:(v-1)
    # L_{k+1} = B_{k+1} * inv(A_k)
    #  tu B_{k+1} = Bblocks[k],  A_k = Ablocks[k],  C_k = Cblocks[k],  itp.

    L = Bblocks[k] * inv(Ablocks[k])  # B_{k+1} * A_k^-1

    # A_{k+1} := A_{k+1} - L * C_k
    Ablocks[k+1] .-= L * Cblocks[k]

    # b_{k+1} := b_{k+1} - L * b_k
    bblocks[k+1] .-= L * bblocks[k]

    # W praktycznej implementacji można:
    # - przechowywać odwrócony A_k tylko raz w zmiennej, jeśli l jest duże
    # - lub stosować np. factor(Ablocks[k]) = LU rozkład, żeby nie robić inv() "wprost".
  end

  # -------------------------------------------------
  # 2) Podstawienie "wstecz" (back substitution)
  # -------------------------------------------------
  xblocks = Vector{Vector{Float64}}(undef, v)

  # x_v = A_v^-1 * b_v
  xblocks[v] = inv(Ablocks[v]) * bblocks[v]

  # x_k = A_k^-1 * ( b_k - C_k * x_{k+1} ), idąc od k=v-1 w dół
  for k in (v-1):-1:1
    tmp = bblocks[k] .- Cblocks[k] * xblocks[k+1]
    xblocks[k] = inv(Ablocks[k]) * tmp
  end

  return xblocks
end

"""
    solve_block_tridiag_pivot!(Ablocks, Bblocks, Cblocks, bblocks)

Rozwiązuje blokowo-trójdiagonalny układ A*x = b 
metodą Gaussa z częściowym pivotowaniem *w obrębie bloków diagonalnych*.

Argumenty:
- `Ablocks`  - wektor długości v,  Ablocks[k] = A_k   (macierz l×l),
- `Bblocks`  - wektor długości v-1, Bblocks[k] = B_{k+1},
- `Cblocks`  - wektor długości v-1, Cblocks[k] = C_k,
- `bblocks`  - wektor długości v,   bblocks[k] = b_k (wektor l).

Zwraca:
- `xblocks` – wektor długości v,   xblocks[k] to rozwiązanie x_k (wektor l).

UWAGA: 
Pivotowanie odbywa się TYLKO w wierszach wewnątrz aktualnego bloku A_k
(+ te same wiersze w B_{k+1} i b_k). 
To poprawia stabilność względem całkowicie naiwnego Gaussa,
ale NIE jest to klasyczny partial pivot w całym n×n.
"""
function solve_block_tridiag_pivot!(Ablocks, Bblocks, Cblocks, bblocks)
  v = length(Ablocks)
  l = size(Ablocks[1], 1)
  @assert all(size(Ablocks[k], 1) == l && size(Ablocks[k], 2) == l for k in 1:v)
  @assert length(Bblocks) == v - 1
  @assert length(Cblocks) == v - 1
  @assert length(bblocks) == v

  # Eliminacja w przód (blokowa)
  for k in 1:(v-1)
    # =============================================================
    #  (a) "częściowe pivotowanie" w obrębie bloku A_k (l×l):
    #      - dla każdej kolumny col od r do l-1, 
    #        szukamy wiersza p (r..l) z max |A_k[p,col]| 
    #        i zamieniamy wiersze p <-> r w (A_k, B_{k+1}, b_k).
    #
    #      Minimalna wersja "pivotu" - bierzemy TYLKO wiersz r
    #      i szukamy p, który daje max w kolumnie col=r. 
    #      (czyli pivot w stylu "1 kolumna -> 1 wiersz" w mikroskali).
    #
    for r in 1:(l-1)
      # szukaj p w zakresie r..l z max |A_k[p, r]|
      subrange = r:l
      p_local = argmax(abs.(Ablocks[k][subrange, r]))  # np. 1..(l-r+1)
      p_global = p_local + (r - 1)  # bo subrange zaczyna się od r

      if p_global != r
        # zamiana wierszy r <-> p_global w A_k
        Ablocks[k][[r, p_global], :] = Ablocks[k][[p_global, r], :]
        # zamiana wierszy r <-> p_global w B_{k+1}
        Bblocks[k][[r, p_global], :] = Bblocks[k][[p_global, r], :]
        # zamiana elementów r <-> p_global w b_k
        bblocks[k][[r, p_global]] = bblocks[k][[p_global, r]]
      end

      # Teraz "zerowanie" poniżej r-tego wiersza w A_k
      pivot = Ablocks[k][r, r]
      @assert abs(pivot) > 1e-14 "Zerowy pivot w bloku A_k. Metoda się wysypie."
      for i in r+1:l
        m = Ablocks[k][i, r] / pivot
        # Zeruj Ablocks[k][i, r], i.e. kolumnę r
        Ablocks[k][i, r] = 0.0
        # Reszta kolumn w A_k
        for cc in r+1:l
          Ablocks[k][i, cc] -= m * Ablocks[k][r, cc]
        end
        # To samo w B_{k+1}
        for cc in 1:size(Bblocks[k], 2)
          Bblocks[k][i, cc] -= m * Bblocks[k][r, cc]
        end
        # Wejścia w wektorze b_k
        bblocks[k][i] -= m * bblocks[k][r]
      end
    end

    # =============================================================
    #  (b) Skoro "wyzerowaliśmy" dolną część A_k, 
    #      możemy teraz zdefiniować L_{k+1} = B_{k+1} * (A_k)^{-1} 
    #      i zaktualizować A_{k+1}, b_{k+1} 
    #      => w klasycznym "blok Gaussie" to: 
    #         A_{k+1} := A_{k+1} - L_{k+1}*C_k
    #         b_{k+1} := b_{k+1} - L_{k+1}*b_k
    #
    #      Ale jeśli A_k jest już w postaci górnotrójkątnej (bo "wyzerowaliśmy" poniżej diag),
    #      możemy wykonać back-substitution w obrębie A_k, by "zdjąć" B_{k+1} = L_{k+1}.
    #      Dla uproszczenia: zrobimy "odwracanie" A_k (lub LU) i potem L_{k+1} = B_{k+1} * A_k^{-1}.
    #
    #  UWAGA: Po powyższej pętli (r=1..l-1) A_k jest w *górnotrójkątnej* formie,
    #         niekoniecznie diagonalnej. Najlepiej byłoby zrobić jej rozkład LU i go użyć.

    # Dla przykładu (mało wydajnie): wprost bierzemy inv(UpperTriangular(A_k))
    A_k_upper = UpperTriangular(Ablocks[k])
    invA_k = inv(A_k_upper)  # mała macierz l×l, akceptowalne

    # L_{k+1} = B_{k+1} * invA_k
    Ltemp = Bblocks[k] * invA_k

    # A_{k+1} -= L_{k+1} * C_k
    Ablocks[k+1] .-= Ltemp * Cblocks[k]

    # b_{k+1} -= L_{k+1} * b_k
    bblocks[k+1] .-= Ltemp * bblocks[k]
  end

  # ---------------------------------------------------
  #  (c) Po etapie "w przód" mamy w A_1..A_v formę 
  #      mniej/bardziej górnotrójkątną (blokowo).
  #      Wystarczy zrobić "podstawienie wstecz" blok po bloku:
  #        x_v = A_v^-1 * b_v
  #        x_{v-1} = A_{v-1}^-1 * ( b_{v-1} - C_{v-1} * x_v )
  #        ...
  # ---------------------------------------------------
  xblocks = Vector{Vector{Float64}}(undef, v)

  # x_v
  #  Uwaga: A_v też może być "górnotrójkątne". Można wziąć inv(UpperTriangular(A_v)).
  A_v_upper = UpperTriangular(Ablocks[v])
  xblocks[v] = inv(A_v_upper) * bblocks[v]

  # x_{v-1}, x_{v-2}, ...
  for k in (v-1):-1:1
    tmp = bblocks[k] .- Cblocks[k] * xblocks[k+1]
    A_k_upper = UpperTriangular(Ablocks[k])
    xblocks[k] = inv(A_k_upper) * tmp
  end

  return xblocks
end



function blockify(A::SparseMatrixCSC{Float64}, b::Vector{Float64}, block_size::Int)
  n = size(A, 1)               # całkowity rozmiar, np. 16
  v = n ÷ block_size           # liczba bloków wzdłuż przekątnej, np. 4
  l = block_size              # rozmiar pojedynczego bloku, np. 4

  # Upewnij się, że n jest faktycznie podzielne przez block_size
  @assert n == v * l "n musi być równe v*l"

  # Przygotuj miejsce na bloki
  Ablocks = Vector{Matrix{Float64}}(undef, v)
  Bblocks = Vector{Matrix{Float64}}(undef, v - 1)
  Cblocks = Vector{Matrix{Float64}}(undef, v - 1)
  bblocks = Vector{Vector{Float64}}(undef, v)

  # ----- Wypełnianie Ablocks i bblocks -----
  for k in 1:v
    rowrange = (k-1)*l+1:k*l
    # Blok diagonalny A_k
    Ablocks[k] = Matrix(A[rowrange, rowrange])
    # Odpowiadający fragment wektora b_k
    bblocks[k] = b[rowrange]
  end

  # ----- Wypełnianie Bblocks i Cblocks -----
  for k in 1:(v-1)
    # B_{k+1} = A[  (k+1)-1 * l + 1 : (k+1)*l ,  (k+1)-2 * l + 1 : (k+1)-1 * l ]
    # ale uważajmy, bo "B_{k+1}" w opisie to wiersz k+1, kolumna k
    # Wygodniej: Bblocks[k] = B_{k+1}
    rowrangeB = k*l+1:(k+1)*l
    colrangeB = (k-1)*l+1:k*l
    Bblocks[k] = Matrix(A[rowrangeB, colrangeB])

    # C_k = A[ (k)-1*l+1 : k*l ,  k*l+1 : (k+1)*l ]
    rowrangeC = (k-1)*l+1:k*l
    colrangeC = k*l+1:(k+1)*l
    Cblocks[k] = Matrix(A[rowrangeC, colrangeC])
  end

  return Ablocks, Bblocks, Cblocks, bblocks
end



end