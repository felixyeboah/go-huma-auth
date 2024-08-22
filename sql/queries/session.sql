-- name: CreateSession :one
INSERT INTO sessions (
    access_token,
    refresh_token,
    user_id,
    expiry_date,
    user_agent,
    ip_address,
    last_accessed_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;
