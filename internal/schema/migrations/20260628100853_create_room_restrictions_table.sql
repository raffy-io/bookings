-- +goose Up
CREATE TABLE room_restrictions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    start_date date NOT NULL,
    end_date date NOT NULL,
    room_id int NOT NULL,
    reservation_id int NOT NULL,
    restriction_id int NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_room_restrictions_rooms
        FOREIGN KEY(room_id) 
        REFERENCES rooms(id)
        ON DELETE CASCADE, 

    CONSTRAINT fk_room_restrictions_reservations
        FOREIGN KEY(reservation_id) 
        REFERENCES reservations(id)
        ON DELETE CASCADE, 

    CONSTRAINT fk_room_restrictions_restrictions
        FOREIGN KEY(restriction_id) 
        REFERENCES restrictions(id)
        ON DELETE CASCADE  
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON room_restrictions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

-- Create Indexes for Room Restrictions (The Go/SQL way)
-- 1. Index for checking room availability via date ranges and room_id
CREATE INDEX idx_room_restrictions_room_id_dates ON room_restrictions (room_id, start_date, end_date);

-- 2. Index for looking up restrictions tied to a specific reservation
CREATE INDEX idx_room_restrictions_reservation_id ON room_restrictions (reservation_id);


-- +goose Down
DROP TABLE room_restrictions;