
#include <iostream>
#include <ostream>

#define CHARACTERISTIC 1234577

class GF_Int {
    //Cialo skonczone o charakterystyce 1234577

    private:
        int value;
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
        int inverse = modInverse(rhs_val, 1234577);
        long long division = static_cast<long long>(this->value) * inverse % 1234577;
        this->value = static_cast<int>(division);
        return *this;
    }

    friend GF_Int operator*(const GF_Int& lhs, const GF_Int& rhs);
    friend GF_Int operator+(const GF_Int& lhs, const GF_Int& rhs);
    friend GF_Int operator-(const GF_Int& lhs, const GF_Int& rhs);
    friend GF_Int operator/(const GF_Int& lhs, const GF_Int& rhs);

    friend bool operator==(const GF_Int& lhs, const GF_Int& rhs);
    friend bool operator!=(const GF_Int& lhs, const GF_Int& rhs);
    friend bool operator<=(const GF_Int& lhs, const GF_Int& rhs);
    friend bool operator>=(const GF_Int& lhs, const GF_Int& rhs);
    friend bool operator<(const GF_Int& lhs, const GF_Int& rhs);
    friend bool operator>(const GF_Int& lhs, const GF_Int& rhs);
    

    friend std::ostream& operator<<(std::ostream& os, const GF_Int& obj) {
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


GF_Int operator+(const GF_Int& lhs, const GF_Int& rhs) {
    return GF_Int((lhs.get_value() + rhs.get_value()));
}

GF_Int operator-(const GF_Int& lhs, const GF_Int& rhs) {
    int val = lhs.get_value() - rhs.get_value();
    if(val < 0){
        return GF_Int(CHARACTERISTIC + val);
    } else {
        return GF_Int(val);
    }
}

GF_Int operator/(const GF_Int& lhs, const GF_Int& rhs) {
    int rhs_val = rhs.get_value();
    if (rhs_val == 0) {
        throw std::invalid_argument("Attempt to divide by zero!");
    }

    int inverse = modInverse(rhs_val, rhs.get_characteristic());
    long long division = static_cast<long long>(lhs.get_value()) * inverse;
    return GF_Int(division);
}

GF_Int operator*(const GF_Int& lhs, const GF_Int& rhs) {
    long long product = static_cast<long long>(lhs.get_value() * rhs.get_value());
    return GF_Int(product);
}

bool operator==(const GF_Int& lhs, const GF_Int& rhs){
    return lhs.get_value() == rhs.get_value();
}

bool operator!=(const GF_Int& lhs, const GF_Int& rhs){
    return lhs.get_value() != rhs.get_value();
}

bool operator<=(const GF_Int& lhs, const GF_Int& rhs){
    return lhs.get_value() <= rhs.get_value();
}

bool operator>=(const GF_Int& lhs, const GF_Int& rhs){
    return lhs.get_value() >= rhs.get_value();
}

bool operator>(const GF_Int& lhs, const GF_Int& rhs){
    return lhs.get_value() > rhs.get_value();
}

bool operator<(const GF_Int& lhs, const GF_Int& rhs){
    return lhs.get_value() < rhs.get_value();
}

int main() {
    GF_Int p1(1234570);
    GF_Int p2(40);

    GF_Int a1(12);
    GF_Int a2(1234577 + 12);
    GF_Int a3 = a1 * a2;


    GF_Int p3 = p1 + p2;
    std::cout << "p3: " << p3.get_value() << std::endl;
    std::cout << "a1: " << a1.get_value() << " a2: " << a2.get_value() << " " << (a1 == a2 ? "true" : "false") << std::endl;
    std::cout << a3.get_value() << std::endl;
    return 0;
}