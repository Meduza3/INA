# Marcin Zubrzycki

using JuMP
Pkg.add("JuMP")
using GLPK

function solve(fuel_cost_matrix::Matrix, max_production::Vector, demand::Vector)
  m, n = size(fuel_cost_matrix)
  model = Model(GLPK.Optimizer)
  set_silent(model)
end