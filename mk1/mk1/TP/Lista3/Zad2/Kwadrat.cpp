#include "Kwadrat.hpp"

Kwadrat::Kwadrat(int a, int kat) noexcept(false) {

    for (int i = 0; i < 4; i++) bok.push_back(a);
    this->bok = bok;
}

std::string Kwadrat::Nazwa() noexcept(true) {
    return "Kwadrat";
}