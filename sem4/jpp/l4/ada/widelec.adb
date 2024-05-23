package body Widelec is
    protected body Widelec is
        entry Podnies when not W_Uzyciu is
        begin
            W_Uzyciu := True;
        end Podnies;

        procedure Odloz is
        begin
            W_Uzyciu := False;
        end Odloz;
    end Widelec;
end Widelec;