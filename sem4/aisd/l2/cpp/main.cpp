#include "main.h"

int main(int argc, char* argv[]) {
    if (argc < 2) {
        std::cout << "🍄:" << argv[0] << " -insertion|-quick|-hybrid\n";
        return 1;
    }

    algorithm picked_algorithm;

    if(strcmp(argv[1], "-insertion") == 0) {
        picked_algorithm = INSERTION;
    } else if(strcmp(argv[1], "-quick") == 0) {
        picked_algorithm = QUICK;
    } else if(strcmp(argv[1], "-hybrid") == 0) {
        picked_algorithm = HYBRID;
    } else {
        printf("🍄: -insertion, -quick or -hybrid");
    }

    std::vector<int> listOfNumbers;
    std::string inputLine;

    std::getline(std::cin, inputLine);

    std::stringstream ss(inputLine);

    int num;
    while (ss >> num) {
        listOfNumbers.push_back(num);
        if (ss.peek() == ',') ss.ignore();
    }

    std::vector<int> copyOfList = listOfNumbers;
    sorts::SortMetrics metrics{0, 0};
    sorts::SortResults results{listOfNumbers, metrics};
    std::cout << "Tablica wejsciowa:";
    sorts::printArray(listOfNumbers);

    
    if(picked_algorithm == INSERTION){
        results = sorts::insertionSort(listOfNumbers, metrics);    
    } else if (picked_algorithm == QUICK){
        results = sorts::quickSort(listOfNumbers, 0, listOfNumbers.size() - 1, metrics);   
    } else {
        results = sorts::hybridSort(listOfNumbers, metrics);   
    }

    std::cout << "Tablica wejsciowa ponownie:";
    sorts::printArray(copyOfList);
    std::cout << "Tablica po posortowaniu   :";
    sorts::printArray(results.arr);
    std::cout << results.metrics.comparisons << " porownan miedzy kluczami\n";
    std::cout << results.metrics.swaps << " przestawien kluczy\n";
}