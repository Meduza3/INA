# Marcin Zubrzycki

import Pkg
Pkg.instantiate()
Pkg.add("JuMP")
Pkg.add("GLPK")
using JuMP
using GLPK

machine_matrix = [5 10 6;
                  3 6 4;
                  4 5 3;
                  4 2 1]

availibility = [3600, 3600, 3600]
machine_costs = [2 / 60, 2 / 60, 3 / 60]
material_costs = [4, 1, 1, 1]
maximum_demand = [400, 100, 150, 500]
prices = [9, 7, 6, 5]

function solve(machine_matrix::Matrix, availibility::Vector, machine_costs::Vector, material_costs::Vector, maximum_demand::Vector, prices::Vector)
  m, n = size(machine_matrix)
  model = Model(GLPK.Optimizer)
  set_silent(model)

  @variable(model, kg_produced[1:m] >= 0, Int)

  # Nie wyprodukuj więcej niż pozwala na to dostępność
  @constraint(model, sum(repeat(kg_produced, 1, n) .* machine_matrix, dims = 1) .<= availibility)

  # Nie wyproduk więcej, niż dasz radę sprzedać
  @constraint(model, kg_produced .<= maximum_demand)

  # maksymalizowac zysk - koszta
  @objective(model, Max, sum((prices .- material_costs) .* kg_produced) - sum(machine_costs .* vec(sum(repeat(kg_produced, 1, n) .* machine_matrix, dims = 1))))
  
  optimize!(model)

  if termination_status(model) == MOI.OPTIMAL
		println("Kg na produkt na maszynę")
		display(value.(kg_produced))
		println("Minuty na produkt na maszynę")
		display(value.(kg_produced) .* machine_matrix)
		println("Profit")
		println(objective_value(model))
	elseif termination_status(model) == MOI.INFEASIBLE
		println("The model is infeasible.")
		return nothing
	else
		println("Solver did not find an optimal solution.")
		return nothing
	end
end

solve(machine_matrix, availibility, machine_costs, material_costs, maximum_demand, prices)
