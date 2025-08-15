-- +goose UP
CREATE TABLE users (
    id uuid PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text UNIQUE NOT NULL
);

-- +goose DOWN
DROP TABLE users;

