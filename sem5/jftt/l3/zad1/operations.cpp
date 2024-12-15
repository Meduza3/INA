#include "operations.hpp"

long long Add(long long a, long long b, long long characteristic) {
    return (a + b) % characteristic;
}

long long Sub(long long a, long long b, long long characteristic) {
    return (a - b + characteristic) % characteristic;
}

long long Mul(long long a, long long b, long long characteristic) {
    return (a * b) % characteristic;
}

long long Div(long long a, long long b, long long characteristic) {
    return (a * Pow(b, characteristic - 2, characteristic)) % characteristic;
}

long long Pow(long long a, long long b, long long characteristic) {
    long long result = 1;
    while (b > 0) {
        if (b % 2 == 1) {
            result = Mul(result, a, characteristic);
        }
        a = Mul(a, a, characteristic);
        b /= 2;
    }
    return result;
}