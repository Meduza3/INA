#include "biblioteka.h"

int silnia(int n){
    if (n > 1) {
        return n * silnia(n - 1);
    } else {
        return 1;
    }
}

int nwd(int a, int b) {
    return (a%b == 0) ? b : nwd(b, a%b);
}

struct dio_result extended_gcd(int a, int b) {
    if (b == 0) {
        return (struct dio_result){1, 0, a};
    } else {
        struct dio_result next = extended_gcd(b, a % b);
        return (struct dio_result){next.y, next.x - (a / b) * next.y, next.gcd};
    }
}

struct dio_result diofantyczne(int a, int b, int c) {
    struct dio_result result = extended_gcd(a, b);
    int g = result.gcd;

    if (c % g != 0) {
        printf("No solution\n");
        return (struct dio_result){0, 0, g};
    }

    int x0 = result.x * (c / g);
    int y0 = result.y * (c / g);

    return (struct dio_result){x0, y0, g};
}
