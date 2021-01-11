CREATE DATABASE start;

USE start;

CREATE TABLE Users (

    login VARCHAR(20) PRIMARY KEY,
    password VARCHAR(20) CHECK( password != "" )

);

INSERT INTO Users(login, password) VALUES("KoiP", "123456");
INSERT INTO Users(login, password) VALUES("Jannet", "ilovefrogs");
INSERT INTO Users(login, password) VALUES("Tort", "graal321");
INSERT INTO Users(login, password) VALUES("Marie", "admin");
INSERT INTO Users(login, password) VALUES("Boris", "654321");