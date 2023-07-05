-- +migrate Up
ALTER TABLE users DROP COLUMN email;

-- +migrate Down
ALTER TABLE users ALTER COLUMN email TEXT NOT NULL;