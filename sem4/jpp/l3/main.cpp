#include "DHSetup.hpp"
#include "User.hpp"

#define CHARACTERISTIC 1234567891

int main() {
    DHSetup<CHARACTERISTIC> dhSetup;
    GF_Int<CHARACTERISTIC> generator = dhSetup.getGenerator();
    GF_Int<CHARACTERISTIC> powerResult = dhSetup.power(generator, 5);

    // std::cout << "Generator: " << generator.get_value() << std::endl;
    // for(int i = 0; i < generator.get_characteristic() - 1; i++) {
    //     std::cout << "Generator^" << i << ": " << dhSetup.power(generator, i).get_value() << "\n";
    // }

    User<CHARACTERISTIC> alice(dhSetup);
    User<CHARACTERISTIC> bob(dhSetup);

    alice.setKey(bob.getPublicKey());
    bob.setKey(alice.getPublicKey());

    GF_Int<CHARACTERISTIC> message(6969420);
    auto encryptedMessage = alice.encrypt(message);
    auto decryptedMessage = bob.decrypt(encryptedMessage);

    std::cout << "Wiadomosc oryginalna: " << message.get_value() << "\n";
    std::cout << "Encrypted Message: " << encryptedMessage.get_value() << "\n";
    std::cout << "Decrypted Message: " << decryptedMessage.get_value() << "\n";

    return 0;
}