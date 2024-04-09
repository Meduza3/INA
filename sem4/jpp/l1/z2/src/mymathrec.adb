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

    function Extended_Euclid(A, B, C : Integer) return Diophantine_Solution is
      		GCD_AB : Integer;
      		Sign_A, Sign_B : Integer;
      		Div : Integer;
      		Solution, Next_Solution : Diophantine_Solution;
   	 begin

		if A < 0 then
			Sign_A := -1;
		else
			Sign_A := 1;
		end if;
		if B < 0 then
			Sign_B := -1;
		else
			Sign_B := 1;
		end if;

            GCD_AB := NWD(Abs(A), Abs(B));
      		Div := C / GCD_AB;
      		Solution.X := Sign_A * Div;
      		Solution.Y := 0;

      		if (C - A * Solution.X) mod B /= 0 then
         		Next_Solution := Extended_Euclid(B, A mod B, C - A * Solution.X);
         		Solution.X := Solution.X + Sign_A * Next_Solution.X;
         		Solution.Y := Solution.Y + Sign_B * Next_Solution.Y;
      		else
         		Solution.Y := (C - A * Solution.X) / B;
      		end if;

      		return Solution;
   	end Extended_Euclid;
end MyMathRec;