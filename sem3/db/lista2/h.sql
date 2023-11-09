DELIMITER //

CREATE TRIGGER UsunMatrycePoUsunieciuAparatu
AFTER DELETE ON Aparat
FOR EACH ROW
BEGIN
    DECLARE LiczbaAparatow INT;

    -- Sprawdź, czy istnieją jeszcze aparaty z daną matrycą
    SELECT COUNT(*) INTO LiczbaAparatow
    FROM Aparat
    WHERE matryca = OLD.matryca;

    -- Jeśli nie ma już aparatu z daną matrycą, usuń matrycę
    IF LiczbaAparatow = 0 THEN
        DELETE FROM Matryca WHERE ID = OLD.matryca;
    END IF;
END//

DELIMITER ;
