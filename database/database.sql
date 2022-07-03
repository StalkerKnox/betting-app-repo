CREATE TABLE offers (
    number VARCHAR(10),
    tv_channel VARCHAR(10),
    offer_id INT PRIMARY KEY,
    title VARCHAR(40),
    has_statistics BOOLEAN,
    time DATETIME
);

CREATE TABLE rates (
    offer_id INT,
    rate float,
    name VARCHAR(6)
);

CREATE TABLE leagues (
    id INT AUTO_INCREMENT PRIMARY KEY ,
    title VARCHAR(40)
);

CREATE TABLE elaborations (
    elaboration_id INT AUTO_INCREMENT PRIMARY KEY,
    league_id INT
);

ALTER TABLE elaborations AUTO_INCREMENT = 10;

CREATE TABLE types (
    type_id INT AUTO_INCREMENT PRIMARY KEY,
    elaboration_id INT,
    name VARCHAR(20)
);

CREATE TABLE connect (
    elaboration_id INT,
    offer_id INT
);


CREATE TABLE players (
    user_name VARCHAR(20),
    balance INT
);

INSERT INTO players(user_name, balance) VALUES("ante_95", 124);
INSERT INTO players(user_name, balance) VALUES("mrnja_53", 552);
INSERT INTO players(user_name, balance) VALUES("white_widow3", 1700);
INSERT INTO players(user_name, balance) VALUES("mali_simun", 70);