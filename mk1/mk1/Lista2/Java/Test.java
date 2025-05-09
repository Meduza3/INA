public class Test {
        public static void main(String[] args) {
            if(args.length == 0){
                System.out.println("Brak argumentów");
                return;
            }

            try{
                int n = Integer.parseInt(args[0]);
                if(n < 0){
                    throw new OutOfRangeException("Liczba spoza zakresu");
                }
                WierszTrojkataPascala wiersz = new WierszTrojkataPascala(n);

                for (int i = 1; i < args.length; i++) {
                    try {
                        int m = Integer.parseInt(args[i]);
                        System.out.printf("%d - %d\n", m , wiersz.wspolczynnik(m));
                    } catch (NumberFormatException e){
                        System.out.printf("%s - nieprawidłowa dana \n", args[i]);
                    } catch (OutOfRangeException e){
                        System.out.printf("%s - liczba spoza zakresu\n", args[i]);
                    }
                }
            } catch (NumberFormatException e) {
                System.out.println("Błąd: Pierwszy argument musi być liczbą całkowitą");
        } catch (OutOfRangeException e) {
            System.out.println("Błąd: Pierwszy argument musi być liczbą nieujemną");
        }
    }
}