def DivRem(a, b):
    b_deg = len(b) - 1
    q = [0] * (len(a) - len(b) + 1)
    r = a[:]

    while len(r) >= len(b):
        # Aktualizacja stopnia r przy każdej iteracji
        r_deg = len(r) - 1
        wspolczynnik = r[r_deg] / b[b_deg]  # Dostosowanie indeksu dla r
        q[r_deg - b_deg] = wspolczynnik  # Dostosowanie indeksu dla q

        for j in range(b_deg, -1, -1):
            if r_deg - b_deg + j < len(r):
                r[r_deg - b_deg + j] -= wspolczynnik * b[j]

        # Usuwanie zer na końcu listy r
        while len(r) > 0 and r[-1] == 0:
            r.pop()

    return q, r


def NWD(a, b):
    while any(b):
        _, reszta = DivRem(a, b)
        a, b = b, reszta
    
    nwd = [x / a[-1] for x in a]
    return nwd


def MultPol(a, b):
    wynik = [0] * (len(a) + len(b) -1)
    for i in range(len(a)):
        for j in range(len(b)):
            wynik[i + j] += a[i] * b[j]
    return wynik

def NWW(a, b):
    nwd = NWD(a, b)
    ab = MultPol(a, b)
    nww, _ = DivRem(ab, nwd)
    nww = [x / nww[-1] for x in nww]
    return nww


a = [1, 0, 1]
b = [1, 2, 1]

c = NWD(a, b)
d = NWW(a, b)
print(c)
print(d)