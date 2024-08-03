-- name: CreateFeedFollow :one
INSERT INTO feed_follow (id, created_at, updated_at, account_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follow WHERE account_id=$1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow WHERE id=$1 AND account_id=$2;