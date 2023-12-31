-- name: CreateUser :one
INSERT INTO users (
  email,
  full_name,
  hashed_password
)
VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;
