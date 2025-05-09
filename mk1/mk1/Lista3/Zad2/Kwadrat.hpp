#ifndef KWADRAT_HPP
#define KWADRAT_HPP

#include "Czworobok.hpp"

class Kwadrat : public Czworobok {
    public:
        Kwadrat(int bok, int kat) noexcept(false);
        std::string Nazwa() noexcept(true) override;
};

#endif