#include "sorts.h"

void sorts::printArray(const std::vector<int>& arr) {
        std::ostream_iterator<int> out_it(std::cout, " ");
        std::copy(arr.begin(), arr.end(), out_it);
        std::cout << std::endl;
}

sorts::SortResults sorts::insertionSort(std::vector<int>& arr, sorts::SortMetrics& metrics) {
    int i, j, key;
    std::vector<int> arr_copy = arr;
    int n = arr.size();
    for (i = 1; i < n; i++) {
        key = arr[i];
        j = i - 1;

        while (j >= 0 && arr[j] > key) {
            metrics.comparisons++;
            arr[j + 1] = arr[j];
            metrics.swaps++;
            j = j - 1;
            if (n < 40 && j%2 == 1) {
                printArray(arr);
            }
        }
        arr[j + 1] = key;
        printArray(arr);
    }
    printf("Original array: ");
    printArray(arr_copy);
    return(sorts::SortResults{arr, metrics});
}

int sorts::partition(std::vector<int>& arr, int low, int high, sorts::SortMetrics& metrics) {
    int pivot = arr[high];
    int i = (low - 1);
    for (int j = low; j <= high - 1; j++) {
        metrics.comparisons++;
        if (arr[j] < pivot) {
            i++;
            std::swap(arr[i], arr[j]);
            metrics.swaps++;
        }
    }
    std::swap(arr[i + 1], arr[high]);
    metrics.swaps++;
    return (i + 1);
}

sorts::SortResults sorts::quickSort(std::vector<int>& arr, int low, int high, sorts::SortMetrics& metrics) {
    if (low < high) {
        int pi = sorts::partition(arr, low, high, metrics);
        quickSort(arr, low, pi - 1, metrics);
        quickSort(arr, pi + 1, high, metrics);
    }
    return {arr, metrics};
}

sorts::SortResults sorts::hybridSort(std::vector<int>& arr, sorts::SortMetrics& metrics) {
    if(arr.size() < 100){
        std::cout << "Using insertionSort:\n";
        insertionSort(arr, metrics);
    } else {
        std::cout << "Using quickSort:\n";
        quickSort(arr, 0, arr.size() - 1, metrics);
    }
    return {arr, metrics};
}