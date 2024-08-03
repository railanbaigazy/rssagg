-- +goose Up
CREATE TABLE feed (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    account_id UUID NOT NULL REFERENCES account(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed;