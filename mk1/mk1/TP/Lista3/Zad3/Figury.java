public class Figury {
    public static void main(String[] args) {
        switch(args[0]){
            case "o":
                System.out.println(Figura.JedenEnum.KOLO.ObliczObwod(Integer.parseInt(args[1])));
                System.out.println(Figura.JedenEnum.KOLO.ObliczPole(Integer.parseInt(args[1])));
                System.out.println(Figura.JedenEnum.KOLO.PodajNazwe());
                break;
            case "c":
                if(args[1] == args[2]){
                    System.out.println(Figura.JedenEnum.KWADRAT.ObliczObwod(Integer.parseInt(args[1])));
                    System.out.println(Figura.JedenEnum.KWADRAT.ObliczPole(Integer.parseInt(args[1])));
                    System.out.println(Figura.JedenEnum.KWADRAT.PodajNazwe());
                } else if ( args[2] == "90"){
                    System.out.println(Figura.DwaEnum.PROSTOKAT.ObliczObwod(Integer.parseInt(args[1]), Integer.parseInt(args[2])));
                    System.out.println(Figura.DwaEnum.PROSTOKAT.ObliczPole(Integer.parseInt(args[1]), Integer.parseInt(args[2])));
                    System.out.println(Figura.DwaEnum.PROSTOKAT.PodajNazwe());
                } else {
                    System.out.println(Figura.DwaEnum.ROMB.ObliczObwod(Integer.parseInt(args[1]), Integer.parseInt(args[2])));
                    System.out.println(Figura.DwaEnum.ROMB.ObliczObwod(Integer.parseInt(args[1]), Integer.parseInt(args[2])));
                    System.out.println(Figura.DwaEnum.ROMB.PodajNazwe());
                }
                break;
            case "p":
                System.out.println(Figura.JedenEnum.PIECIOKAT.ObliczObwod(Integer.parseInt(args[1])));
                System.out.println(Figura.JedenEnum.PIECIOKAT.ObliczPole(Integer.parseInt(args[1])));
                System.out.println(Figura.JedenEnum.PIECIOKAT.PodajNazwe());
                break;
            case "s":
                System.out.println(Figura.JedenEnum.SZESCIOKAT.ObliczObwod(Integer.parseInt(args[1])));
                System.out.println(Figura.JedenEnum.SZESCIOKAT.ObliczPole(Integer.parseInt(args[1])));
                System.out.println(Figura.JedenEnum.SZESCIOKAT.PodajNazwe());
                break;
            default:
                System.out.println("To nie jest znana mi figura.");
        }
    }
}
