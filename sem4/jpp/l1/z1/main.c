#include "MyMath.h"
#include <stdio.h>

int main() {
    printf("%d\n", silnia(6));
    printf("NWD 43423 i 123123 %d\n", nwd(43424, 123122));
    struct diophantine_results res = extended_euclid(20, 43);
    printf("Results of diophantine equation with 20 and 43: x = %d, y = %d\n", res.x, res.y);

}