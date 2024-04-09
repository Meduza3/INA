with Ada.Text_IO;
with Interfaces.C;
with Main; 

procedure main_program is
    use Ada.Text_IO;
    use Main;

    N : Interfaces.C.Int := 5;
    A : Interfaces.C.Int := 20;
    B : Interfaces.C.Int := 27;
    EA : Interfaces.C.Int := 10;
    EB : Interfaces.C.Int := 35;
    Solution1 : Diophantine_Solution;

    NL : Interfaces.C.Int := 8;
    AL : Interfaces.C.Int := 24;
    BL : Interfaces.C.Int := 33;
    EAL : Interfaces.C.Int := 5;
    EBL : Interfaces.C.Int := 140;
    Solution1L : Diophantine_Solution;


begin

    Put_Line("NL silnia to " & Interfaces.C.int'Image(Main.silnia_l(NL)));
    Put_Line("AL i BL po NWD to " & Interfaces.C.int'Image(Main.nwd_l(AL, BL)));
    Solution1L := extended_euclid_l(EAL, EBL);
    Put_Line("Rozwiazanie EAL i EBL to " & Interfaces.C.int'Image(Solution1L.X) & " i " & Interfaces.C.int'Image(Solution1L.Y));

    Put_Line("N silnia to " & Interfaces.C.int'Image(Main.silnia_r(N)));
    Put_Line("A i B po NWD to " & Interfaces.C.int'Image(Main.nwd_r(A, B)));
    Solution1 := extended_euclid_r(EA, EB);
    Put_Line("Rozwiazanie EA i EB to " & Interfaces.C.int'Image(Solution1.X) & " i " & Interfaces.C.int'Image(Solution1.Y));
end main_program;