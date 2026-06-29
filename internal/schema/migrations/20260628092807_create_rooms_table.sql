-- +goose Up
CREATE TABLE rooms (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    room_name text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON rooms
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();


-- +goose Down
DROP TABLE rooms;



