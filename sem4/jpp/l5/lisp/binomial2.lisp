(defun pascals-triangle-row (n)
    (if (= n 0)
        '(1)
        (let ((previous-row (pascals-triangle-row (- n 1))))
        (cons 1 (append (mapcar #'+ previous-row (cdr previous-row)) '(1))))))

(defun binomial2 (n k)
    (nth k (pascals-triangle-row n)))