class Romb extends Czworokat{
    Romb(float a, float b, float c, float d, float alf){
        this.bok1 = a;
        this.bok2 = b;
        this.bok3 = c;
        this.bok4 = d;
        this.kat = alf;
    }

    public float Pole(){
        return (float) (bok1*bok2*Math.sin(kat));
    }

    public float Obwod(){
        return bok1+bok2+bok3+bok4;
    }

    public String Nazwa(){
        return "Romb";
    }
}