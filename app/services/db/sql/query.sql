-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: AddUser :one
INSERT INTO users (
  name
) VALUES (
  $1
)
RETURNING *;