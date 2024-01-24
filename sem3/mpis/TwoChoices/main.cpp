#include <iostream>
#include <random>
#include <fstream>
#include <vector>
#include <algorithm>

int main() {
    const int minN = 1000;
    const int maxN = 100000;
    const int k = 50;

    std::ofstream dataFile("data.txt");
    std::mt19937 generator(std::random_device{}());

    std::cout.rdbuf(dataFile.rdbuf());

    for(int currentN = minN; currentN <= maxN; currentN += 1000) {
        for(int d = 1; d <= 2; d++){
            
            int MaxLoadAvg = 0;
            for(int i = 0; i < k; i++){
                int MaxLoad = 0;                
                int numberOfBalls = 0;
                std::vector<int> bins(currentN, 0);
                std::uniform_int_distribution<> distribution(0, currentN - 1);

                while (true) {
                    std::vector<int> bin_indexes(d);
                    for(int j = 0; j < d; j++){
                        bin_indexes[j] = distribution(generator);
                    }
                    auto smallest = std::min_element(bin_indexes.begin(), bin_indexes.end(), 
                                                     [&](const int& a, const int& b) {
                                                         return bins[a] < bins[b];
                                                     });
                    bins[*smallest]++;
                    numberOfBalls++;
                    if(numberOfBalls == currentN){
                        MaxLoad = *std::max_element(bins.begin(), bins.end());
                        break;
                    }
                }
                MaxLoadAvg += MaxLoad;
                std::cout << currentN << "\t" << d << "\t" << MaxLoad;
                if(i == 49){
                    std::cout << "\t" << MaxLoadAvg/50;
                    MaxLoadAvg = 0;
                }
                std::cout << "\n";
            }
        }
    }

    dataFile.close();
    return 0;
}