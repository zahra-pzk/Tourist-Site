-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    google_id TEXT,
    name TEXT,
    email TEXT UNIQUE,
    picture TEXT
);

-- +goose Down
DROP TABLE IF EXISTS users;
