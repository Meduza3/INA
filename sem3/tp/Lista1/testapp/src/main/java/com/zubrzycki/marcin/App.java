package com.zubrzycki.marcin;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        Customer marcin = new Customer("Marcin");
        Product mleko = new Product("Mleko", 3);
        Product chleb = new Product("Chleb", 2);
        marcin.addToCart(mleko);
        marcin.addToCart(chleb);
        System.out.println(marcin.getName() + " kupił " + marcin.getCart().getProducts().size() + " produktów za " + marcin.getCart().calculateTotal() + " zł.");
    }
}
