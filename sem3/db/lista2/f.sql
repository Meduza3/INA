DELIMITER //

CREATE TRIGGER DodajModelPoWstawieniu
AFTER INSERT ON Aparat
FOR EACH ROW
BEGIN
    DECLARE ProducentIstnieje INT;

    -- Sprawdź, czy producent istnieje w tabeli Producent
    SELECT COUNT(*) INTO ProducentIstnieje
    FROM Producent
    WHERE ID = NEW.producent;

    -- Jeśli producent nie istnieje, dodaj go
    IF ProducentIstnieje = 0 THEN
        INSERT INTO Producent (ID, nazwa, kraj, adresKorespondencyjny)
        VALUES (NEW.producent, 'Nazwa Producenta', 'Nieznany', 'Brak Adresu');
    END IF;
END//

DELIMITER ;
