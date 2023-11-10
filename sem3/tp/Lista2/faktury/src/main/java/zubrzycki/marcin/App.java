package zubrzycki.marcin;

import java.util.Scanner;
import java.util.ArrayList;
public class App 
{
    public static void main( String[] args )
    {
        ArrayList<Faktura> faktury = new ArrayList<Faktura>();
        int wybor2;
        System.out.println("Witaj w Zubrzycki Faktury!");
        while (true) {
            
            System.out.println("1. Wyświetl listę faktur, 2. Dodaj fakturę, 3. Usuń fakturę, 4. Edytuj fakturę, 5. Bazy danych, 6. Wyjdź");
            Scanner scanner = new Scanner(System.in);
            int wybor = scanner.nextInt();
            switch (wybor) {
                case 1:
                    for(Faktura faktura : faktury) {
                        faktura.wyswietl();
                    }
                    break;
                
                case 2:
                    Faktura faktura = new Faktura();
                    do{
                        System.out.println("Podaj nazwę pozycji:");
                        String nazwa = scanner.next();
                        System.out.println("Podaj cenę pozycji:");
                        float cena = scanner.nextFloat();
                        System.out.println("Podaj ilość pozycji:");
                        int ilosc = scanner.nextInt();
                        Pozycja pozycja = new Pozycja(nazwa, cena, ilosc);
                        faktura.dodajPozycje(pozycja);
                        System.out.println("Czy chcesz dodać kolejną pozycję? 1. Tak, 2. Nie");
                        wybor2 = scanner.nextInt();
                    } while(wybor2 != 2);
                    faktury.add(faktura);
                    break;
            
                case 3:
                    System.out.println("Podaj numer faktury do usunięcia:");
                    int numer = scanner.nextInt();
                    for(Faktura fak : faktury){
                        if(fak.getID() == numer){
                            faktury.remove(fak);
                        }
                    }
                    break;

                case 4:
                    System.out.println("Podaj numer faktury do edycji:");
                    int numer2 = scanner.nextInt();
                    Faktura edytowanaFaktura = new Faktura();
                    for(Faktura fak : faktury) {
                        if(fak.getID() == numer2){
                            edytowanaFaktura = fak;
                        }
                    }

                    edytowanaFaktura.wyswietl();
                    
                    System.out.println("1. Dodaj pozycję, 2. Usuń pozycję, 3. Edytuj pozycję.");
                    int wybor3 = scanner.nextInt();
                    switch(wybor3){
                        case 1:
                            System.out.println("Podaj nazwę pozycji:");
                            String nazwa = scanner.next();
                            System.out.println("Podaj cenę pozycji:");
                            scanner.nextLine();
                            float cena = scanner.nextFloat();
                            System.out.println("Podaj ilość pozycji:");
                            int ilosc = scanner.nextInt();
                            Pozycja pozycja = new Pozycja(nazwa, cena, ilosc);
                            edytowanaFaktura.dodajPozycje(pozycja);
                            break;
                        case 2:
                            System.out.println("Podaj numer pozycji do usunięcia:");
                            int numer4 = scanner.nextInt() - 1;
                            edytowanaFaktura.getPozycje().remove(numer4);
                            break;
                        case 3:
                            System.out.println("Podaj numer pozycji do edycji:");
                            int numer3 = scanner.nextInt() - 1;
                            System.out.print(edytowanaFaktura.getPozycje().get(numer3).getNazwa() + " --> ");
                            String nowaNazwa = scanner.next();
                            edytowanaFaktura.getPozycje().get(numer3).setNazwa(nowaNazwa);
                            scanner.nextLine();
                            System.out.print(edytowanaFaktura.getPozycje().get(numer3).getCena() + " --> ");
                            float nowaCena = scanner.nextFloat();
                            edytowanaFaktura.getPozycje().get(numer3).setCena(nowaCena);
                            System.out.print(edytowanaFaktura.getPozycje().get(numer3).getIlosc() + " --> ");
                            int nowaIlosc = scanner.nextInt();
                            edytowanaFaktura.getPozycje().get(numer3).setIlosc(nowaIlosc);

                        case 4:
                            break;

                        default:
                            System.out.println("Nie ma takiej opcji!");
                            break;
                    }
                    break;
                case 5:
                    
                case 6:

                    System.out.println("Do widzenia!");
                    scanner.close();
                    System.exit(0);
                    break;
            }
        }
    }
}