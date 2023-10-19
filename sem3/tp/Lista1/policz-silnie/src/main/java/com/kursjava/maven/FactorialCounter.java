package com.kursjava.maven;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
public class FactorialCounter {
  private static final Logger log =
      LogManager.getLogger(FactorialCounter.class);
  public static void main(String[] args) {
    System.out.println("Przyklad liczenia silni.");
    System.out.println("Silnia 5 = " + factorial(5));
  }
  public static int factorial(int n) {
    if (n < 0) {
      log.error("Nieprawidlowa wartosc n: {}", n);
      throw new IllegalArgumentException(
          "Silnia moze byc liczona tylko dla n >= 0"
      );
    }
    int result = 1;
    for (int i = 2; i <= n; i++) {
      result *= i;
    }
    return result;
  }
}