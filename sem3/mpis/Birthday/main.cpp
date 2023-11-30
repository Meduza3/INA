#include <iostream>
#include <random>
#include <vector>
#include <fstream>

bool checkCn(std::vector<int> &urny){
    int n = urny.size();
    bool res;
    for(int i = 1; i < n; i++){
       if(urny[i] == 0){
            return 0;
       }
    }
    return 1;
}

bool checkDn(std::vector<int> &urny){
    int n = urny.size();
    bool res;
    for(int i = 1; i < n; i++){
        if(urny[i] < 2){
            return 0;
        }
    }
    return 1;
}

int main() {
    std::ofstream outfile("data/data99.txt");

        for(int n = 99000; n <= 1000000; n += 1000){
                unsigned int B_n_avg = 0;
                unsigned int U_n_avg = 0;
                unsigned int C_n_avg = 0;
                unsigned int D_n_avg = 0;
                unsigned int DC_n_avg = 0;

            for(int k = 0; k < 50; k++){
                std::vector<int> urny(n);
                unsigned int B_n = 0;
                unsigned int U_n = 0;
                unsigned int C_n = 0;
                unsigned int D_n = 0;


                std::mt19937 mt(std::random_device{}());

                std::uniform_int_distribution<int> distribution(1, n);

                for(int i = 0; ; i++){
                    int current_urna = distribution(mt);

                    if(B_n == 0 && urny[current_urna] != 0){
                        B_n = i + 1;
                    }

                    urny[current_urna]++;

                    if(checkCn(urny) && C_n == 0){
                        C_n = i + 1;
                    }

                    if(i == n){
                        for(int i = 1; i <= n; i++){
                            if(urny[i] == 0){
                                U_n++;
                            }
                        }
                    }

                    if(checkDn(urny) && D_n == 0){
                        D_n = i + 1;
                        break;
                    }
                }

                std::cout.rdbuf(outfile.rdbuf());
                std::cout << n << "\t";
                std::cout << B_n << "\t";
                B_n_avg += B_n;
                std::cout << U_n << "\t";
                U_n_avg += U_n;
                std::cout << C_n << "\t";
                C_n_avg += C_n;
                std::cout << D_n << "\t";
                D_n_avg += D_n;
                std::cout << D_n - C_n;
                DC_n_avg += (D_n - C_n);
                if(k == 49){
                    std::cout << "\t";
                    std::cout << B_n_avg/50 << "\t";
                    std::cout << U_n_avg/50 << "\t";
                    std::cout << C_n_avg/50 << "\t";
                    std::cout << D_n_avg/50 << "\t";
                    std::cout << DC_n_avg/50;
                }
                std::cout << std::endl;

            }
        }
        return 0;
}