CREATE DATABASE OlympicsFanzone;

CREATE TABLE Schedule(
    ID INTEGER PRIMARY KEY AUTO_INCREMENT,
    EVENTS VARCHAR(255) NOT NULL,
    STARTDATE VARCHAR(255) NOT NULL, 
    ENDDATE VARCHAR(255) NOT NULL,
    COUNTRY VARCHAR(255) NOT NULL
);