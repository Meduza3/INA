class Prostokat extends Czworokat{
    
    Prostokat(float a, float b, float c, float d, float alf){
        this.bok1 = a;
        this.bok2 = b;
        this.bok3 = c;
        this.bok4 = d;
        this.kat = alf;
    }

    public float Pole(){
        return bok1*bok2;
    }

    public float Obwod(){
        return 2*bok1+2*bok2;
    }

    public String Nazwa(){
        return "Prostokat";
    }


}