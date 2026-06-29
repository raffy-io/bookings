-- name: CreateRoom :one
INSERT INTO rooms (
  room_name
) VALUES (
  $1
)
RETURNING *;