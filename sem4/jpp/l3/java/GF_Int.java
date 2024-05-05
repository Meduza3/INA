public class GF_Int {
    private int value;
    private int CHARACTERISTIC;

    public GF_Int(int value, int characteristic) {
        this.CHARACTERISTIC = characteristic;
        this.value = (value % CHARACTERISTIC + CHARACTERISTIC) % CHARACTERISTIC;
    }

    public int getValue() {
        return value;
    }

    public int getCharacteristic() {
        return CHARACTERISTIC;
    }

    public GF_Int add(GF_Int rhs) {
        int newValue = (value + rhs.value) % CHARACTERISTIC;
        return new GF_Int(newValue, CHARACTERISTIC);
    }

    public GF_Int subtract(GF_Int rhs) {
        int newValue = (value - rhs.value + CHARACTERISTIC) % CHARACTERISTIC;
        return new GF_Int(newValue, CHARACTERISTIC);
    }

    public GF_Int multiply(GF_Int rhs) {
        long product = (long) value * rhs.value % CHARACTERISTIC;
        return new GF_Int((int) product, CHARACTERISTIC);
    }

    public GF_Int divide(GF_Int rhs) {
        if (rhs.value == 0) throw new IllegalArgumentException("Attempt to divide by zero!");
        int inverse = modInverse(rhs.value, CHARACTERISTIC);
        long newValue = (long) value * inverse % CHARACTERISTIC;
        return new GF_Int((int) newValue, CHARACTERISTIC);
    }

    private static int modInverse(int a, int m) {
        int m0 = m, t, q;
        int x0 = 0, x1 = 1;
        if (m == 1) return 0;
        while (a > 1) {
            q = a / m;
            t = m;
            m = a % m;
            a = t;
            t = x0;
            x0 = x1 - q * x0;
            x1 = t;
        }
        if (x1 < 0) x1 += m0;
        return x1;
    }

    @Override
    public String toString() {
        return String.valueOf(value);
    }
}