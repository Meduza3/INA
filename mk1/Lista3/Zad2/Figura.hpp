#ifndef FIGURA_HPP
#define FIGURA_HPP

#include <string>

class Shape {
    public:
        virtual double Pole() noexcept(true) = 0;
        virtual double Obwod() noexcept(true) = 0;
        virtual std::string Nazwa() noexcept(true) = 0;
};

#endif