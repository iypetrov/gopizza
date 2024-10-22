-- name: CreateUser :one
INSERT INTO users (id, email, address, status, confirmed_at)
    VALUES ($1, $2, $3, 'pending', NULL)
RETURNING
    id, email, address, status, confirmed_at;

-- name: ConfirmUser :one
UPDATE
    users
SET
    status = 'confirmed',
    confirmed_at = $2
WHERE
    id = $1
RETURNING
    id,
    email,
    address,
    status,
    confirmed_at;

