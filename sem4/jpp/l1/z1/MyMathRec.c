#include "MyMath.h"

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



struct diophantine_results extended_euclid(int a, int b) {
    if (a == 0) {
        struct diophantine_results base_case = {0, 1};
        return base_case;
    }

    struct diophantine_results temp = extended_euclid(b % a, a);
    struct diophantine_results res;

    res.x = temp.y - (b / a) * temp.x;
    res.y = temp.x;

    return res;
}