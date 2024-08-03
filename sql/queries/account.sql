-- name: CreateAccount :one
INSERT INTO account (id, created_at, updated_at, name, api_key)
    VALUES ($1, $2, $3, $4, 
    encode(sha256(random()::text::bytea), 'hex')
    )
    RETURNING *;

-- name: GetAccountByAPIKey :one
SELECT * FROM account WHERE api_key=$1;