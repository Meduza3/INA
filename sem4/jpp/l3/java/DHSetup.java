import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class DHSetup {
    private int CHARACTERISTIC;
    private GF_Int generator;
    private Random random = new Random();


    public int getCharacteristic() {
        return CHARACTERISTIC;
    }

    public DHSetup(int characteristic) {
        this.CHARACTERISTIC = characteristic;
        generateGenerator();
    }

    private int powerMod(int x, int y, int p) {
        int res = 1;
        x = x % p;
        while (y > 0) {
            if ((y & 1) == 1) {
                res = (res * x) % p;
            }
            y = y >> 1;
            x = (x * x) % p;
        }
        return res;
    }

    private boolean isPrime(int n) {
        if (n <= 1) return false;
        if (n <= 3) return true;
        if (n % 2 == 0 || n % 3 == 0) return false;
        for (int i = 5; i * i <= n; i += 6)
            if (n % i == 0 || n % (i + 2) == 0)
                return false;
        return true;
    }

    private List<Integer> primeFactors(int n) {
        List<Integer> factors = new ArrayList<>();
        while (n % 2 == 0) {
            factors.add(2);
            n = n / 2;
        }
        for (int i = 3; i <= Math.sqrt(n); i = i + 2) {
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
        List<Integer> factors = primeFactors(CHARACTERISTIC - 1);

        while (true) {
            int candidate = 2 + random.nextInt(CHARACTERISTIC - 2);
            boolean found = true;
            for (int factor : factors) {
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
