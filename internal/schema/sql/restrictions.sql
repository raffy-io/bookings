-- name: CreateRestriction :one
INSERT INTO restrictions (
  restriction_name
) VALUES (
  $1
)
RETURNING *;