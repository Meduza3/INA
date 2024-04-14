template<typename T>
class DHSetup {

    private:
        GF_Int<T> field;
        T characteristic = 1234577;
        T generator = 0;

        void generateGenerator(){
            
        }

    public:

    DHSetup(){
        generateGenerator();
    }

    T power(T a, unsigned long b) {
        unsigned long result = 1;

        for(int i = 0; i < b; i++){
            result = result * a % characteristic;
        }

        return static_cast<T>(result);
    }
};


int main() {
}