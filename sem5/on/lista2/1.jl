# Marcin Zubrzycki

function algo1(x, y)
  Sum = BigFloat(0)
  for i in 1:5
    Sum += x[i]*y[i]
  end
  return Sum
end

function algo2(x, y)
  Sum = BigFloat(0)
  for i in 5:-1:1
    Sum += x[i]*y[i]
  end
  return Sum
end

function algo3(x, y)
  products = x .* y
  positive_products = sort(filter(z -> z > 0, products), rev = true)
  negative_products = sort(filter(z -> z < 0, products))
  sum_pos = sum(positive_products)
  sum_neg = sum(negative_products)
  return sum_pos + sum_neg
end

function algo4(x, y)
  products = x .* y
  positive_products = sort(filter(z -> z > 0, products))
  negative_products = sort(filter(z -> z < 0, products), rev = true)
  sum_pos = sum(positive_products)
  sum_neg = sum(negative_products)
  return sum_pos + sum_neg
end

true_value = -1.00657107000000e-11


x_f32 = Float32[2.718281828, -3.141592654, 1.414213562, 0.5772156649, 0.3010299957]
y_f32 = Float32[1486.2497, 878366.9879, -22.37492, 4773714.647, 0.000185049]
x2_f32 = Float32[2.718281828, -3.141592654, 1.414213562, 0.577215664, 0.301029995]

x_f64 = Float64[2.718281828, -3.141592654, 1.414213562, 0.5772156649, 0.3010299957]
y_f64 = Float64[1486.2497, 878366.9879, -22.37492, 4773714.647, 0.000185049]
x2_f64 = Float64[2.718281828, -3.141592654, 1.414213562, 0.577215664, 0.301029995]

x_big = BigFloat[2.718281828, -3.141592654, 1.414213562, 0.5772156649, 0.3010299957]
y_big = BigFloat[1486.2497, 878366.9879, -22.37492, 4773714.647, 0.000185049]
x2_big = BigFloat[2.718281828, -3.141592654, 1.414213562, 0.577215664, 0.301029995]

algorithms = [algo1, algo2, algo3, algo4]
algorithm_names = ["Algorithm 1", "Algorithm 2", "Algorithm 3", "Algorithm 4"]

precisions = [
    (x_f32, x2_f32, y_f32, "Float32"),
    (x_f64, x2_f64, y_f64, "Float64"),
    (x_big, x2_big, y_big, "BigFloat")
]

for (x, x2, y, typename) in precisions
  println("\nUsing precision: $typename")
  for (i, algo) in enumerate(algorithms)
      sum1 = algo(x, y)
      sum2 = algo(x2, y)
      difference = sum1 - sum2
      println("$(algorithm_names[i]):")
      println("  Result with x:  $sum1")
      println("  Result with x2: $sum2")
      println("  Difference:     $difference\n")
  end
end