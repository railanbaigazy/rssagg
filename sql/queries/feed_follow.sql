-- name: CreateFeedFollow :one
INSERT INTO feed_follow (id, created_at, updated_at, account_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *;