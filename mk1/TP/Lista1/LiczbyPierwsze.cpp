#include "LiczbyPierwsze.hpp"
#include <vector>
#include <stdexcept>

LiczbyPierwsze::LiczbyPierwsze(int n) {
    if (n >= 2) {
        std::vector<bool> isPrime(n + 1, true);
        isPrime[0] = isPrime[1] = false;

        for (int i = 2; i * i <= n; i++) {
            if (isPrime[i]) {
                for (int j = i * i; j <= n; j += i) {
                    isPrime[j] = false;
                }
            }
        }

        for (int i = 2; i <= n; i++) {
            if (isPrime[i]) {
                liczbyPierwsze.push_back(i);
            }
        }
    }
}

int LiczbyPierwsze::liczba(int m) {
    if ( m >= 0 && m < liczbyPierwsze.size()) {
        return liczbyPierwsze[m];
    } else {
        throw std::out_of_range("liczba spoza zakresu");    }
}