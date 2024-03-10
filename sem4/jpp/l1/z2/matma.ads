package Matma is
    function Silnia(N : Natural) return Natural;
    
    function NWD(A, B : Integer) return Integer;
    
    type Solution is record
        X : Integer;
        Y : Integer;
        Exists: Boolean;
    end record;
    function NWDRozszerzony(A, B : Integer; X : out Integer; Y : out Integer) return Integer;
    function Diophantine(A, B, C : Integer) return Solution;
end Matma;