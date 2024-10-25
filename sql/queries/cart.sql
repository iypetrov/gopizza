-- name: AddPizzaToCart :one
INSERT INTO cart (id, user_id, pizza_id, created_at)
    VALUES ($1, $2, $3, $4)
RETURNING
    id, user_id, pizza_id, created_at;

-- name: EmptyCartByUserID :one
DELETE FROM cart
WHERE user_id = $1
RETURNING
    id,
    user_id,
    pizza_id,
    created_at;
 