with Widelec;

package Filozof is
    task type Filozof_Task (Id : Integer; lewy, prawy : not null access Widelec.Widelec; Ilosc_Posilkow : Integer) is
        entry Start;
    end Filozof_Task;
end Filozof;