#include "DHSetup.h"

template<int CHARACTERISTIC>
class DHSetup {

    private:
        GF_Int<CHARACTERISTIC> generator;

        void generateGenerator(){
            generator = GF_Int<CHARACTERISTIC>(3);
        }

    public:

    DHSetup(){
        generateGenerator();
    }

    GF_Int<CHARACTERISTIC> power(GF_Int<CHARACTERISTIC> a, unsigned long b) {
        GF_Int<CHARACTERISTIC> result(1);

        for(int i = 0; i < b; i++){
            result = result * a;
        }

        return result;
    }

    GF_Int<CHARACTERISTIC> getGenerator() {
        return generator;
    }
};


int main() {
    DHSetup<1234577> dhSetup;
    GF_Int<1234577> generator = dhSetup.getGenerator();
    GF_Int<1234577> powerResult = dhSetup.power(generator, 5);

    std::cout << "Generator: " << generator.get_value() << std::endl;
    std::cout << "Generator^5: " << powerResult.get_value() << std::endl;

    return 0;
}