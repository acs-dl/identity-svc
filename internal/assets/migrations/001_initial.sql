-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    email TEXT NOT NULL,
    position TEXT NOT NULL
);
INSERT INTO users (name, surname, email, position) VALUES ('Serhii', 'Pomohaiev', 'serhii.pomohaiev@distributedlab.com', 'GOD');

-- +migrate Down

DROP TABLE IF EXISTS users;
