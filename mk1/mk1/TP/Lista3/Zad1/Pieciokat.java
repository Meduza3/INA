class Pieciokat extends Figura{
    public float bok;

    Pieciokat(float bok){
        this.bok = bok;
    }

    public float Pole(){
        return (float) ((float) (5*bok*bok)/(4*Math.tan(3.14159/5)));
    }

    public float Obwod(){
        return 5*bok;
    }

    public String Nazwa(){
        return "Pięciokąt";
    }
}