package MyMathLoop is
    function Silnia(N : Natural) return Natural;
    function NWD(A, B : Natural) return Natural;

    type Diophantine_Solution is record
        X : Integer;
        Y : Integer;
        GCD : Natural;
    end record;

    function Extended_Euclid(A, B : Integer) return Diophantine_Solution;
end MyMathLoop;