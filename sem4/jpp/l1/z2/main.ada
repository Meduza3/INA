with Ada.Text_IO; use Ada.Text_IO;
with Matma; use Matma;

procedure Main is
    N : Natural;
    M : Natural;
    Result : Natural;
    Sol : Solution;
begin
    N := 14;
    M := 645;
    -- Result := Matma.Silnia(N);
    -- Put_Line(Natural'Image(Result));

    Result := Matma.NWD(N, M);
    Sol := Diophantine(15, -3, 9);
    Put_Line(Natural'Image(Result));
    
end Main;