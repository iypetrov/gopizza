-- +goose Up
CREATE TYPE currency AS ENUM (
    'usd'
);

CREATE TYPE order_status AS ENUM (
    'created',
    'charged'
);

CREATE TABLE IF NOT EXISTS orders (
    id uuid PRIMARY KEY,
	intent_id text UNIQUE NOT NULL,
    user_id uuid REFERENCES users (id) ON DELETE SET NULL,
	amount double precision NOT NULL,
	currency currency NOT NULL,
    order_status order_status NOT NULL,
    updated_at timestamp,
    created_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE orders;

DROP TYPE order_status;

DROP TYPE currency;
