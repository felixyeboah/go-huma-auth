-- name: CreateUser :one
INSERT INTO users(
    name,
    email,
    phone_number,
    password,
    role_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * from users
WHERE email = $1 LIMIT 1;

-- name: VerifyUser :exec
UPDATE users
SET is_verified = true
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users
SET password = $2
WHERE id = $1;