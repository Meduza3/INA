with Ada.Numerics.Float_Random; use Ada.Numerics.Float_Random;
with Ada.Text_IO; use Ada.Text_IO;

package body Filozof is
    task body Filozof_Task is
        Rng : Generator;
        Min_Dobranoc : Float := 5.0;
        Max_Dobranoc : Float := 10.0;
    begin
        accept Start;
        Reset (Rng);
        for I in 1 .. Ilosc_Posilkow loop
            Put_Line ("Philosopher" & Integer'Image(Id) & " is thinking.");
            delay Duration (Min_Dobranoc + Random (Rng) * (Max_Dobranoc - Min_Dobranoc));
            Put_Line ("Philosopher" & Integer'Image(Id) & " stopped thinking.");

            lewy.Podnies;
            prawy.Odloz;

            Put_Line ("Philosopher" & Integer'Image(Id) & " is eating.");
            delay Duration (Min_Dobranoc + Random (Rng) * (Max_Dobranoc - Min_Dobranoc));
            Put_Line ("Philosopher" & Integer'Image(Id) & " stopped eating.");

            lewy.Podnies;
            prawy.Odloz;
        end loop;
    end Filozof_Task;
end Filozof;