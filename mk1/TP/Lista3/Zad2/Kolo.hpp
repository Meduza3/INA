#ifndef KOLO_HPP
#define KOLO_HPP

#include "Figura.hpp"

class Kolo : public Figura {
    private:
        double promien;

    public:
    Kolo(double promien) noexcept(true);
    double Pole() noexcept(true) override;
    double Obwod() noexcept(true) override;
    std::string Nazwa() noexcept(true) override;
};

#endif