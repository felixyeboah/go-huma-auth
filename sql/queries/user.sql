-- name: CreateUser :one
INSERT INTO users(
    name,
    email,
    avatar,
    phone_number,
    password,
    is_verified,
    role_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;