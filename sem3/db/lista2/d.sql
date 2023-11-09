DELIMITER //

CREATE PROCEDURE GenerujModeleAparatow()
BEGIN
    DECLARE i INT DEFAULT 1;
    DECLARE model VARCHAR(30);
    DECLARE producent_id INT;
    DECLARE matryca_id INT;
    DECLARE obiektyw_id INT;

    CREATE TEMPORARY TABLE IF NOT EXISTS UnikalneModele (
        model VARCHAR(30)
    );

    WHILE i <= 100 DO
        SET model = CONCAT('Model', i);

        IF NOT EXISTS (SELECT 1 FROM Aparat WHERE model = model) THEN
            INSERT INTO Aparat (model, producent, matryca, obiektyw, waga, typ)
            VALUES (model, FLOOR(RAND() * 5) + 1, FLOOR(RAND() * 15) + 100, FLOOR(RAND() * 15) + 100, RAND() * 2.0 + 1.0, 
                CASE 
                    WHEN RAND() < 0.25 THEN 'kompaktowy'
                    WHEN RAND() < 0.5 THEN 'lustrzanka'
                    WHEN RAND() < 0.75 THEN 'profesjonalny'
                    ELSE 'inny'
                END);
            INSERT INTO UnikalneModele (model) VALUES (model);
            SET i = i + 1;
        END IF;
    END WHILE;
    DROP TEMPORARY TABLE IF EXISTS UnikalneModele;
END//

DELIMITER ;
