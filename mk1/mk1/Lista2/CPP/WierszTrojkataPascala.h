#ifndef WIERSZTROJKATAPASCALA_H
#define WIERSZTROJKATAPASCALA_H

#include <vector>
#include <stdexcept>
#include <string>

using namespace std;

class OutOfRangeException : public std::runtime_error {
    public:
    explicit OutOfRangeException(const string &message) : runtime_error(message) {}  
};

class WierszTrojkataPascala {
    public:
    WierszTrojkataPascala(int n);
    ~WierszTrojkataPascala();
    long long wspolczynnik(int m);

    private:
    vector<int> wiersz;
};

#endif