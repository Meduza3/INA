with Widelec;
with Filozof;

procedure Main is
   Ilosc_Filozofow : Integer := 5;
   Ilosc_Posilkow : Integer := 5;

   Widelce : array (1 .. Ilosc_Filozofow) of aliased Widelec.Widelec;

   type wsk is access Filozof.Filozof_Task;
   wsks : array (1 .. Ilosc_Filozofow) of wsk;

   Lewy_Index : Integer;
   Prawy_Index : Integer;
begin
   for I in 1 .. Ilosc_Filozofow loop
      if I mod 2 = 0 then
         Lewy_Index := (I mod Ilosc_Filozofow) + 1;
         Prawy_Index := I;
      else
         Lewy_Index := I;
         Prawy_Index := (I mod Ilosc_Filozofow) + 1;
      end if;

      wsks (I) := new Filozof.Filozof_Task (
         I, Widelce (Lewy_Index)'Access, Widelce (Prawy_Index)'Access, Ilosc_Posilkow
      );
   end loop;

   for I in 1 .. Ilosc_Filozofow loop
      wsks (I).Start;
   end loop;
end Main;