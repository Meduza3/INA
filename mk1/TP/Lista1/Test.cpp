#include <iostream>
#include <cstdlib>
#include "LiczbyPierwsze.hpp"

int main(int argc, char *argv[]) {
    if (argc < 2) {
        std::cout << "Nieprawidłowy zakres" << std::endl;
        return 1;
    }

    int n;
    try {
        n = std::stoi(argv[1]);
    } catch (const std::invalid_argument &) {
        std::cout << argv[1] << " - Nieprawidłowy zakres" << std::endl;
        return 1;
    }

    if (n < 2) {
        std::cout << n << " - Nieprawidłowy zakres" << std::endl;
        return 1;
    }

    LiczbyPierwsze liczbyPierwsze(n);

    for (int i = 2; i < argc; i++) {
        int m;
        try {
            m = std::stoi(argv[i]);
        } catch (const std::invalid_argument &) {
            std::cout << argv[i] << " - nieprawidłowa dana" << std::endl;
            continue;
        }

        try {
            std::cout << liczbyPierwsze.liczba(m) << std::endl;
        } catch (const std::out_of_range &) {
            std::cout << m << " - liczba spoza zakresu" << std::endl;
        }
    }
    return 0;
}
