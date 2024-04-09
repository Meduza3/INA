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
    Local_A : Integer := A;
    Local_B : Integer := B; -- Use local variables for manipulation
    Old_R, R : Integer := Local_A; -- Start with Local_A
    Old_S, S : Integer := 1;
    Old_T, T : Integer := 0;
    Quotient, Temp : Integer;
begin
    while Local_B /= 0 loop -- Use Local_B in condition
        Quotient := Old_R / Local_B;

        -- Update R
        Temp := R;
        R := Old_R - Quotient * Local_B;
        Old_R := Temp;

        -- Update S
        Temp := S;
        S := Old_S - Quotient * S;
        Old_S := Temp;

        -- Update T
        Temp := T;
        T := Old_T - Quotient * T;
        Old_T := Temp;

        -- Prepare for the next iteration
        Temp := Local_A;
        Local_A := Local_B;            -- Local_A takes the value of Local_B
        Local_B := Temp mod Local_B;   -- Local_B takes the remainder
    end loop;

    return Diophantine_Solution'(X => Old_S, Y => Old_T, GCD => Natural(Old_R));
end Extended_Euclid;

end MyMathLoop;