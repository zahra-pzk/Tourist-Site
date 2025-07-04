-- +goose Up
ALTER TABLE users ADD COLUMN google_id TEXT;

-- +goose Down
ALTER TABLE users DROP COLUMN google_id;
