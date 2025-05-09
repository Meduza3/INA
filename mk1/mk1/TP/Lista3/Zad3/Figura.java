public class Figura {
    public enum JedenEnum implements Jeden {
        KOLO{
            @Override
            public float ObliczObwod(float a) {
                return (float) 3.14159 * 2 * a;
            }

            @Override
            public float ObliczPole(float a) {
                return (float) 3.14159 * a * a;
            }

            @Override
            public String PodajNazwe() {
                return "Koło";
            }
        },
        KWADRAT{
            @Override
            public float ObliczObwod(float a) {
                return 4 * a;
            }

            @Override
            public float ObliczPole(float a) {
                return a * a;
            }

            @Override
            public String PodajNazwe() {
                return "Kwadrat";
            }
        },
        PIECIOKAT{
            @Override
            public float ObliczObwod(float a) {
                return 5 * a;
            }

            @Override
            public float ObliczPole(float a) {
                return (float) ((float) (5*a*a)/(4*Math.tan(3.14159/5)));
            }

            @Override
            public String PodajNazwe() {
                return "Pięciokąt";
            }
        },
        SZESCIOKAT{
            @Override
            public float ObliczObwod(float a) {
                return 6 * a;
            }

            @Override
            public float ObliczPole(float a) {
                return (float) ((float) (6*a*a)/(4*Math.tan(3.14159/6)));
            }

            @Override
            public String PodajNazwe() {
                return "Sześciokąt";
            }

        }

    }

    public enum DwaEnum implements Dwa {
        PROSTOKAT{
            @Override
            public float ObliczObwod(float a, float b) {
                return 2 * a + 2 * b;
            }

            @Override
            public float ObliczPole(float a, float b) {
                return a * b;
            }

            @Override
            public String PodajNazwe() {
                return "Prostokąt";
            }
        },
        ROMB{
            @Override
            public String PodajNazwe() {
                return "Romb";
            }

            @Override
            public float ObliczObwod(float a, float b) {
                return 4 * a;
            }

            @Override
            public float ObliczPole(float a, float b) {
                return (float) (a * a * Math.sin(b));
            }
        }
    }

}

    
