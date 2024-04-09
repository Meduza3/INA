#include "MyMath.h"

int silnia_r(int n){
    if (n > 1) {
        return n * silnia_r(n - 1);
    } else {
        return 1;
    }
}

int nwd_r(int a, int b) {
    return (a%b == 0) ? b : nwd_r(b, a%b);
}



struct diophantine_results extended_euclid_r(int a, int b) {
    if (a == 0) {
        struct diophantine_results base_case = {0, 1};
        return base_case;
    }

    struct diophantine_results temp = extended_euclid_r(b % a, a);
    struct diophantine_results res;

    res.x = temp.y - (b / a) * temp.x;
    res.y = temp.x;

    return res;
}