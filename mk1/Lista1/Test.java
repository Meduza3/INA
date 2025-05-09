
public class Test {
    public static void main(String[] args) {
        if (args.length < 1) {
            System.out.println("Nieprawidłowy zakres");
            return;
        }

        int n;
        try {
            n = Integer.parseInt(args[0]);
        } catch (NumberFormatException e) {
            System.out.println(args[0] + " - Nieprawidłowy zakres");
            return;
        }

        if (n < 2) {
            System.out.println(n + " - Nieprawidłowy zakres");
            return;
        }
            // Tworzymy obiekt klasy LiczbyPierwsze z podanym zakresem n
        LiczbyPierwsze liczbyPierwsze = new LiczbyPierwsze(n);

        for (int i = 1; i < args.length; i++){
            int m;
            try {
                m = Integer.parseInt(args[i]);
            } catch (NumberFormatException e) {
                System.out.println(args[i] + " - nieprawidłowa dana");
                continue;
            }

            try {
                System.out.println(m + " - " + liczbyPierwsze.liczba(m));
            } catch (IndexOutOfBoundsException e) {
                System.out.println(m + " - liczba spoza zakresu");
            }
        }
    }
}
