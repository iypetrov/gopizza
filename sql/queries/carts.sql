-- name: AddPizzaToCart :one
INSERT INTO carts (id, user_id, pizza_id, created_at)
    VALUES ($1, $2, $3, $4)
RETURNING
    id, user_id, pizza_id, created_at;

-- name: GetCartByUserID :many
SELECT
    c.id AS cart_id,
    COALESCE(pizzas.name) AS product_name,
    COALESCE(pizzas.image_url) AS product_image_url,
    COALESCE(pizzas.price) AS product_price
FROM
    carts c
    LEFT JOIN pizzas ON c.pizza_id = pizzas.id
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
    created_at;

