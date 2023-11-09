CREATE VIEW WidokAparatowProducentow AS
SELECT 
    P.nazwa AS nazwa_producenta,
    P.kraj,
    A.model
FROM 
    Aparat A
JOIN 
    Producent P ON A.producent = P.ID;
