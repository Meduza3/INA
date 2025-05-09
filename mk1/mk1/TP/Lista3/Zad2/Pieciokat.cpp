#include "Pieciokat.hpp"

#include <cmath>

Pieciokat::Pieciokat(int bok) noexcept(true) {
    this->bok = bok;
}

double Pieciokat::Pole() noexcept(true) {
    return sqrt(25 + 10 * sqrt(5)) * ((bok * bok) / 4);
}

double Pieciokat::Obwod() noexcept(true) {
    return 5 * bok;
}

std::string Pieciokat::Nazwa() noexcept(true) {
    return "Pieciokat";
}