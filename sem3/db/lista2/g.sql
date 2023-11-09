CREATE FUNCTION LiczbaModeliZMatryca(IDMatrycy INT) 
RETURNS INT
BEGIN
    DECLARE LiczbaModeli INT;

    SELECT COUNT(*) INTO LiczbaModeli
    FROM Aparat
    WHERE matryca = IDMatrycy;

    RETURN LiczbaModeli;
END;
