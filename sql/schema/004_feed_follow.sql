-- +goose Up
CREATE TABLE feed_follow (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    account_id UUID NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feed(id) ON DELETE CASCADE,
    UNIQUE(account_id, feed_id)
);


-- +goose Down
DROP TABLE feed_follow;