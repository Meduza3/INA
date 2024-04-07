package body MyMathLoop is
    function Silnia(N : Natural) return Natural is
        Res : Natural := 1;
    begin
        for J in 1..N loop
            Res := Res * J;
        end loop;
        return Res;
    end Silnia;

    function NWD(A, B : Natural) return Natural is
        Temp : Natural;
        X : Natural := A;
        Y : Natural := B;
    begin
        while Y /= 0 loop
            Temp := X mod Y;
            X := Y;
            Y := Temp;
        end loop;
        return X;
    end NWD;

    function Extended_Euclid(A, B : Integer) return Diophantine_Solution is
        Old_R, R : Integer := A;
        Old_S, S : Integer := 1;
        Old_T, T : Integer := 0;
        Quotient, Temp : Integer;
    begin
        while B /= 0 loop
            Quotient := Old_R / B;
            Temp := R;
            R := Old_R - Quotient * B;
            Old_R := Temp;

            Temp := S;
            S := Old_S - Quotient * S;
            Old_S := Temp;

            Temp := T;
            T := Old_T - Quotient * T;
            Old_T := Temp;

            Temp := A;
            A := B;
            B := Temp;
        end loop;
        return Diophantine_Solution'(X => Old_S, Y => Old_T, GCD => Natural(Old_R)); -- Cast Old_R to Natural as GCD cannot be negative.
    end Extended_Euclid;
end MyMathLoop;