#include <iostream>
#include <random>
#include <vector>
#include <fstream>

bool checkCn(std::vector<int> urny){
    int n = urny.size();
    bool res;
    for(int i = 1; i < n; i++){
       if(urny[i] == 0){
            return 0;
       }
    }
    return 1;
}

bool checkDn(std::vector<int> urny){
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

    std::ofstream outfile_B("data/data_b.txt");
    std::ofstream outfile_U("data/data_u.txt");
    std::ofstream outfile_C("data/data_c.txt");
    std::ofstream outfile_D("data/data_d.txt");
    std::ofstream outfile_DC("data/data_dc.txt");
    for(int k = 0; k < 10; k++){
        for(int n = 1000; n <= 10000; n += 1000){
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
            std::cout.rdbuf(outfile_B.rdbuf());
            std::cout << B_n << std::endl;
            std::cout.rdbuf(outfile_U.rdbuf());
            std::cout << U_n << std::endl;
            std::cout.rdbuf(outfile_C.rdbuf());
            std::cout << C_n << std::endl;
            std::cout.rdbuf(outfile_D.rdbuf());
            std::cout << D_n << std::endl;
            std::cout.rdbuf(outfile_DC.rdbuf());
            std::cout << D_n - C_n << std::endl;

        }
        
        return 0;
    }
}