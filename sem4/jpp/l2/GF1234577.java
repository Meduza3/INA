public class GF1234577 {
    private int value;
    private int characteristic = 1234577;

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
            return GF1234577(characteristic + res);
        } else {
            return GF1234577(res);
        }
    }

    public GF1234577 mult(GF1234577 rhs) {
        return GF1234577(value * rhs.getValue());
    }

    public GF1234577 div(GF1234577 rhs) {
        if(rhs.getValue() == 0){
            throw IllegalArgumentException("Can't divide by zero!");
        }

        int inverse = modInverse(rhs.getValue(), rhs.getCharacteristic());

        
    }

    @Override
    public String toString() {
        return Integer.parseInt(value);
    }
}
