#ifndef CZWOROBOK_HPP
#define CZWOROBOK_HPP

#include "Figura.hpp"

#include <vector>

class Czworobok : public Figura {
    protected:
        std::vector<int> bok;
        int kat;

    public:
        double Pole() noexcept(true) override;
        double Obwod() noexcept(true) override;
};

#endif