class OutOfRangeException extends Exception {
    public OutOfRangeException(String message) {
        super(message);
    }
}

class InvalidInputException extends Exception {
    public InvalidInputException(String message) {
        super(message);
    }
}

public class WierszTrojkataPascala {
    private int[] wiersz;
        public WierszTrojkataPascala(int n) {
            wiersz = new int[n+1];
            wiersz[0] = 1;

            for(int i = 1; i <= n; i++){
                wiersz[i] = (int) ((long) wiersz[i-1] * (n- i + 1) /1);
            }
        }

        public int wspolczynnik(int m) throws OutOfRangeException {
            if (m < 0 || m >= wiersz.length) {
                throw new OutOfRangeException("Liczba spoza zakresu");
            }
            return wiersz[m];
        }
}
