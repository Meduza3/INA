#ifndef PROSTOKAT_HPP
#define PROSTOKAT_HPP

#include "Czworobok.hpp"

class Prostokat : public Czworobok {
    public:
        Prostokat(int a, int b, int c, int d, int kat) noexcept(false);
        std::string Nazwa() noexcept(true) override;
};

#endif