#include "Kolo.hpp"
#include "Szesciokat.hpp"
#include "Pieciokat.hpp"
#include "Czworobok.hpp"
#include "Prostokat.hpp"
#include "Romb.hpp"
#include "Figura.hpp"
#include "Kwadrat.hpp"

#include <iostream>
#include <cstring>
#include <stdexcept>

int main(int argc, const char* const argv[]) {
    if (argc < 3) {
        std::cout << "Oczekiwałem więcej (3) argumentów, niż " << argc << '\n';
        return 0;
    }

    Figura *Figura;

    if (std::strcmp(argv[1], "c") == 0) {
        if (argc != 7) {
            std::cout << "Oczekiwałem 7 argumentów, nie " << argc << '\n';
            return 0;
        }
        
        double bok[4];
        double kat;

        if (bok[0] == bok[1] && bok[0] == bok[2] && bok[0] == bok[3]) {
            if (angle == 90) Figura = new Kwadrat(bok[0], kat); //Kwadrat
            else { //Romb
                try {
                    Figura = new Romb(bok[0], kat);
                }
                catch (const std::invalid_argument &e) {
                    std::cout << "To nie jest romb ani żaden czworokąt jaki znam." << '\n';
                    return 0;
                }
            }
        } 
        else { //Prostokąt
            try {
                Figura = new Prostokat(bok[0], bok[1], bok[2], bok[3], kat);
            }
            catch (const std::invalid_argument &e) {
                std::cout << "To nie jest Prostokat ani żaden czworokąt jaki znam.\n";
                return 0;
            }
        }


    }
    else if (std::strcmp(argv[1], "o") == 0 || std::strcmp(argv[1], "p") == 0 || std::strcmp(argv[1], "s") == 0) { //Koło, pięciokąt, sześciokąt
        if (argc != 3) {
            std::cout << "Spodziewałem się trzech argumentów, a nie " << argc << '\n';
            return 0;
        }

        double n;

        if (std::strcmp(argv[1], "o") == 0) { //Koło
            Figura = new Kolo(n);
        }
        else if (std::strcmp(argv[1], "p") == 0) { //Pięciokąt
            Figura = new Pieciokat(n);
        }
        else if (std::strcmp(argv[1], "s") == 0) { //Sześciokąt
            Figura = new Szesciokat(n);
        }            
    } 
    else {
        std::cout << "Spodziewałem się c, o, p lub s jako pierwszy argument." << '\n';
        return 0;
    } 

    std::cout << "Nazwa: " << Figura->Nazwa() << '\n';
    std::cout << "Pole: " << Figura->Pole() << '\n';
    std::cout << "Obwod: " << Figura->Obwod() << '\n';

    delete Figura;

    return 0;
}