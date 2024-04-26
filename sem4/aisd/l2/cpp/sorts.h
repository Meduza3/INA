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

    SortResults insertionSort(std::vector<int>& arr, SortMetrics& metrics, int n);
    SortResults quickSort(std::vector<int>& arr, int low, int high, sorts::SortMetrics& metrics, int n);
    SortResults hybridSort(std::vector<int>& arr, SortMetrics& metrics, int n);
    SortResults mergeSort(std::vector<int>& arr, int l, int r, sorts::SortMetrics, int n);
    SortResults dualPivotQuickSort(std::vector<int>& arr, int low, int high, sorts::SortMetrics& metrics, int n);
    SortResults customSort(std::vector<int>& arr, SortMetrics& metrics, int n);
    void findRuns(std::vector<int>& arr, std::vector<int>& runs);
    void mergeRuns(std::vector<int>& arr, int start, int mid, int end, sorts::SortMetrics& metrics, int n);
    void adaptiveMergeSort(std::vector<int>& arr, std::vector<int>& runs, int low, int high, sorts::SortMetrics& metrics, int n);
    void merge(std::vector<int>& arr, int l, int m, int r, sorts::SortMetrics& metrics, int n);
    int partition(std::vector<int>& arr, int low, int high, SortMetrics& metrics);
    void dualPivotPartition(std::vector<int>& arr, int low, int high, int& lp, int& rp, sorts::SortMetrics& metrics, int n);
    void printArray(const std::vector<int>& arr);
    bool isSorted(const std::vector<int>& vec);
}

#endif // SORTS_H
