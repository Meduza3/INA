#ifndef DHSETUP_H
#define DHSETUP_H

    #include "Gf_int.hpp"
    #include <vector>
    #include <random>

    template<int CHARACTERISTIC>
    class DHSetup {

        private:
            GF_Int<CHARACTERISTIC> generator;

            int powerMod(int x, unsigned int y, int p) {
                //zwraca (x^y) % p
                int res = 1;
                x = x % p;
                while (y > 0) {
                    if (y & 1) {
                        res = (res*x) % p;
                    }
                    y = y >> 1;
                    x = (x*x) % p;
                }
                return res;
            }

            bool isPrime(int n) {
                if (n <= 1) return false;
                if (n <= 3) return true;
                if (n % 2 == 0 || n % 3 == 0) return false;
                for (int i = 5; i * i <= n; i += 6)
                    if (n % i == 0 || n % (i+2) == 0)
                        return false;
                return true;
            }

            std::vector<int> primeFactors(int n) {
                std::vector<int> factors;
                while (n % 2 == 0) {
                    factors.push_back(2);
                    n = n/2;
                }
                for (int i = 3; i <= sqrt(n); i = i + 2) {
                    while (n % i == 0) {
                        factors.push_back(i);
                        n = n/i;
                    }
                }
                if (n > 2)
                    factors.push_back(n);
                return factors;
            }

            void generateGenerator() {
                std::mt19937 rng(std::random_device{}());
                std::uniform_int_distribution<> dist(2, CHARACTERISTIC - 2);
                std::vector<int> factors = primeFactors(CHARACTERISTIC - 1);

                while (true) {
                    int candidate = dist(rng);
                    bool found = true;
                    for (int factor : factors) {
                        if (powerMod(candidate, (CHARACTERISTIC - 1) / factor, CHARACTERISTIC) == 1) {
                            found = false;
                            break;
                        }
                    }
                    if (found) {
                        generator = GF_Int<CHARACTERISTIC>(candidate);
                        return;
                    }
                }
            }



        public:

        DHSetup(){
            generateGenerator();
        }

        GF_Int<CHARACTERISTIC> power(GF_Int<CHARACTERISTIC> a, unsigned long b) {
            GF_Int<CHARACTERISTIC> result(1);

            for(int i = 0; i < b; i++){
                result = result * a;
            }

            return result;
        }

        GF_Int<CHARACTERISTIC> getGenerator() {
            return generator;
        }
    };

#endif