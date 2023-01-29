CREATE TABLE users
(
    email     VARCHAR(256) PRIMARY KEY UNIQUE,
    firstName VARCHAR(50)  NOT NULL,
    lastName  VARCHAR(50),
    password  VARCHAR(100) NOT NULL,
    createdAt DATE         NOT NULL,
    lastLogin DATE         NOT NULL
);