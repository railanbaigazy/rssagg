-- +goose Up

CREATE TABLE account (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name VARCHAR(128) NOT NULL
);


-- +goose Down
DROP TABLE account;