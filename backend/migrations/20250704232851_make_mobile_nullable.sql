-- +goose Up
ALTER TABLE users ALTER COLUMN mobile DROP NOT NULL;

-- +goose Down
ALTER TABLE users ALTER COLUMN mobile SET NOT NULL;
