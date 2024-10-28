-- +goose Up
CREATE TABLE IF NOT EXISTS salads (
    id uuid PRIMARY KEY,
    name text NOT NULL UNIQUE,
    tomatoes boolean NOT NULL,
    garlic boolean NOT NULL,
    onion boolean NOT NULL,
    parmesan boolean NOT NULL,
	chicken boolean NOT NULL,
    image_url text NOT NULL,
    price double precision NOT NULL,
    updated_at timestamp NOT NULL,
    created_at timestamp NOT NULL
);

ALTER TYPE product_type ADD VALUE 'salad';

ALTER TABLE carts
ADD COLUMN salad_id uuid REFERENCES salads (id) ON DELETE SET NULL;

-- +goose Down
DROP TABLE salads;

ALTER TABLE carts DROP COLUMN IF EXISTS salad_id;

CREATE TYPE product_type_old AS ENUM ('pizza');

ALTER TABLE carts
  ALTER COLUMN product_type TYPE product_type_old
  USING product_type::text::product_type_old;

DROP TYPE product_type;
ALTER TYPE product_type_old RENAME TO product_type;
