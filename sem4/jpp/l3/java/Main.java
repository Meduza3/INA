public class Main {
    public static void main(String[] args) {
        DHSetup dhSetup = new DHSetup(1234567891);
        System.out.println("Generator: " + dhSetup.getGenerator());

        User alice = new User(dhSetup);
        User bob = new User(dhSetup);
        GF_Int alicePublicKey = alice.getPublicKey();
        GF_Int bobPublicKey = bob.getPublicKey();
        System.out.println("Alice's public key: " + alicePublicKey);
        System.out.println("Bob's public key: " + bobPublicKey);

        alice.setKey(bobPublicKey);
        bob.setKey(alicePublicKey);

        GF_Int message = new GF_Int(4206969, dhSetup.getCharacteristic()); 
        GF_Int encryptedMessage = alice.encrypt(message); 
        GF_Int decryptedMessage = bob.decrypt(encryptedMessage); 

        System.out.println("Original Message: " + message);
        System.out.println("Encrypted Message: " + encryptedMessage);
        System.out.println("Decrypted Message: " + decryptedMessage);
    }
}