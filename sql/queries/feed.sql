-- name: CreateFeed :one
INSERT INTO feed (id, created_at, updated_at, name, url, account_id) 
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;