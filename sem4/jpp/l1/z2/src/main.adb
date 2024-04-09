with Ada.Text_IO; use Ada.Text_IO;
with MyMathLoop; use MyMathLoop;

procedure Main is
    N : Natural := 10;
    Result : Natural;
begin
    Result := Silnia(N);
    Put_Line("Factorial of " & Natural'Image(N) & " is " & Natural'Image(Result));
end Main;