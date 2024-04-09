from ctypes import *

mymath = CDLL("./mymath.so")

class DiophantineSolution(Structure):
    _fields_ = [("x", c_int),
                ("y", c_int),
                ("gcd", c_uint)]

print(mymath.silnia_r(5))
print(mymath.silnia_l(6))
print(mymath.nwd_l(54, 23))
print(mymath.nwd_r(32, 76))

mymath.extended_euclid_r.restype = DiophantineSolution

solution1 = mymath.extended_euclid_r(52, 423)
print(f"{solution1.x} {solution1.y}")
solution2 = mymath.extended_euclid_r(523, 423)
print(f"{solution2.x} {solution2.y}")