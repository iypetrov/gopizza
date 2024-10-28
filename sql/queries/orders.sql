-- name: CreateOrder :one
INSERT INTO orders (id, intent_id, user_id, amount, currency, order_status, updated_at, created_at)
    VALUES ($1, $2, $3, $4, 'usd', 'created', NULL, $5)
RETURNING
    id, intent_id, user_id, amount, currency, order_status, updated_at, created_at;

-- name: GetOrderByIntentID :one
SELECT
    o.id,
    o.intent_id,
    o.user_id,
    o.amount,
    o.currency,
    o.order_status,
    o.updated_at,
    o.created_at,
    u.address
FROM orders o
JOIN users u ON o.user_id = u.id
WHERE o.intent_id = $1;

-- name: ChargeOrder :one
UPDATE orders
SET order_status = 'charged', updated_at = $2
WHERE id = $1
RETURNING
    id, intent_id, user_id, amount, currency, order_status, updated_at, created_at;
