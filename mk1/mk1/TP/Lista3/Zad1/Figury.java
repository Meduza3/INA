class Figury{
    public static void main(String[] args) {
        Figura a = null;
       switch(args[0]){
        case "o":
            a = new Kolo(Integer.parseInt(args[1]));
            break;
        
        case "c":
        //kwadrat, prostokat lub romb
            if (args[1] == args[2] && args[2] == args[3] && args[3] == args[4] && Integer.parseInt(args[5]) == 90){
                a = new Kwadrat(Integer.parseInt(args[1]), Integer.parseInt(args[2]), Integer.parseInt(args[3]), Integer.parseInt(args[4]), Integer.parseInt(args[5]));
            } else if (Integer.parseInt(args[5]) == 90) {
                a = new Prostokat(Integer.parseInt(args[1]), Integer.parseInt(args[2]), Integer.parseInt(args[3]), Integer.parseInt(args[4]), Integer.parseInt(args[5]));
            } else if (args[1] == args[2] && args[2] == args[3] && args[3] == args[4] && Integer.parseInt(args[5]) != 90){
                a = new Romb(Integer.parseInt(args[1]), Integer.parseInt(args[2]), Integer.parseInt(args[3]), Integer.parseInt(args[4]), Integer.parseInt(args[5]));
            }
            break;

        case "p":
            a = new Pieciokat(Integer.parseInt(args[1]));
            break;

        case "s":
            a = new Szesciokat(Integer.parseInt(args[1]));
            break;

        default:
            System.out.println("To nie jest znana mi figura.");
       }
       if(a != null){
            System.out.println("Twoja figura to " + a.Nazwa() + ".");
            System.out.println("Jej pole to " + a.Pole() + " jednostek.");
            System.out.println("Jej obwód to " + a.Obwod() + " jednostek.");
       }
    }
}