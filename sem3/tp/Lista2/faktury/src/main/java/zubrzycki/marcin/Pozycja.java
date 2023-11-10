package zubrzycki.marcin;

public class Pozycja {
    private float cena;
    private int ilosc;
    private String nazwa;

    public Pozycja(String nazwa, float cena, int ilosc) {
        this.nazwa = nazwa;
        this.cena = cena;
        this.ilosc = ilosc;
    }

    public float getCena() {
        return this.cena;
    }

    public int getIlosc() {
        return this.ilosc;
    }

    public String getNazwa() {
        return this.nazwa;
    }

    public void setCena(float cena) {
        this.cena = cena;
    }

    public void setIlosc(int ilosc) {
        this.ilosc = ilosc;
    }

    public void setNazwa(String nazwa) {
        this.nazwa = nazwa;
    }

    
}
