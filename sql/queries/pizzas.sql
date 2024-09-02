-- name: CreatePizza :one

INSERT INTO pizzas (
    id, name, tomatoes, garlic, onion, parmesan, cheddar, pepperoni, sausage, ham, bacon, chicken, salami, ground_beef, mushrooms, olives, spinach, pineapple, arugula, anchovies, capers, image_url, price, updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24
)
RETURNING *;
