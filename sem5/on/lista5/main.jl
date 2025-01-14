using SparseArrays
include("blocksys.jl")
using .blocksys

A = nothing
b = nothing


function read_matrix(filename::String)
  open(filename, "r") do f
    params = split(strip(readline(f)))
    size = parse(Int, params[1])
    block_size = parse(Int, params[2])
    A = spzeros(Float64, size, size)
    while !eof(f)
      line = split(strip(readline(f)))
      if length(line) == 3
        i = parse(Int, line[1])
        j = parse(Int, line[2])
        v = parse(Float64, line[3])
        A[i, j] = v
      else
        println("Pominięto nieprawidłowy wiersz: ", join(line, " "))
      end
    end
    return A
  end
end

function read_vector(filename::String)
  open(filename, "r") do f
    n = parse(Int, strip(readline(f)))
    b = zeros(n)
    for i in 1:n
      b[i] = parse(Float64, strip(readline(f)))
    end
    return b
  end
end


while true
  print("#\$%> ")
  command = split(readline())
  if isempty(command) || command[1] == "exit"
    println()
    break
  elseif command[1] == "read"
    if length(command) < 3
      println("Użycie: read A b")
      continue
    end
    try
      global A = read_matrix(String(command[2]))
      global b = read_vector(String(command[3]))
      println("Pomyślnie wczytano pliki")
      Ablocks, Bblocks, Cblocks, bblocks = blockify(A, b, 4)
      xblocks = solve_block_tridiag_pivot!(Ablocks, Bblocks, Cblocks, bblocks)
      x = join_blocks(xblocks)
      println("Rozwiązanie x = ", x)
    catch err
      println("Nie udało się wczytać plików: ", err)
    end
  elseif command[1] == "experiment"
    if length(command) < 3
      println("Użycie: experiment A b")
      continue
    end
    try
      global A = read_matrix(String(command[2]))
      global b = read_vector(String(command[3]))
      Ablocks, Bblocks, Cblocks, bblocks = blockify(A, b, 4)

      tridiag_time = @elapsed begin
        xblocks_tridiag = solve_block_tridiag!(Ablocks, Bblocks, Cblocks, bblocks)
        x_tridiag = join_blocks(xblocks_tridiag)
      end

      pivot_time = @elapsed begin
        xblocks_pivot = solve_block_tridiag_pivot!(Ablocks, Bblocks, Cblocks, bblocks)
        x_pivot = join_blocks(xblocks_pivot)
      end

      standard_time = @elapsed begin
        x_standard = A \ b
      end
      println("No pivot: ", tridiag_time)
      println("Pivot: ", pivot_time)
      println("Standard: ", standard_time)

    catch err
      println("Nie udało się wczytać plików: ", err)
    end
  else
    println("Użycie: read A b")
  end
end
