public class GF_Int {
    private long value;
    private long CHARACTERISTIC;

    public GF_Int(long value, long characteristic) {
        this.CHARACTERISTIC = characteristic;
        this.value = (value % CHARACTERISTIC + CHARACTERISTIC) % CHARACTERISTIC;
    }

    public long getValue() {
        return value;
    }

    public GF_Int mod() {
        long newValue = value % CHARACTERISTIC;
        if (newValue < 0) newValue += CHARACTERISTIC;
        return new GF_Int(newValue, CHARACTERISTIC);
    }

    public GF_Int inverse() {
        long t = 0, newt = 1;
        long r = CHARACTERISTIC, newr = value;

        while (newr != 0) {
            long quotient = r / newr;

            long temp = newt;
            newt = t - quotient * newt;
            t = temp;

            temp = newr;
            newr = r - quotient * newr;
            r = temp;
        }

        if (r > 1) {
            throw new ArithmeticException("Inverse does not exist");
        }
        if (t < 0) {
            t = t + CHARACTERISTIC;
        }

        return new GF_Int(t, CHARACTERISTIC);
    }

    public long getCharacteristic() {
        return CHARACTERISTIC;
    }

    public GF_Int add(GF_Int rhs) {
        long newValue = (value + rhs.value) % CHARACTERISTIC;
        return new GF_Int(newValue, CHARACTERISTIC);
    }

    public GF_Int subtract(GF_Int rhs) {
        long newValue = (value - rhs.value + CHARACTERISTIC) % CHARACTERISTIC;
        return new GF_Int(newValue, CHARACTERISTIC);
    }

    public GF_Int multiply(GF_Int rhs) {
        long product = (long) value * rhs.value % CHARACTERISTIC;
        return new GF_Int((long) product, CHARACTERISTIC);
    }

    public GF_Int divide(GF_Int rhs) {
        if (rhs.value == 0) throw new IllegalArgumentException("Attempt to divide by zero!");
        long inverse = modInverse(rhs.value, CHARACTERISTIC);
        long newValue = (long) value * inverse % CHARACTERISTIC;
        return new GF_Int((long) newValue, CHARACTERISTIC);
    }

    private static long modInverse(long a, long m) {
        long m0 = m, t, q;
        long x0 = 0, x1 = 1;
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
    public static void main(String[] args) {
        GF_Int a = new GF_Int(3,11);
        GF_Int b = a.inverse();
        System.err.println(b);

        DHSetup setup = new DHSetup(a.getCharacteristic());
        setup.power(a, 3);
        System.err.println(a);
    }
}
