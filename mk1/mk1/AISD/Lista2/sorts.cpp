#include "sorts.h"

void sorts::printArray(const std::vector<int>& arr) {
        std::ostream_iterator<int> out_it(std::cout, " ");
        std::copy(arr.begin(), arr.end(), out_it);
        std::cout << std::endl;
}

sorts::SortResults sorts::insertionSort(std::vector<int>& arr, sorts::SortMetrics& metrics, int n) {
    int i, j, key;
    std::vector<int> arr_copy = arr;
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

sorts::SortResults sorts::quickSort(std::vector<int>& arr, int low, int high, sorts::SortMetrics& metrics, int n) {
    if (low < high) {
        int pi = sorts::partition(arr, low, high, metrics);
        quickSort(arr, low, pi - 1, metrics, n);
        quickSort(arr, pi + 1, high, metrics, n);
        if(n <= 40) {
        printArray(arr);
        }
    }
    return {arr, metrics};
}

void sorts::merge(std::vector<int>& arr, int l, int m, int r, sorts::SortMetrics& metrics, int n) {
    int i, j, k;
    int n1 = m - l + 1;
    int n2 = r - m;

    std::vector<int> L(n1), R(n2);

    for (i = 0; i < n1; i++)
        L[i] = arr[l + i];
    for (j = 0; j < n2; j++)
        R[j] = arr[m + 1 + j];

    // Merge the temp arrays back into arr[l..r]
    i = 0; // Initial index of first subarray
    j = 0; // Initial index of second subarray
    k = l; // Initial index of merged subarray
    while (i < n1 && j < n2) {
        metrics.comparisons++;
        if (L[i] <= R[j]) {
            arr[k] = L[i];
            i++;
        } else {
            arr[k] = R[j];
            j++;
        }
        k++;
    }

    while (i < n1) {
        metrics.swaps++;
        arr[k] = L[i];
        i++;
        k++;
    }

    while (j < n2) {
        metrics.swaps++;
        arr[k] = R[j];
        j++;
        k++;
    }

    if(n <= 40) {
        printArray(arr);
    }
}

sorts::SortResults sorts::mergeSort(std::vector<int>& arr, int l, int r, sorts::SortMetrics metrics, int n) {
    if (l < r) {
        int m = l + (r - l) / 2;

        mergeSort(arr, l, m, metrics, n);
        mergeSort(arr, m + 1, r, metrics, n);

        merge(arr, l, m, r, metrics, n);
    }
    return {arr, metrics};  
}

sorts::SortResults sorts::hybridSort(std::vector<int>& arr, sorts::SortMetrics& metrics, int n) {
    if(arr.size() < 17){
        std::cout << "Using insertionSort:\n";
        insertionSort(arr, metrics, n);
    } else {
        std::cout << "Using quickSort:\n";
        quickSort(arr, 0, arr.size() - 1, metrics, n);
    }
    return {arr, metrics};
}

void sorts::dualPivotPartition(std::vector<int>& arr, int low, int high, int& lp, int& rp, sorts::SortMetrics& metrics, int n) {
    if (arr[low] > arr[high])
        std::swap(arr[low], arr[high]);

    int j = low + 1;
    int g = high - 1, k = low + 1;
    int p = arr[low], q = arr[high];

    while (k <= g) {
        // Compare with smaller pivot
        if (arr[k] < p) {
            metrics.comparisons++;
            std::swap(arr[k], arr[j]);
            metrics.swaps++;
            j++;
        } else if (arr[k] >= q) {  // Compare with larger pivot
            metrics.comparisons++;
            while (arr[g] > q && k < g) {
                metrics.comparisons++;
                g--;
            }
            std::swap(arr[k], arr[g]);
            metrics.swaps++;
            g--;
            if (arr[k] < p) {
                std::swap(arr[k], arr[j]);
                metrics.swaps++;
                j++;
            }
        }
        k++;
    }
    j--;
    g++;

    // Swap pivots to final positions
    std::swap(arr[low], arr[j]);
    std::swap(arr[high], arr[g]);
    metrics.swaps += 2;

    lp = j;  // Update left pivot index
    rp = g;  // Update right pivot index
}

sorts::SortResults sorts::dualPivotQuickSort(std::vector<int>& arr, int low, int high, sorts::SortMetrics& metrics, int n) {
    if (low < high) {
        int lp, rp;
        dualPivotPartition(arr, low, high, lp, rp, metrics, n);
        dualPivotQuickSort(arr, low, lp - 1, metrics, n);
        dualPivotQuickSort(arr, lp + 1, rp - 1, metrics, n);
        dualPivotQuickSort(arr, rp + 1, high, metrics, n);

        if( n <= 40) {
            printArray(arr);
        }
    }
    return {arr, metrics};
}

void sorts::findRuns(std::vector<int>& arr, std::vector<int>& runs) {
    int n = arr.size();
    int i = 0;
    while (i < n) {
        int start = i;
        while (i < n - 1 && arr[i] <= arr[i + 1]) {
            i++;
        }
        runs.push_back(start);  // Store start of run
        runs.push_back(i);      // Store end of run
        i++;
    }
}

void sorts::mergeRuns(std::vector<int>& arr, int start, int mid, int end, sorts::SortMetrics& metrics, int n) {
    std::vector<int> temp(end - start + 1);
    int i = start, j = mid + 1, k = 0;

    while (i <= mid && j <= end) {
        metrics.comparisons++;
        if (arr[i] <= arr[j]) {
            temp[k++] = arr[i++];
        } else {
            temp[k++] = arr[j++];
        }
    }

    while (i <= mid) {
        metrics.swaps++;
        temp[k++] = arr[i++];
    }

    while (j <= end) {
        metrics.swaps++;
        temp[k++] = arr[j++];
    }

    for (i = start, k = 0; i <= end; i++, k++) {
        arr[i] = temp[k];
    }

    if( n <= 40) {
        printArray(arr);
    }
}

void sorts::adaptiveMergeSort(std::vector<int>& arr, std::vector<int>& runs, int low, int high, sorts::SortMetrics& metrics, int n) {
    if (low < high) {
        int mid = low + (high - low) / 2;
        adaptiveMergeSort(arr, runs, low, mid, metrics, n);
        adaptiveMergeSort(arr, runs, mid + 1, high, metrics, n);
        mergeRuns(arr, runs[low], runs[mid], runs[high], metrics, n);
    }
}

sorts::SortResults sorts::customSort(std::vector<int>& arr, sorts::SortMetrics& metrics, int n) {
    std::vector<int> runs;
    findRuns(arr, runs);
    adaptiveMergeSort(arr, runs, 0, runs.size() - 1, metrics, n);
    return {arr, metrics};
}

bool sorts::isSorted(const std::vector<int>& vec) {
    for (size_t i = 1; i < vec.size(); i++) {
        if (vec[i-1] > vec[i]) {
            return false;
        }
        return true;
    }
    return true;
}