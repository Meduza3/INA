#ifndef HEXAGON_HPP
#define HEXAGON_HPP

#include "Figura.hpp"

class Szesciokat : public Figura {
    private:
        double bok;

    public:
        Szesciokat(int side) noexcept(true);
        double Pole() noexcept(true) override;
        double Obwod() noexcept(true) override;
        std::string Nazwa() noexcept(true) override;
};

#endif