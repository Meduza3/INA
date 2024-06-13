module Binomial where


binomial :: Integer -> Integer -> Integer
binomial n k
 | k == 0 = 1
 | k == n = 1
 | k > n = 0
 | k > n - k = binomial n (n - k)
 | otherwise = binomial (n - 1) k + binomial (n - 1) (k - 1)