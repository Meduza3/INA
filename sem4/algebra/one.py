
# N(a + bj) = a**2 + b**2
def N(number):
    return number.real ** 2 + number.imag ** 2
# x | y jesli jest takie z, ze y = x * z

def DivRem(x, y):
    q = x/y
    q_real = round(q.real)
    q_imag = round(q.imag)

    q = complex(q_real, q_imag)

    r = x - q * y
    return q, r

def NWD(u, v):
    while N(v) != 0:
        q, _ = DivRem(u, v)
        u, v = v, u - q * v
    return u

def NWW(u, v):
    nwd = NWD(u, v)
    nww = (u * v) / nwd
    return complex(round(nww.real), round(nww.imag))

c = NWD(3 + 4j, 1 + 3j)
d = NWW(3 + 4j, 1 + 3j)

print(c)
print(d)