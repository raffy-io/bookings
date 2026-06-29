-- +goose Up
CREATE TABLE restrictions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    restriction_name text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON restrictions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();


-- +goose Down
DROP TABLE restrictions;


