import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class DHSetup {
    private long CHARACTERISTIC;
    private GF_Int generator;
    private Random random = new Random();


    public long getCharacteristic() {
        return CHARACTERISTIC;
    }

    public DHSetup(long characteristic) {
        this.CHARACTERISTIC = characteristic;
        generateGenerator();
    }

    private long powerMod(long x, long y, long p) {
        if (y == 0) {
            return x;
        }

        if (x == 0) {
            return y;
        }
        long res = 1;
        x = x % p;
        if (x < 0) x += p;
        while (y > 0) {
            if ((y & 1) == 1) {
                res = (res * x) % p;
            }
            y = y >> 1;
            x = (x * x) % p;
            if (x < 0) x += p;
        }
        return res;
    }

    public boolean isGenerator(long g, long p) {
        List<Long> factors = primeFactors(p - 1);
        for (long factor : factors) {
            if (powerMod(g, (p - 1) / factor, p) == 1) {
                return false;
            }
        }
        return true;
    }

    private List<Long> primeFactors(long n) {
        List<Long> factors = new ArrayList<>();
        while (n % 2 == 0) {
            factors.add(2L);
            n = n / 2;
        }
        for (long i = 3; i <= Math.sqrt(n); i = i + 2) {
            while (n % i == 0) {
                factors.add(i);
                n = n / i;
            }
        }
        if (n > 2)
            factors.add(n);
        return factors;
    }

    private void generateGenerator() {
        List<Long> factors = primeFactors(CHARACTERISTIC - 1);

        while (true) {
            long candidate = 2L + (long) (Math.random() * (CHARACTERISTIC - 2));
            boolean found = true;
            for (long factor : factors) {
                if (powerMod(candidate, (CHARACTERISTIC - 1) / factor, CHARACTERISTIC) == 1) {
                    found = false;
                    break;
                }
            }
            if (found) {
                generator = new GF_Int(candidate, CHARACTERISTIC);
                return;
            }
        }
    }

    public GF_Int power(GF_Int a, long b) {
        GF_Int result = new GF_Int(1, CHARACTERISTIC);
        GF_Int base = a;
    
        while (b > 0) {
            if (b % 2 != 0) {
                result = result.multiply(base);
                result = result.mod();
            }
            base = base.multiply(base);
            b /= 2;
        }
    
        return result;
    }

    public GF_Int getGenerator() {
        return generator;
    }

}
