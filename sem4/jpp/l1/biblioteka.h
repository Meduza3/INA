#ifndef BIBLIOTEKA_H
#define BIBLIOTEKA_H

int silnia(int n);
int nwd(int a, int b);
struct dio_result diofantyczne(int a, int b, int c);
struct dio_result extended_gcd(int a, int b);

struct dio_result
{
    int x;
    int y;
    int gcd;
};

#endif
