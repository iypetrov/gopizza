-- +goose Up
CREATE TABLE IF NOT EXISTS cart (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    pizza_id uuid,
    created_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE cart;
