-- +goose Up
CREATE TYPE product_type AS ENUM (
    'pizza'
);

CREATE TABLE IF NOT EXISTS carts (
    id uuid PRIMARY KEY,
    user_id uuid REFERENCES users (id) ON DELETE SET NULL,
    pizza_id uuid REFERENCES pizzas (id) ON DELETE SET NULL,
    product_type product_type NOT NULL,
    created_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE carts;

DROP TYPE product_type;
