import java.util.ArrayList;

public class TrojkatPascala {
    private final ArrayList<ArrayList<Long>> wiersze;

    public TrojkatPascala(int liczbaWierszy){
        wiersze = new ArrayList<ArrayList<Long>>();
        for(int i = 0; i < liczbaWierszy; i++){
            wiersze.add(new ArrayList<Long>());
            for(int j = 0; j < i+1; j++){
                wiersze.get(i).add(newton(i,j));
            }
        }
    }

    private static long newton(int n, int k) {
        long numerator = 1L;
        long denominator = 1L;
        for (int i = 1; i <= k; i++) {
            numerator *= n - (k - i);
            denominator *= i;
        }
        return numerator / denominator;
    }


    public String printTriangle(){
        StringBuilder result = new StringBuilder();
        for(int i = 0; i < wiersze.size(); i++){
            for(int j = 0; j <= i; j++){
                result.append(" ").append(newton(i, j));
            }
            result.append('\n');
        }
        return result.toString();
    }

    public static void main(String[] args) {
        TrojkatPascala trojkat = new TrojkatPascala(15);
        System.out.println(trojkat.printTriangle());
    }
}
