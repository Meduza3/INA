DELIMITER //

CREATE FUNCTION NajmniejszaPrzekatnaMatrycy(IDProducenta INT) 
RETURNS VARCHAR(30)
BEGIN
    DECLARE ModelAparatu VARCHAR(30);

    SELECT A.model INTO ModelAparatu
    FROM Aparat A
    JOIN Matryca M ON A.matryca = M.ID
    WHERE A.producent = IDProducenta
    ORDER BY M.przekatna ASC
    LIMIT 1;

    RETURN ModelAparatu;
END//

DELIMITER ;
