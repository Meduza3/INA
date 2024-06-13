(defun binomial2 (n k)
  (if (< n k)
    nil
      (let ((row (list 1)))
        (dotimes (i n)
          (setq row (cons 1 (mapcar #'+ row (append (list 0) row)))))
        (nth k row))))