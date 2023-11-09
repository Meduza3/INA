CREATE TABLE Aparat (
    model varchar(30),
    producent int,
    matryca int,
    obiektyw int,
    waga float CHECK (waga >= 0),
    typ enum('kompaktowy', 'lustrzanka', 'profesjonalny', 'inny')
);

CREATE TABLE Matryca (
    ID int PRIMARY KEY AUTO_INCREMENT,
    przekatna decimal(4,2) CHECK (przekatna >= 0),
    rozdzielczosc decimal(3,1) CHECK (rozdzielczosc >= 0),
    typ varchar(10)
);
ALTER TABLE Matryca AUTO_INCREMENT = 100;

CREATE TABLE Obiektyw (
    ID int PRIMARY KEY AUTO_INCREMENT,
    model varchar(30),
    minPrzeslona float CHECK (minPrzeslona >= 0),
    maxPrzeslona float CHECK (maxPrzeslona >= 0)
);


CREATE TABLE Producent (
    ID int PRIMARY KEY AUTO_INCREMENT,
    nazwa varchar(50) NOT NULL,
    kraj varchar(20) DEFAULT 'nieznany',
    adresKorespondencyjny varchar(100) DEFAULT 'nieznany'
);
