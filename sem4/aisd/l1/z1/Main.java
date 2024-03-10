
public class Main {

    public static void main(String[] args) {
        Queue kolejka = new Queue(50);
        Stack stos = new Stack(50);

        System.out.println("Do struktur dodajemy: ");
        for(int i = 1; i <= 50; i++){
            int item = (int)(Math.random() * 50 + 1);
            System.out.print(item + " ");
            kolejka.enqueue(item);
            stos.push(item);
        }
        System.out.println();
        System.out.println("Odczytujemy teraz ze stosu: ");
        for(int i = 0; i < 50; i++){
            System.out.print(stos.pop() + " ");
        }
        System.out.println();
        System.out.println("Odczytujemy teraz z kolejki: ");
        for(int i = 0; i < 50; i++){
            System.out.print(kolejka.dequeue() + " ");
        }

        System.out.println();
    }
}
