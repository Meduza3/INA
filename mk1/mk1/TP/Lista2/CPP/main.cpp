#include <iostream>
#include <string>
#include <stdexcept>
#include "WierszTrojkataPascala.h"

using namespace std;

int main(int argc, char * argv[]) {
    if (argc < 2) {
        cerr << "Brak Argumentów" << endl;
        return 1;
    }

    try {
        int n = stoi(argv[1]);
        WierszTrojkataPascala wiersz(n);

        for(int i = 2; i < argc; ++i) {
            try {
                int m = stoi(argv[i]);
                if ( m < 0 || m > n) {
                    throw OutOfRangeException("Liczba spoza zakresu");
                }
                cout << m << " - " << wiersz.wspolczynnik(m) << endl;
                
        } catch (const invalid_argument &) {
            cout << argv[i] << " - nieprawidłowa dana" << endl;
        } catch (const OutOfRangeException &) {
            cout << argv[i] << " - liczba spoza zakresu" << endl;
        }
    }
} catch (const invalid_argument &){
    cerr << "Błędny pierwszy argument" << endl;
    return 1;
} catch (const OutOfRangeException &) {
    cerr << "Błędny pierwszy argument" << endl;
    return 1;
}

return 0;
}
