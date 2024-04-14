#include "main.h"

int main(int argc, char* argv[]) {
    if (argc < 2) {
        std::cout << "🍄:" << argv[0] << " <-insertion|-quick|-hybrid|-merge|-custom|-dualpivot>\n";
        return 1;
    }

    algorithm picked_algorithm;

    if(strcmp(argv[1], "-insertion") == 0) {
        picked_algorithm = INSERTION;
    } else if(strcmp(argv[1], "-quick") == 0) {
        picked_algorithm = QUICK;
    } else if(strcmp(argv[1], "-hybrid") == 0) {
        picked_algorithm = HYBRID;
    } else if(strcmp(argv[1], "-merge") == 0) {
        picked_algorithm = MERGE;
    } else if(strcmp(argv[1], "-custom") == 0) {
        picked_algorithm = CUSTOM;
    } else if(strcmp(argv[1], "-dualpivot") == 0) {
        picked_algorithm = DUALPIVOT;
    } else {
        printf("🍄: -insertion, -quick, -hybrid, -merge, -custom, -dualpivot");
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

    std::string filename;
    if(picked_algorithm == INSERTION) {
        results = sorts::insertionSort(listOfNumbers, metrics);
        filename = "data/insertion.txt";
    } else if (picked_algorithm == QUICK) {
        results = sorts::quickSort(listOfNumbers, 0, listOfNumbers.size() - 1, metrics);
        filename = "data/quick.txt";   
    } else if (picked_algorithm == HYBRID) {
        results = sorts::hybridSort(listOfNumbers, metrics);
        filename = "data/hybrid.txt";   
    } else if (picked_algorithm == MERGE) {
        results = sorts::mergeSort(listOfNumbers, 0, listOfNumbers.size() - 1, metrics);
        filename = "data/merge.txt";   
    } else if (picked_algorithm == DUALPIVOT) {
        results = sorts::dualPivotQuickSort(listOfNumbers, 0, listOfNumbers.size() - 1, metrics);
        filename = "data/dualpivot.txt";   
    } else if (picked_algorithm == CUSTOM) {
        results = sorts::customSort(listOfNumbers, metrics);
        filename = "data/custom.txt";   
    }

    std::ofstream outFile(filename, std::ios::app);

    std::cout << "Tablica wejsciowa ponownie:";
    sorts::printArray(copyOfList);
    std::cout << "Tablica po posortowaniu   :";
    sorts::printArray(results.arr);
    std::cout << results.metrics.comparisons << " porownan miedzy kluczami\n";
    std::cout << results.metrics.swaps << " przestawien kluczy\n";

    std::cout << sorts::isSorted(results.arr) ? "Tablica jest poprawnie posortowana :) \n" : "Tablica nie jest poprawnie posortowana :( \n";

    

    outFile << listOfNumbers.size() - 1 << "\t" << results.metrics.comparisons << "\t" << results.metrics.swaps << '\n';


}
