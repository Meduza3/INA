#include "WierszTrojkataPascala.h"

WierszTrojkataPascala::WierszTrojkataPascala(int n) {
    wiersz.resize(n+1);
    wiersz[0] = 1;

    for(int i = 1; i <= n; ++i) {
        wiersz[i] = static_cast<long long>(wiersz[i - 1]) * (n - i + 1) / i;
    }
}

WierszTrojkataPascala::~WierszTrojkataPascala() {
}

long long WierszTrojkataPascala::wspolczynnik(int m) {
    if (m < 0 || m >= static_cast<int>(wiersz.size())) {
        throw OutOfRangeException("Liczba spoza zakresu");
    }
    return wiersz[m];
}