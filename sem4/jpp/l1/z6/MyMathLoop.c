#include "MyMath.h"

int silnia_l(int a){
    int res = 1;
    for(int i = 2; i <= a; i++){
        res *= i;
    }
    return res;
}

int nwd_l(int a, int b) {
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

struct diophantine_results extended_euclid_l(int a, int b) {
    int x0 = 1, y0 = 0; // Represents x, y for the current step
    int x1 = 0, y1 = 1; // Represents x, y for the next step
    int q, temp;

    while (b != 0) {
        q = a / b; // Quotient

        // Calculate next x and y
        temp = x1;
        x1 = x0 - q * x1;
        x0 = temp;

        temp = y1;
        y1 = y0 - q * y1;
        y0 = temp;

        // Apply the Euclid's algorithm step
        temp = a % b;
        a = b;
        b = temp;
    }

    struct diophantine_results result = {x0, y0}; // x0 and y0 now hold the result
    return result;
}