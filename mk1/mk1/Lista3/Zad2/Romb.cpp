#include "Romb.hpp"

Romb::Romb(int a, int kat) noexcept(false) {
    for (int i = 0; i < 4; i++) bok.push_back(a);
    this->kat = kat;
}

std::string Romb::Nazwa() noexcept(true) {
    return "Romb";
}