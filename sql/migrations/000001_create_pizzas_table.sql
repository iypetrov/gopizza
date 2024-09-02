-- +goose Up

CREATE TABLE IF NOT EXISTS pizzas (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    tomatoes BOOLEAN NOT NULL,
    garlic BOOLEAN NOT NULL,
    onion BOOLEAN NOT NULL,
    parmesan BOOLEAN NOT NULL,
    cheddar BOOLEAN NOT NULL,
    pepperoni BOOLEAN NOT NULL,
    sausage BOOLEAN NOT NULL,
    ham BOOLEAN NOT NULL,
    bacon BOOLEAN NOT NULL,
    chicken BOOLEAN NOT NULL,
    salami BOOLEAN NOT NULL,
    ground_beef BOOLEAN NOT NULL,
    mushrooms BOOLEAN NOT NULL,
    olives BOOLEAN NOT NULL,
    spinach BOOLEAN NOT NULL,
    pineapple BOOLEAN NOT NULL,
    arugula BOOLEAN NOT NULL,
    anchovies BOOLEAN NOT NULL,
    capers BOOLEAN NOT NULL,
    image_url TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE pizzas;
