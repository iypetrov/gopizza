-- name: CreateSalad :one
INSERT INTO salads (id, name, tomatoes, garlic, onion, parmesan, chicken, image_url, price, updated_at, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING
    id, name, tomatoes, garlic, onion, parmesan, chicken, image_url, price, updated_at, created_at;

-- name: GetSaladByID :one
SELECT
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    chicken,
    image_url,
    price,
    updated_at,
    created_at
FROM
    salads
WHERE
    id = $1;

-- name: GetAllSalads :many
SELECT
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    chicken,
    image_url,
    price,
    updated_at,
    created_at
FROM
    salads;

-- name: UpdateSalad :one
UPDATE
    salads
SET
    name = $2,
    tomatoes = $3,
    garlic = $4,
    onion = $5,
    parmesan = $6,
    chicken = $7,
    image_url = $8,
    price = $9,
    updated_at = $10,
    created_at = $11
WHERE
    id = $1
RETURNING
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    chicken,
    image_url,
    price,
    updated_at,
    created_at;

-- name: DeleteSaladByID :one
DELETE FROM salads
WHERE id = $1
RETURNING
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    chicken,
    image_url,
    price,
    updated_at,
    created_at;

