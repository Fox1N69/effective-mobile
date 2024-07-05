CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_number VARCHAR(20),
    surname VARCHAR(255),
    name VARCHAR(255),
    patronymic VARCHAR(255),
    address TEXT
);