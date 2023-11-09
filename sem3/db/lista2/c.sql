INSERT INTO Aparat (model, producent, matryca, obiektyw, waga, typ)
VALUES
    ('Model1', 1, 101, 1, 1.5, 'kompaktowy'),
    ('Model2', 2, 102, 2, 2, 'kompaktowy'), 
    ('Model3', 3, 103, 3, 1.5, 'inny'), 
    ('Model4', 4, 104, 4, 3.5, 'kompaktowy'), 
    ('Model5', 5, 105, 5, 2.5, 'lustrzanka'), 
    ('Model6', 6, 106, 6, 2.5, 'profesjonalny'), 
    ('Model7', 7, 107, 7, 1.5, 'lustrzanka'), 
    ('Model8', 8, 108, 8, 2.5, 'lustrzanka'), 
    ('Model9', 9, 109, 9, 1.5, 'kompaktowy'), 
    ('Model10', 10, 110, 10, 1.5, 'profesjonalny'), 
    ('Model11', 11, 111, 11, 2.5, 'kompaktowy'), 
    ('Model12', 12, 112, 12, 1.5, 'lustrzanka'), 
    ('Model13', 13, 113, 13, 2.5, 'inny'), 
    ('Model14', 14, 114, 14, 1.5, 'profesjonalny'), 
    ('Model15', 15, 115, 15, 2.5, 'lustrzanka');

INSERT INTO Matryca (przekatna, rozdzielczosc, typ)
VALUES 
  (2.8, 12.0, 'Matryca1'),
  (3.2, 15.0, 'Matryca2'),
  (3.5, 18.0, 'Matryca3'),
  (4.0, 20.0, 'Matryca4'),
  (4.5, 22.0, 'Matryca5'),
  (5.0, 25.0, 'Matryca6'),
  (5.5, 30.0, 'Matryca7'),
  (6.0, 35.0, 'Matryca8'),
  (6.5, 40.0, 'Matryca9'),
  (7.0, 45.0, 'Matryca10'),
  (7.5, 50.0, 'Matryca11'),
  (8.0, 55.0, 'Matryca12'),
  (8.5, 60.0, 'Matryca13'),
  (9.0, 65.0, 'Matryca14'),
  (9.5, 70.0, 'Matryca15');

INSERT INTO Obiektyw (model, minPrzeslona, maxPrzeslona)
VALUES 
  ('Obiektyw1', 1.8, 16.0),
  ('Obiektyw2', 2.0, 22.0),
  ('Obiektyw3', 2.8, 32.0),
  ('Obiektyw4', 3.5, 45.0),
  ('Obiektyw5', 4.0, 50.0),
  ('Obiektyw6', 4.5, 55.0),
  ('Obiektyw7', 5.0, 60.0),
  ('Obiektyw8', 5.5, 65.0),
  ('Obiektyw9', 6.0, 70.0),
  ('Obiektyw10', 6.5, 75.0),
  ('Obiektyw11', 7.0, 80.0),
  ('Obiektyw12', 7.5, 85.0),
  ('Obiektyw13', 8.0, 90.0),
  ('Obiektyw14', 8.5, 95.0),
  ('Obiektyw15', 9.0, 100.0);

INSERT INTO Producent (nazwa, kraj, adresKorespondencyjny)
VALUES 
  ('Producent1', 'Polska', 'Adres1'),
  ('Producent2', 'Niemcy', 'Adres2'),
  ('Producent3', 'USA', 'Adres3'),
  ('Producent4', 'Francja', 'Adres4'),
  ('Producent5', 'Chiny', 'Adres5'),
  ('Producent6', 'Chiny', 'Adres6'),
  ('Producent7', 'Chiny', 'Adres7'),
  ('Producent8', 'Chiny', 'Adres8'),
  ('Producent9', 'Chiny', 'Adres9'),
  ('Producent10', 'Japonia', 'Adres10'),
  ('Producent11', 'Korea Południowa', 'Adres11'),
  ('Producent12', 'Brazylia', 'Adres12'),
  ('Producent13', 'Indie', 'Adres13'),
  ('Producent14', 'Rosja', 'Adres14'),
  ('Producent15', 'Kanada', 'Adres15');