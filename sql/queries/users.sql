-- name: CreateUser :one
INSERT INTO users (id, email, address, status, confirmed_at, created_at)
    VALUES ($1, $2, $3, 'pending', NULL, $4)
RETURNING
    id, email, address, status, confirmed_at, created_at;

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
    confirmed_at,
    created_at;

-- name: GetUserByEmail :one
SELECT
    id,
    email,
    address,
    status,
    confirmed_at,
    created_at
FROM
    users
WHERE
    email = $1;
