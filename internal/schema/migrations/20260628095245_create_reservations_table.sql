-- +goose Up
CREATE TABLE reservations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    phone text  NULL,
    start_date date NOT  NULL,
    end_date date NOT  NULL,
    room_id int NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_reservations_rooms
        FOREIGN KEY(room_id) 
        REFERENCES rooms(id)
        ON DELETE CASCADE
);



CREATE TRIGGER set_timestamp
BEFORE UPDATE ON reservations
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();


-- +goose Down
DROP TABLE reservations;
