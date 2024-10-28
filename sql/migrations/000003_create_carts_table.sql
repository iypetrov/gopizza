-- +goose Up
CREATE TABLE IF NOT EXISTS carts (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    pizza_id uuid REFERENCES pizzas (id) ON DELETE SET NULL,
    created_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE carts;

