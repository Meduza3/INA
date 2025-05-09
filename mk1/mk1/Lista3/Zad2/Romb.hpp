#ifndef ROMB_HPP
#define ROMB_HPP

#include "Czworobok.hpp"

class Romb : public Czworobok {
    public:
        Romb(int bok, int kat) noexcept(false);
        std::string Nazwa() noexcept(true) override;
};

#endif