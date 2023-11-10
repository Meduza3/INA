package zubrzycki.marcin;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

import org.junit.Before;
import org.junit.Test;
import java.util.ArrayList;

/**
 * Unit test for simple App.
 */
public class AppTest 
{
    private App app;

    @Before
    public void setUp() {
        app = new App();
    }

    private ArrayList<Faktura> faktury = new ArrayList<Faktura>();

    @Test
    public void testAddFaktura() {
        Faktura faktura = new Faktura();
        Pozycja pozycja = new Pozycja("Test", 10.0f, 2);
        faktura.dodajPozycje(pozycja);

        assertEquals(1, faktura.getPozycje().size());
    }

    @Test
    public void testRemoveFaktura() {
        Faktura faktura = new Faktura();
        faktury.add(faktura);

        assertEquals(1, faktury.size());

        faktury.remove(faktura);

        assertEquals(0, faktury.size());
    }

    @Test
    public void testEditPozycja() {
        Faktura faktura = new Faktura();
        Pozycja pozycja = new Pozycja("Stara", 10.0f, 2);
        faktura.dodajPozycje(pozycja);

        assertEquals("Stara", faktura.getPozycje().get(0).getNazwa());

        faktura.getPozycje().get(0).setNazwa("Nowa");

        assertEquals("Nowa", faktura.getPozycje().get(0).getNazwa());
    }
}
