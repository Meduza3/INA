import java.util.ArrayList;

public class main {

  public static final String BLACK = "\u001B[30m";
  public static final String GREEN = "\u001B[32m";
  public static final String YELLOW = "\u001B[33m";
  public static final String BLUE = "\u001B[34m";
  public static final String PURPLE = "\u001B[35m";
  public static final String CYAN = "\u001B[36m";
  public static final String WHITE = "\u001B[37m";


  public static int networkLength = 60;
  public static int printSpeed = 80;
  public static int delayRange = 100;

  public static void main(String[] args) throws InterruptedException {

    String[] network = new String[networkLength];


    Transmitter t1 = new Transmitter("A", 30, YELLOW);
    Transmitter t2 = new Transmitter("B",10, CYAN);
    Transmitter t3 = new Transmitter("C",50, GREEN);

    ArrayList<Transmitter> transmitters = new ArrayList<Transmitter>();

    transmitters.add(t1);
    transmitters.add(t2);
    transmitters.add(t3);

    Controller ctrl = new Controller(network,transmitters, printSpeed, delayRange);

    ctrl.run();


  }



}