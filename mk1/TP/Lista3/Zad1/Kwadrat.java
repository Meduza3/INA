class Kwadrat extends Czworokat{

    Kwadrat(float a, float b, float c, float d, float alf){
        this.bok1 = a;
        this.bok2 = b;
        this.bok3 = c;
        this.bok4 = d;
        this.kat = alf;
    }

    public float Pole(){
        return bok1*bok1;
    }

    public float Obwod(){
        return bok1*4;
    }

    public String Nazwa(){
        return "Kwadrat";
    }
}