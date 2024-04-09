#include <stdio.h>

extern unsigned int silnia_l(unsigned int N);
extern unsigned int silnia_r(unsigned int N);

extern unsigned int nwd_l(unsigned int A, unsigned int B);
extern unsigned int nwd_r(unsigned int A, unsigned int B);

typedef struct {
    int X;
    int Y;
    unsigned int GCD;
} Diophantine_Solution;

extern Diophantine_Solution exeuclid_l(int A, int B);
extern Diophantine_Solution exeuclid_r(int A, int B, int C); 

int main()  {
    unsigned int silnia_l_result = silnia_l(5);
    unsigned int silnia_r_result = silnia_r(10);
    printf("Silnia(5) = %u\n", silnia_l_result);
    printf("Silnia(10) = %u\n", silnia_r_result);

    int x = 24, y = 72, z = 42;
    unsigned int nwd_l_result = nwd_l(x, y);
    unsigned int nwd_r_result = nwd_r(y, z); 

    printf("nwd(%d, %d) = %u\n", x, y, nwd_l_result);
    printf("nwd(%d, %d) = %u\n", y, z, nwd_r_result);

    int c = 35,  d = 10, e = 55;
    Diophantine_Solution euclid_r_result = exeuclid_r(c, d, e);
    printf("Extended Euclid of (%d, %d):\n", c, d);
    printf("X = %d, Y = %d, GCD = %u\n", euclid_r_result.X, euclid_r_result.Y, euclid_r_result.GCD);

        int a = 252, b = 105;
    Diophantine_Solution euclid_l_result = exeuclid_l(a, b);
    printf("Extended Euclid of (%d, %d):\n", a, b);
    printf("X = %d, Y = %d, GCD = %u\n", euclid_l_result.X, euclid_l_result.Y, euclid_l_result.GCD);

}