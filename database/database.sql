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
    user_name VARCHAR(20) PRIMARY KEY,
    balance DECIMAL(6,2)
);

INSERT INTO players(user_name, balance) VALUES("ante_95", 124);
INSERT INTO players(user_name, balance) VALUES("mrnja_53", 552);
INSERT INTO players(user_name, balance) VALUES("white_widow3", 1700.1);
INSERT INTO players(user_name, balance) VALUES("mali_simun", 70);


CREATE TABLE tickets (
    ticket_id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(20),
    payment_amount DECIMAL(6,2),
    prize_money DECIMAL(6,2)
);

ALTER TABLE tickets AUTO_INCREMENT = 1000;

CREATE TABLE played_offers (
    ticket_id INT,
    offer_id INT,
    name VARCHAR(6),
    rate VARCHAR(20)
);