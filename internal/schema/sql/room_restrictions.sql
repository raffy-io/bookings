-- name: CreateRoomRestriction :one
INSERT INTO room_restrictions (
  start_date,end_date,room_id,reservation_id, restriction_id 
) VALUES (
  $1,$2,$3,$4,$5
)
RETURNING *;