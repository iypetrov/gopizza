-- +goose Up
CREATE TABLE IF NOT EXISTS pizzas (
    id uuid PRIMARY KEY,
    name text NOT NULL UNIQUE,
    tomatoes boolean NOT NULL,
    garlic boolean NOT NULL,
    onion boolean NOT NULL,
    parmesan boolean NOT NULL,
    cheddar boolean NOT NULL,
    pepperoni boolean NOT NULL,
    sausage boolean NOT NULL,
    ham boolean NOT NULL,
    bacon boolean NOT NULL,
    chicken boolean NOT NULL,
    salami boolean NOT NULL,
    ground_beef boolean NOT NULL,
    mushrooms boolean NOT NULL,
    olives boolean NOT NULL,
    spinach boolean NOT NULL,
    pineapple boolean NOT NULL,
    arugula boolean NOT NULL,
    anchovies boolean NOT NULL,
    capers boolean NOT NULL,
    image_url text NOT NULL,
    price double precision NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_pizzas_price ON pizzas (price);

-- +goose Down
DROP INDEX IF EXISTS idx_pizzas_price;

DROP TABLE pizzas;

