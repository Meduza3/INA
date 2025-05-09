class Szesciokat extends Figura{
    public float bok;

    Szesciokat(float bok){
        this.bok = bok;
    }

    public float Pole(){
        return (float) ((float) (6*bok*bok)/(4*Math.tan(3.14159/6)));
    }

    public float Obwod(){
        return 6*bok;
    }

    public String Nazwa(){
        return "Sześciokąt";
    }
}