#include "DHSetup.hpp"
#include <random>

template<int CHARACTERISTIC>
class User {
    private:
        DHSetup<CHARACTERISTIC>& dhSetup;
        GF_Int<CHARACTERISTIC> secret;
        GF_Int<CHARACTERISTIC> publicKey;
        GF_Int<CHARACTERISTIC> encryptionKey;

        std::mt19937 rng{std::random_device{}()};

        void generateSecret() {    
            std::uniform_int_distribution<> dist(1, CHARACTERISTIC - 2);
            secret = GF_Int<CHARACTERISTIC>(dist(rng));
            publicKey = dhSetup.power(dhSetup.getGenerator(), secret.get_value());
        }

    public:
        User(DHSetup<CHARACTERISTIC>& setup) : dhSetup(setup) {
            generateSecret();
        }

        GF_Int<CHARACTERISTIC> getPublicKey() const {
            return publicKey;
        }

        void setKey(GF_Int<CHARACTERISTIC> a) {
            encryptionKey = dhSetup.power(a, secret.get_value());
        }

        GF_Int<CHARACTERISTIC> encrypt(GF_Int<CHARACTERISTIC> m) {
            return m * encryptionKey;
        }

        GF_Int<CHARACTERISTIC> decrypt(GF_Int<CHARACTERISTIC> c) {
                return c / encryptionKey;
        }
};