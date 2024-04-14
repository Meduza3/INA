#include <iostream>
#include <stdexcept>

template<int CHARACTERISTIC>
class GF_Int {
private:
    int value;

    static int modInverse(int a, int m) {
        int m0 = m, t, q;
        int x0 = 0, x1 = 1;
        if (m == 1) return 0;
        while (a > 1) {
            q = a / m;
            t = m;
            m = a % m;
            a = t;
            t = x0;
            x0 = x1 - q * x0;
            x1 = t;
        }
        if (x1 < 0) x1 += m0;
        return x1;
    }

public:
    GF_Int(int val = 0) {
        value = (val % CHARACTERISTIC + CHARACTERISTIC) % CHARACTERISTIC;
    }

    int get_value() const {
        return value;
    }

    static int get_characteristic() {
        return CHARACTERISTIC;
    }

    GF_Int& operator+=(const GF_Int& rhs) {
        value = (value + rhs.value) % CHARACTERISTIC;
        return *this;
    }

    GF_Int& operator-=(const GF_Int& rhs) {
        value = (value - rhs.value + CHARACTERISTIC) % CHARACTERISTIC;
        return *this;
    }

    GF_Int& operator*=(const GF_Int& rhs) {
        long long product = static_cast<long long>(value) * rhs.value % CHARACTERISTIC;
        value = static_cast<int>(product);
        return *this;
    }

    GF_Int& operator/=(const GF_Int& rhs) {
        int rhs_val = rhs.get_value();
        if (rhs_val == 0) throw std::invalid_argument("Attempt to divide by zero!");
        int inverse = modInverse(rhs_val, CHARACTERISTIC);
        value = static_cast<long long>(value) * inverse % CHARACTERISTIC;
        return *this;
    }

    friend GF_Int operator+(GF_Int lhs, const GF_Int& rhs) { return lhs += rhs; }
    friend GF_Int operator-(GF_Int lhs, const GF_Int& rhs) { return lhs -= rhs; }
    friend GF_Int operator*(GF_Int lhs, const GF_Int& rhs) { return lhs *= rhs; }
    friend GF_Int operator/(GF_Int lhs, const GF_Int& rhs) { return lhs /= rhs; }
};

template<int M>
std::ostream& operator<<(std::ostream& os, const GF_Int<M>& obj) {
    os << obj.get_value();
    return os;
}
