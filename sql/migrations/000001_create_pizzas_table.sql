-- +goose Up

CREATE TABLE IF NOT EXISTS PIZZAS (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    image_url TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE PIZZAS;
