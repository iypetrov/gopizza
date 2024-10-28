-- +goose Up
CREATE TYPE user_status AS ENUM (
    'pending',
    'confirmed'
);

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    email text NOT NULL UNIQUE,
    address text NOT NULL,
    user_status user_status NOT NULL,
    confirmed_at timestamp,
    created_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE users;

DROP TYPE user_status;
