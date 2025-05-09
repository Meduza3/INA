#include "Kolo.hpp"

#include <cmath>

Kolo::Kolo(int promien) noexcept(true){
    this->promien = promien;
}

double Kolo::Pole() noexcept(true) {
    return M_PI * promien * promien;
}

double Kolo::Obwod() noexcept(true) {
    return 2 * M_PI * promien;
}

std::string Kolo::Nazwa() noexcept(true) {
    return "Koło";
}