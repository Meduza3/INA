# Marcin Zubrzycki


import Pkg
Pkg.instantiate()
Pkg.add("JuMP")
Pkg.add("GLPK")
using JuMP
using GLPK

function solve(fuel_cost_matrix::Matrix, max_production::Vector, demand::Vector)
  m, n = size(fuel_cost_matrix)
  model = Model(GLPK.Optimizer)
  set_silent(model)

  @variable(model, solution_matrix[1:m, 1:n] >= 0, Int)

	@constraint(model, vec(sum(solution_matrix, dims = 1)) .<= max_production)

  @constraint(model, vec(sum(solution_matrix, dims = 2)) .== demand)

  @objective(model, Min, sum(fuel_cost_matrix .* solution_matrix))

  optimize!(model)

  if termination_status(model) == MOI.OPTIMAL
		println("Wysyłane paliwo")
		display(value.(solution_matrix))
		println("Suma paliwa dla firm")
		println(vec(sum(value.(solution_matrix), dims = 1)))
		println("Koszt")
		println(objective_value(model))
	elseif termination_status(model) == MOI.INFEASIBLE
		println("The model is infeasible.")
		return nothing
	else
		println("Solver did not find an optimal solution.")
		return nothing
	end


end

fuel_cost_matrix = [10 7 8;
                    10 11 14;
                    9 12 4;
                    11 13 9]

max_production = [275000, 550000, 660000]
demand = [110000, 220000, 330000, 440000]

solve(fuel_cost_matrix, max_production, demand)