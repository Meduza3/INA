// sorts.h
#ifndef SORTS_H
#define SORTS_H

#include <vector>
#include <iostream>
#include <iterator>

namespace sorts {

    class SortMetrics {
    public:
        int comparisons = 0;
        int swaps = 0;

        void reset() {
            comparisons = 0;
            swaps = 0;
        }
    };

    struct SortResults {
        public:
            std::vector<int> arr;
            sorts::SortMetrics metrics;
    };

    SortResults insertionSort(std::vector<int>& arr, SortMetrics& metrics);
    SortResults quickSort(std::vector<int>& arr, int low, int high, sorts::SortMetrics& metrics);
    SortResults hybridSort(std::vector<int>& arr, SortMetrics& metrics);
    int partition(std::vector<int>& arr, int low, int high, SortMetrics& metrics);
    void printArray(const std::vector<int>& arr); 
}

#endif // SORTS_H
