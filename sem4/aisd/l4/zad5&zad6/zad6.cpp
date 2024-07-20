#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <ctime>
#include "splay.hpp"

constexpr int ITERATIONS = 20;

struct ExperimentResult {
    long long totalComparisons = 0;
    long long totalPointerOperations = 0;
    int maxHeight = 0;

    void add(const ExperimentResult& result) {
        totalComparisons += result.totalComparisons;
        totalPointerOperations += result.totalPointerOperations;
        maxHeight = std::max(maxHeight, result.maxHeight);
    }

    void divide(int factor) {
        totalComparisons /= factor;
        totalPointerOperations /= factor;
        maxHeight /= factor;
    }
};

std::vector<int> generateRandomVector(int size) {
    std::vector<int> vec(size);
    for (int i = 0; i < size; ++i) {
        vec[i] = i + 1;
    }
    std::random_shuffle(vec.begin(), vec.end());
    return vec;
}

int main() {
    std::ofstream resultsFile("results_sorted.csv");
    resultsFile << "N,Average Comparisons,Average Substitutions,Average Height,Max Comparisons,Max Substitutions,Max Height\n";

    for (int n = 10000; n <= 100000; n += 10000) {
        ExperimentResult avgResults;
        ExperimentResult maxResults;

        for (int i = 0; i < ITERATIONS; ++i) {
            SplayTree tree;
            auto elements = generateRandomVector(n);
            sort(elements.begin(), elements.end());
            ExperimentResult currentResult;

            for (int element : elements) {
                tree.insertNode(element);
            }

            currentResult.totalComparisons = tree.getKeyComparisons();
            currentResult.totalPointerOperations = tree.getPointerOperations();
            currentResult.maxHeight = tree.height();

            // Randomly delete elements
            std::random_shuffle(elements.begin(), elements.end());
            for (int element : elements) {
                tree.deleteNode(element);
            }

            if (currentResult.totalComparisons > maxResults.totalComparisons) {
                maxResults.totalComparisons = currentResult.totalComparisons;
            }
            if (currentResult.totalPointerOperations > maxResults.totalPointerOperations) {
                maxResults.totalPointerOperations = currentResult.totalPointerOperations;
            }
            if (currentResult.maxHeight > maxResults.maxHeight) {
                maxResults.maxHeight = currentResult.maxHeight;
            }

            avgResults.add(currentResult);
        }

        avgResults.divide(ITERATIONS);

        resultsFile << n << "," 
                    << avgResults.totalComparisons << ","
                    << avgResults.totalPointerOperations << ","
                    << avgResults.maxHeight << ","
                    << maxResults.totalComparisons << ","
                    << maxResults.totalPointerOperations << ","
                    << maxResults.maxHeight << "\n";
    }

    resultsFile.close();
    return 0;
}
