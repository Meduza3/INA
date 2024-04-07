package body MyMathRec is
    function Silnia(N : Natural) return Natural is
    begin
        if N = 0 then
            return 1;
        else
            return N * Silnia(N - 1);
        end if;
    end Silnia;

    function NWD(A, B : Natural) return Natural is
    begin
        if B = 0 then
            return A;
        else
            return NWD(B, A mod B);
        end if;
    end NWD;

    function Extended_Euclid(A, B : Integer) return Diophantine_Solution is
    function Recursive_Extended_Euclid(A, B : Integer; X1, Y1, X2, Y2 : Integer) return Diophantine_Solution is
        Temp_X, Temp_Y : Integer;
    begin
        if B = 0 then
            return Diophantine_Solution(X => X1, Y => Y1, GCD => Natural(A)); -- Base case
        else
            Temp_X := X2 - (A / B) * X1;
            Temp_Y := Y2 - (A / B) * Y1;
            return Recursive_Extended_Euclid(B, A mod B, X2, Y2, Temp_X, Temp_Y);
        end if;
    end Recursive_Extended_Euclid;
    begin
        return Recursive_Extended_Euclid(A, B, 1, 0, 0, 1); -- Initial call with starting values
    end Extended_Euclid;
end MyMathRec;