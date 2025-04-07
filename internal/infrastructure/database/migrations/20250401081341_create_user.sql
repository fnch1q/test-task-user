-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100) NULL,
    age INTEGER NOT NULL,
    gender VARCHAR(10) NOT NULL,
    nationality VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
