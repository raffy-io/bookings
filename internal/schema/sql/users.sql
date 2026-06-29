-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, email, password, access_level
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;