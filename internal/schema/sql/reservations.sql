-- name: CreateReservation :one
INSERT INTO reservations (
  first_name,last_name,email,phone,start_date,end_date,room_id
) VALUES (
  $1,$2,$3,$4,$5,$6,$7
)
RETURNING *;