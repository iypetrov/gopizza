-- name: InitOrder :one
INSERT INTO orders (id, intent_id, user_id, amount, currency, order_status, updated_at, created_at)
    VALUES ($1, $2, $3, $4, 'usd', 'created', NULL, $5)
RETURNING
    id, intent_id, user_id, amount, currency, order_status, updated_at, created_at;

-- name: GetOrderByIntentID :one
SELECT
    id, intent_id, user_id, amount, currency, order_status, updated_at, created_at
FROM orders
WHERE intent_id = $1;

-- name: ChargeOrder :one
UPDATE orders
SET order_status = 'charged', updated_at = $2
WHERE id = $1
RETURNING
    id, intent_id, user_id, amount, currency, order_status, updated_at, created_at;
