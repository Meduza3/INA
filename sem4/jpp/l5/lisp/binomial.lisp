(defun binomial (n k)
  (if (or (= k 0) (= n k))
  1
  (+ (binomial (- n 1) k) (binomial (- n 1) (- k 1)))))