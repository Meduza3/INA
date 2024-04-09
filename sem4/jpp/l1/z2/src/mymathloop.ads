package MyMathLoop is
    function Silnia(N : Natural) return Natural;
    pragma Export(C, Silnia, "silnia_l");
    function NWD(A, B : Natural) return Natural;
    pragma Export(C, NWD, "nwd_l");

    type Diophantine_Solution is record
        X : Integer;
        Y : Integer;
        GCD : Natural;
    end record;

    function Extended_Euclid(A, B : Integer) return Diophantine_Solution;
    pragma Export(C, Extended_Euclid, "exeuclid_l");
end MyMathLoop;