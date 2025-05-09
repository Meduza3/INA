import java.util.*;

// Klasa LiczbyPierwsze generuje listę liczb pierwszych do określonej liczby n
public class LiczbyPierwsze {

    private List<Integer> liczbyPierwsze; // Lista przechowująca liczby pierwsze

    public LiczbyPierwsze(int n) {
        liczbyPierwsze = new ArrayList<>();

        if( n >= 2 ){
            boolean[] isPrime = new boolean[n + 1];
            Arrays.fill(isPrime, true); // Wypełniamy tablicę wartościami true
            isPrime[0] = isPrime[1] = false; // 0 i 1 nie są liczbami pierwszymi
            // Sito Eratostenesa
            for (int i = 2; i * i <= n; i++) {
                if (isPrime[i]) {
                    for (int j = i * i; j <=n; j += i) {
                        isPrime[j] = false;
                    }
                }
            }
            
            // Dodajemy liczby pierwsze do listy
            for (int i = 2; i <= n; i++) {
                if (isPrime[i]) {
                    liczbyPierwsze.add(i);
                }
            }
        }
    }

    public int liczba(int m) {
        if (m >= 0 && m < liczbyPierwsze.size()) {
            return liczbyPierwsze.get(m);
        } else {
            throw new IndexOutOfBoundsException("liczba spoza zakresu");
        }
    }
}
