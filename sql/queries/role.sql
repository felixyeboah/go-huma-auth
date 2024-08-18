-- name: GetRole :one
SELECT * from roles
WHERE id = $1 LIMIT 1;

-- name: GetRoleByName :one
SELECT id, name, description
FROM roles
WHERE name = $1 LIMIT 1;