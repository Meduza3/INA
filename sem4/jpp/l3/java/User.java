import java.util.Random;

public class User {
    private DHSetup dhSetup;
    private GF_Int secret;
    private GF_Int publicKey;
    private GF_Int encryptionKey;

    private Random rng = new Random();

    public User(DHSetup setup) {
        this.dhSetup = setup;
        generateSecret();
    }

    public GF_Int getSecret() {
        return secret;
    }

    private void generateSecret() {
        long secretValue = 1 + rng.nextInt((int) (dhSetup.getCharacteristic() - 2));
        secret = new GF_Int(secretValue, dhSetup.getCharacteristic());
        publicKey = dhSetup.power(dhSetup.getGenerator(), secret.getValue());
    }

    public GF_Int getPublicKey() {
        return publicKey;
    }

    public void setKey(GF_Int a) {
        encryptionKey = dhSetup.power(a, secret.getValue());
    }

    public GF_Int encrypt(GF_Int m) {
        //return m.multiply(encryptionKey);
        return dhSetup.power(m, encryptionKey.getValue());
    }

    public GF_Int decrypt(GF_Int c) {
        GF_Int inverseKey = encryptionKey.inverse();
        return inverseKey.multiply(c);
    }
}