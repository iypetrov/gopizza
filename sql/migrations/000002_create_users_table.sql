-- +goose Up
CREATE TYPE status AS ENUM (
  'pending',
  'confirmed'
);

CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY,
    email text NOT NULL UNIQUE,
    address text NOT NULL,
	status status  NOT NULL,
    confirmed_at timestamp
);

-- +goose Down
DROP TABLE users;
