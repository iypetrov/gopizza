-- name: AddPizzaToCart :one
INSERT INTO carts (id, user_id, pizza_id, product_type, created_at)
    VALUES ($1, $2, $3, 'pizza', $4)
RETURNING
    id, user_id, pizza_id, product_type, created_at;

-- name: AddSaladToCart :one
INSERT INTO carts (id, user_id, salad_id, product_type, created_at)
    VALUES ($1, $2, $3, 'salad', $4)
RETURNING
    id, user_id, salad_id, product_type, created_at;

-- name: GetCartByUserID :many
SELECT
    c.id AS cart_id,
    CASE
        WHEN c.product_type = 'pizza' THEN pizzas.name::text
        WHEN c.product_type = 'salad' THEN salads.name::text
        ELSE NULL::text
    END AS product_name,
    CASE
        WHEN c.product_type = 'pizza' THEN pizzas.image_url::text
        WHEN c.product_type = 'salad' THEN salads.image_url::text
        ELSE NULL::text
    END AS product_image_url,
    c.product_type::text AS product_type,
    CASE
        WHEN c.product_type = 'pizza' THEN pizzas.price::float8
        WHEN c.product_type = 'salad' THEN salads.price::float8
        ELSE NULL::float8
    END AS product_price
FROM
    carts c
    LEFT JOIN pizzas ON c.pizza_id = pizzas.id
    LEFT JOIN salads ON c.salad_id = salads.id
WHERE
    c.user_id = $1;

-- name: RemoveItemFromCart :exec
DELETE FROM carts
WHERE id = $1;

-- name: EmptyCartByUserID :one
DELETE FROM carts
WHERE user_id = $1
RETURNING
    id,
    user_id,
    pizza_id,
    product_type,
    created_at;
