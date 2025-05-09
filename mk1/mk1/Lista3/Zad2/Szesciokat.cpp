#include "Szesciokat.hpp"

#include <cmath>

Szesciokat::Szesciokat(int bok) noexcept(true) {
    this->bok = bok;
}

double Szesciokat::Pole() noexcept(true) {
    return (3 * sqrt(3) * bok * bok) / 2;
}

double Szesciokat::Obwod() noexcept(true) {
    return 6 * bok;
}

std::string Szesciokat::Nazwa() noexcept(true) {
    return "Szesciokat";
}