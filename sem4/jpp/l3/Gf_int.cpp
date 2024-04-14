
#include <iostream>
#include <ostream>

template<int CHARACTERISTIC>
class GF_Int {

    private:
        int value;
        static int modInverse(int a){
            int m0 = CHARACTERISTIC, t, q;
            int x0 = 0, x1 = 1;
            if(CHARACTERISTIC == 1) {
                return 0;
            }
            while(a > 1) {
                q = a / CHARACTERISTIC;
                t = CHARACTERISTIC;
                CHARACTERISTIC = a % CHARACTERISTIC, a = t;
                t = x0;
                x0 = x1 - q * x0;
                x1 = t;
            }
            if (x1 < 0){
                x1 += m0;
            }
            return x1;
        }


    public:
        int get_value() const {
            return value;
        }
        void set_value(int x){
            value = x % CHARACTERISTIC;
        }

        int get_characteristic() const {
            return CHARACTERISTIC;
        }



    GF_Int(int val = 0){
        set_value(val);
    }


    GF_Int& operator=(const GF_Int& rhs) {
        if (this != &rhs) {
            this->value = rhs.get_value();
        } else {
            throw std::invalid_argument("Attempt to self-assign!");
        }
        return *this;
    }

    GF_Int& operator+=(const GF_Int& rhs) {
        this->value = (this->value + rhs.get_value()) % 1234577;
        return *this;
    }

    GF_Int& operator-=(const GF_Int& rhs) {
        this->value = (this->value - rhs.get_value()) % 1234577;
        if (this->value < 0) this->value += 1234577;
        return *this;
    }

    GF_Int& operator*=(const GF_Int& rhs) {
        long long product = static_cast<long long>(this->value) * rhs.get_value() % 1234577;
        this->value = static_cast<int>(product);
        return *this;
    }

    GF_Int& operator/=(const GF_Int& rhs) {
        int rhs_val = rhs.get_value();
        if (rhs_val == 0) {
            throw std::invalid_argument("Attempt to divide by zero!");
        }
        int inverse = modInverse(rhs_val, rhs.get_characteristic());
        long long division = static_cast<long long>(this->value) * inverse % 1234577;
        this->value = static_cast<int>(division);
        return *this;
    }
    template<int M>
    friend GF_Int<M> operator*(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend GF_Int<M> operator+(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend GF_Int<M> operator-(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend GF_Int<M> operator/(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend bool operator==(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend bool operator!=(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend bool operator<=(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend bool operator>=(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend bool operator<(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

    template<int M>
    friend bool operator>(const GF_Int<M>& lhs, const GF_Int<M>& rhs);
    
    template<int M>
    friend std::ostream& operator<<(std::ostream& os, const GF_Int<M>& obj) {
        os << obj.get_value();
        return os;
    }

};

static int modInverse(int a, int mod){
    int m0 = mod, t, q;
    int x0 = 0, x1 = 1;
    if(mod == 1) {
        return 0;
    }
    while(a > 1) {
        q = a / mod;
        t = mod;
        mod = a % mod, a = t;
        t = x0;
        x0 = x1 - q * x0;
        x1 = t;
    }
    if (x1 < 0){
        x1 += m0;
    }
    return x1;
}

template<int M>
GF_Int<M> operator+(const GF_Int<M>& lhs, const GF_Int<M>& rhs) {
    return GF_Int<M>((lhs.get_value() + rhs.get_value()));
}

template<int M>
GF_Int<M> operator-(const GF_Int<M>& lhs, const GF_Int<M>& rhs) {
    int val = lhs.get_value() - rhs.get_value();
    if(val < 0){
        return GF_Int<M>(lhs.get_characteristics() + val);
    } else {
        return GF_Int<M>(val);
    }
}

template<int M>
GF_Int<M> operator/(const GF_Int<M>& lhs, const GF_Int<M>& rhs) {
    int rhs_val = rhs.get_value();
    if (rhs_val == 0) {
        throw std::invalid_argument("Attempt to divide by zero!");
    }

    int inverse = modInverse(rhs_val, rhs.get_characteristic());
    long long division = static_cast<long long>(lhs.get_value()) * inverse;
    return GF_Int<M>(division);
}

template<int M>
GF_Int<M> operator*(const GF_Int<M>& lhs, const GF_Int<M>& rhs) {
    long long product = static_cast<long long>(lhs.get_value() * rhs.get_value());
    return GF_Int<M>(product);
}

template<int M>
bool operator==(const GF_Int<M>& lhs, const GF_Int<M>& rhs){
    return lhs.get_value() == rhs.get_value();
}

template<int M>
bool operator!=(const GF_Int<M>& lhs, const GF_Int<M>& rhs){
    return lhs.get_value() != rhs.get_value();
}

template<int M>
bool operator<=(const GF_Int<M>& lhs, const GF_Int<M>& rhs){
    return lhs.get_value() <= rhs.get_value();
}

template<int M>
bool operator>=(const GF_Int<M>& lhs, const GF_Int<M>& rhs){
    return lhs.get_value() >= rhs.get_value();
}

template<int M>
bool operator>(const GF_Int<M>& lhs, const GF_Int<M>& rhs){
    return lhs.get_value() > rhs.get_value();
}

template<int M>
bool operator<(const GF_Int<M>& lhs, const GF_Int<M>& rhs){
    return lhs.get_value() < rhs.get_value();
}

int main() {
    GF_Int<1234577> p1(1234570);
    GF_Int<1234577> p2(40);

    GF_Int<1234577> a1(12);
    GF_Int<1234577> a2(1234577 + 12);
    GF_Int<1234577> a3 = a1 * a2;


    GF_Int<1234577> p3 = p1 + p2;
    std::cout << "p3: " << p3.get_value() << std::endl;
    std::cout << "a1: " << a1.get_value() << " a2: " << a2.get_value() << " " << (a1 == a2 ? "true" : "false") << std::endl;
    std::cout << a3.get_value() << std::endl;
    return 0;
}