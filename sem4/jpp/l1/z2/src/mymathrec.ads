package MyMathRec is
    function Silnia(N : Natural) return Natural;
    pragma Export(C, Silnia, "silnia_r");
    function NWD(A, B : Natural) return Natural;
    pragma Export(C, NWD, "nwd_r");

    type Diophantine_Solution is record
        X : Integer;
        Y : Integer;
        GCD : Natural;
    end record;

    function Extended_Euclid(A, B, C : Integer) return Diophantine_Solution;
    pragma Export(C, Extended_Euclid, "exeuclid_r");
end MyMathRec;