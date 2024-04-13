def factorial(n):
	result = 1
	for i in range(2,n+1):
		result *= i
	return result

def gcd(a, b):
	if b == 0:
		return a
	else:
		return gcd(b, a % b)

def gcd_extended(a, b):
    if a == 0:
        return b, 0, 1
    gcd, x1, y1 = gcd_extended(b % a, a)
    x = y1 - (b // a) * x1
    y = x1
    return gcd, x, y

def diophantine(a, b, c):
    if a == 0 and b == 0:
        print("Invalid equation: both coefficients cannot be zero.")
        exit(1)
    
    gcd_ab, x, y = gcd_extended(a, b)
    
    if c % gcd_ab != 0:
        print("No solution: c is not divisible by gcd(a, b).")
        exit(1)
    
    # Scale the solutions x and y by c // gcd_ab
    x *= c // gcd_ab
    y *= c // gcd_ab
    
    # Output the solutions which include all modifications due to sign changes
    return x, y
