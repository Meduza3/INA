#include <iostream>
#include <random>
#include <fstream>
#include <vector>
#include <algorithm>

void InsertionSort(std::vector<int>& vector, int& wykonanePorownania, int& przestawioneKlucze) {
    for(int j = 1; j <= vector.size(); j++){
        wykonanePorownania++;
        int key = vector[j];
        int i = j - 1;
        while(i >= 0 && vector[i] >= key){
            wykonanePorownania++;
            vector[i + 1] = vector[i];
            przestawioneKlucze++;
            i--;
        }
        vector[i + 1] = key;
        przestawioneKlucze++;
    }
}

int main() {

    std::mt19937 generator(std::random_device{}());
    std::ofstream dataFile("data_insertion.txt");
    std::cout.rdbuf(dataFile.rdbuf());

    int minN = 100;
    int maxN = 10000;

    for(int currentN = minN; currentN <= maxN; currentN += 100){
        int wykPorAvg = 0;
        int przeKluAvg = 0;
        for(int k = 0; k < 50; k++){
            int wykonanePorownania = 0;
            int przestawioneKlucze = 0;
            std::vector<int> A(currentN);
            std::iota(A.begin(), A.end(), 0);
            std::shuffle(A.begin(), A.end(), generator);
            InsertionSort(A, wykonanePorownania, przestawioneKlucze);
            std::cout << currentN << " " << wykonanePorownania << " " << przestawioneKlucze;
            
            wykPorAvg += wykonanePorownania;
            przeKluAvg += przestawioneKlucze;

            if(k == 49){
                std::cout << " " << wykPorAvg/50 << " " << przeKluAvg/50;
                wykPorAvg = 0;
                przeKluAvg = 0;
            }
            std::cout << std::endl;
            }
    }
}

