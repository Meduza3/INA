CREATE VIEW WidokAparatowLustrzanek AS
SELECT 
    A.model,
    A.waga,
    P.nazwa AS nazwa_producenta,
    M.przekatna,
    M.rozdzielczosc,
    O.minPrzeslona,
    O.maxPrzeslona
FROM 
    Aparat A
JOIN 
    Producent P ON A.producent = P.ID
JOIN 
    Matryca M ON A.matryca = M.ID
JOIN 
    Obiektyw O ON A.obiektyw = O.ID
WHERE 
    A.typ = 'lustrzanka'
    AND P.kraj != 'Chiny';
