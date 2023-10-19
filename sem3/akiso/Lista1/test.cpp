#include "test.h"

void test() {
    assert(GCD(50, 160) == 10);
    assert(GCD(69, 420) == 3);
    assert(GCD(282, 78) == 6);
    std::cout << "All tests passed!\n" << std::endl;
}