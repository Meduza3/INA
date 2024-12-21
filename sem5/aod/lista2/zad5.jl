# Marcin Zubrzycki
using JuMP
using GLPK


min_cars_matrix = [
                   2 4 3;
                   3 6 5;
                   5 7 8]

max_cars_matrix = [
                   3 5 7;
                   5 7 10;
                   8 12 10]

min_car_per_shift = [10, 20, 18]
min_cars_per_district = [10, 14, 13]

function solve(min_cars_matrix::Matrix, max_cars_matrix::Matrix, min_car_per_shift::Vector, min_cars_per_district::Vector)

	m, n = size(min_cars_matrix)
	model = Model(GLPK.Optimizer)
	set_silent(model)

    # Liczba radiowozów
	@variable(model, 0 <= cars_used[1:m, 1:n]) #, Int)

	@constraint(model, min_cars_matrix .<= cars_used .<= max_cars_matrix)
	@constraint(model, vec(sum(cars_used, dims = 1)) .>= min_car_per_shift)
	@constraint(model, vec(sum(cars_used, dims = 2)) .>= min_cars_per_district)

    # Jak najmniej radiowozów
	@objective(model, Min, sum(cars_used))

	optimize!(model)

	if termination_status(model) == MOI.OPTIMAL
		println("Rozmieszczenie radiowozów")
		display(value.(cars_used))
		println("Liczba radiowozów")
		println(objective_value(model))
	elseif termination_status(model) == MOI.INFEASIBLE
		println("The model is infeasible.")
		return nothing
	else
		println("Solver did not find an optimal solution.")
		return nothing
	end
end

solve(min_cars_matrix, max_cars_matrix, min_car_per_shift, min_cars_per_district)