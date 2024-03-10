package body Matma is
    function Silnia(N : Natural) return Natural is
        Result : Natural := 1;
        begin
            if N = 0 then
                return Result;
            else
                for I in 1..N loop
                    Result := Result * I;
                end loop;
                return Result;
            end if;
        end Silnia;

        function NWD(A, B : Integer) return Integer is
            Temp : Integer;
            X : Integer := A;
            Y : Integer := B;
            begin
                while Y /= 0 loop
                    Temp := X mod Y;
                    X := Y;
                    Y := Temp;
                end loop;
                return X;
        end NWD;

        function NWDRozszerzony(A, B : Integer; X : out Integer; Y : out Integer) return Integer is
            X0, Y0, X1, Y1 : Integer;
            Temp : Integer;
            A_Temp, B_Temp : Integer := A;
        begin
            X := 0; Y := 1; X1 := 1; Y1 := 0;
            while B_Temp /= 0 loop
                Temp := X;
                X := X1 - (A_Temp / B_Temp) * X;
                X1 := Temp;

                Temp := Y;
                Y := Y1 - (A_Temp / B_Temp) * Y;
                Y1 := Temp;

                Temp := A_Temp mod B_TEmp;
                A_Temp := B_Temp;
                B_Temp := Temp;
            end loop;
            X := X1;
            Y := Y1;
            return A_Temp;
        end NWDRozszerzony;

        function Diophantine(A, B, C : Integer) return Solution is
            X, Y, D : Integer;
            Sol : Solution;
        begin
            D := NWDRozszerzony(A, B, X, Y);
            if C mod D /= 0 then
                Sol.Exists := False;
                return Sol;
            else
                Sol.X := X * (C / D);
                Sol.Y := Y * (C / D);
                Sol.Exists := True;
                return Sol;
            end if;
        end Diophantine;
    end Matma;