DROP DATABASE simple;
CREATE DATABASE simple;

USE simple;

CREATE TABLE Users (

    `logn` VARCHAR(20) PRIMARY KEY,
    `passwd` VARCHAR(20) CHECK( `passwd` != "" ),
    `name` VARCHAR(20),
    `age` INT,
    `gender` VARCHAR(20) 

);

INSERT INTO Users(logn, passwd) VALUES("KoiP", "123456");
INSERT INTO Users(logn, passwd) VALUES("Jannet", "ilovefrogs");
INSERT INTO Users(logn, passwd) VALUES("Tort", "graal321");
INSERT INTO Users(logn, passwd) VALUES("Marie", "admin");
INSERT INTO Users(logn, passwd) VALUES("Boris", "654321");