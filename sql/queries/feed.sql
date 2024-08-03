-- name: CreateFeed :one
INSERT INTO feed (id, created_at, updated_at, name, url, account_id) 
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feed;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feed
    ORDER BY last_fetched_at NULLS FIRST
    LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feed 
    SET last_fetched_at=NOW(), updated_at=NOW()
    WHERE id=$1
    RETURNING *;