public class GF1234577 {
    private int value;
    private int characteristic = 1234577;

    private int modInverse(int a, int mod){
        int m0 = mod, t, q;
        int x0 = 0, x1 = 1;
        if(mod == 1) {
            return 0;
        }
        while(a > 1) {
            q = a / mod;
            t = mod;
            mod = a % mod;
            a = t;
            t = x0;
            x0 = x1 - q * x0;
            x1 = t;
        }
        if (x1 < 0){
            x1 += m0;
        }
        return x1;
    }

    public int getValue() {
        return this.value;
    }

    public void setValue(int value) {
        this.value = value % characteristic;
    }

    public int getCharacteristic() {
        return this.characteristic;
    }

    GF1234577(int value) {
        setValue(value);
    }

    GF1234577() {
        setValue(0);
    }

    public GF1234577 add(GF1234577 rhs) {
        return new GF1234577(value + rhs.getValue());
    }

    public GF1234577 sub(GF1234577 rhs) {
        int res = value - rhs.getValue();
        if (res < 0 ) {
            return new GF1234577(characteristic + res);
        } else {
            return new GF1234577(res);
        }
    }

    public GF1234577 mult(GF1234577 rhs) {
        return new GF1234577(value * rhs.getValue());
    }

    public GF1234577 div(GF1234577 rhs) {
        if(rhs.getValue() == 0){
            throw new IllegalArgumentException("Can't divide by zero!");
        }

        int inverse = modInverse(rhs.getValue(), rhs.getCharacteristic());
        return new GF1234577(value * inverse);
    }

    public Boolean eq(GF1234577 rhs) {
        return value == rhs.getValue();
    }

    @Override
    public String toString() {
        return Integer.toString(value);
    }

    public static void main(String[] args) {
        GF1234577 a = new GF1234577(4591);
        GF1234577 b = new GF1234577(1435);
        GF1234577 c = new GF1234577(5925);
        GF1234577 d = new GF1234577(14854);

        System.out.println("a + b = " + a.add(b));
        System.out.println("a - d = " + a.sub(d));
        System.out.println("d * c = " + d.mult(c));
        System.out.println("c / b = " + c.sub(b));
        System.out.println("is a equal to b*d? " + (a.eq(b.mult(d)) ? "Yes" : "No"));
        
    }
}
