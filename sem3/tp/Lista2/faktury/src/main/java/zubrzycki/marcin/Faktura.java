package zubrzycki.marcin;
import java.util.ArrayList;
import java.util.Date;

public class Faktura {
    private ArrayList<Pozycja> pozycje;
    private int ID;
    private static int nextID = 0;
    private Date dataWystawienia;

    public Faktura() {
        this.pozycje = new ArrayList<Pozycja>();
        this.ID = nextID++;
        this.dataWystawienia = new Date();
    }

    public Faktura(ArrayList<Pozycja> pozycje) {
        this.pozycje = pozycje;
        this.ID = nextID++;
        this.dataWystawienia = new Date();
    }

    public void dodajPozycje(Pozycja pozycja) {
        this.pozycje.add(pozycja);
    }

    public void usunPozycje(int index) {
        this.pozycje.remove(index);
    }

    public float obliczWartosc() {
        float wartosc = 0;
        for (Pozycja pozycja : this.pozycje) {
            wartosc += pozycja.getCena() * pozycja.getIlosc();
        }
        return wartosc;
    }

    public int getID() {
        return this.ID;
    }

    public ArrayList<Pozycja> getPozycje() {
        return this.pozycje;
    }

    public void wyswietl() {
        System.out.println();
        System.out.println("Faktura nr " + this.ID + " z dnia " + this.dataWystawienia);
        int i = 1;
        for (Pozycja pozycja : this.pozycje) {
            System.out.println(Integer.toString(i) + ". " + pozycja.getNazwa() + " " + pozycja.getIlosc() + "x " + pozycja.getCena());
            i++;
        }
        System.out.println("Wartosc: " + this.obliczWartosc());
        System.out.println("-----------------------------");
    }

}
