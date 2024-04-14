
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



    GF_Int<T>(int val = 0){
        set_value(val);
    }


    GF_Int<T>& operator=(const GF_Int<T>& rhs);

    GF_Int<T>& operator+=(const GF_Int<T>& rhs);

    GF_Int<T>& operator-=(const GF_Int<T>& rhs);

    GF_Int<T>& operator*=(const GF_Int<T>& rhs);

    GF_Int<T>& operator/=(const GF_Int<T>& rhs);

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
    friend std::ostream& operator<<(std::ostream& os, const GF_Int<M>& obj);

};

static int modInverse(int a, int mod);

template<int M>
GF_Int<M> operator+(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
GF_Int<M> operator-(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
GF_Int<M> operator/(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
GF_Int<M> operator*(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
bool operator==(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
bool operator!=(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
bool operator<=(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
bool operator>=(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
bool operator>(const GF_Int<M>& lhs, const GF_Int<M>& rhs);

template<int M>
bool operator<(const GF_Int<M>& lhs, const GF_Int<M>& rhs);