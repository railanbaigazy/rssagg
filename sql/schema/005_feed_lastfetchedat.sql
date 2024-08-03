-- +goose Up
ALTER TABLE feed ADD last_fetched_at TIMESTAMPTZ;

-- +goose Down
ALTER TABLE feed DROP COLUMN last_fetched_at;