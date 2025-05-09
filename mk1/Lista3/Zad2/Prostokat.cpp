#include "Prostokat.hpp"

Prostokat::Prostokat(int a, int b, int c, int d, int kat) noexcept(false) {
    bok.push_back(a);
    bok.push_back(b);
    bok.push_back(c);
    bok.push_back(d);
    this->kat = kat;
}

std::string Prostokat::Nazwa() noexcept(true) {
    return "Prostokat";
}