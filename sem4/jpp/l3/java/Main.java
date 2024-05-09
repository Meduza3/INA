public class Main {
    public static void main(String[] args) {
        DHSetup dhSetup = new DHSetup(1234567891);
        System.out.println("Generator: " + dhSetup.getGenerator());
        System.out.print("Is generator actually a generator?: ");
        System.out.println(dhSetup.isGenerator(dhSetup.getGenerator().getValue(), 1234567891));
        User alice = new User(dhSetup);
        User bob = new User(dhSetup);
        System.out.println("Alice secret: " + alice.getSecret());
        System.out.println("Bob secret: " + bob.getSecret());
        GF_Int alicePublicKey = bob.getPublicKey();
        GF_Int bobPublicKey = alice.getPublicKey();
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