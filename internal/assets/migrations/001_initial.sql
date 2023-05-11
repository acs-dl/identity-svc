-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    email TEXT NOT NULL,
    telegram TEXT,
    position TEXT NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS users;
