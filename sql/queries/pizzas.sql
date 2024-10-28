-- name: CreatePizza :one
INSERT INTO pizzas (id, name, tomatoes, garlic, onion, parmesan, cheddar, pepperoni, sausage, ham, bacon, chicken, salami, ground_beef, mushrooms, olives, spinach, pineapple, arugula, anchovies, capers, image_url, price, updated_at, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
RETURNING
    id, name, tomatoes, garlic, onion, parmesan, cheddar, pepperoni, sausage, ham, bacon, chicken, salami, ground_beef, mushrooms, olives, spinach, pineapple, arugula, anchovies, capers, image_url, price, updated_at, created_at;

-- name: GetPizzaByID :one
SELECT
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    cheddar,
    pepperoni,
    sausage,
    ham,
    bacon,
    chicken,
    salami,
    ground_beef,
    mushrooms,
    olives,
    spinach,
    pineapple,
    arugula,
    anchovies,
    capers,
    image_url,
    price,
    updated_at,
    created_at
FROM
    pizzas
WHERE
    id = $1;

-- name: GetAllPizzas :many
SELECT
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    cheddar,
    pepperoni,
    sausage,
    ham,
    bacon,
    chicken,
    salami,
    ground_beef,
    mushrooms,
    olives,
    spinach,
    pineapple,
    arugula,
    anchovies,
    capers,
    image_url,
    price,
    updated_at,
    created_at
FROM
    pizzas;

-- name: UpdatePizza :one
UPDATE
    pizzas
SET
    name = $2,
    tomatoes = $3,
    garlic = $4,
    onion = $5,
    parmesan = $6,
    cheddar = $7,
    pepperoni = $8,
    sausage = $9,
    ham = $10,
    bacon = $11,
    chicken = $12,
    salami = $13,
    ground_beef = $14,
    mushrooms = $15,
    olives = $16,
    spinach = $17,
    pineapple = $18,
    arugula = $19,
    anchovies = $20,
    capers = $21,
    image_url = $22,
    price = $23,
    updated_at = $24,
    created_at = $25
WHERE
    id = $1
RETURNING
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    cheddar,
    pepperoni,
    sausage,
    ham,
    bacon,
    chicken,
    salami,
    ground_beef,
    mushrooms,
    olives,
    spinach,
    pineapple,
    arugula,
    anchovies,
    capers,
    image_url,
    price,
    updated_at,
    created_at;

-- name: DeletePizzaByID :one
DELETE FROM pizzas
WHERE id = $1
RETURNING
    id,
    name,
    tomatoes,
    garlic,
    onion,
    parmesan,
    cheddar,
    pepperoni,
    sausage,
    ham,
    bacon,
    chicken,
    salami,
    ground_beef,
    mushrooms,
    olives,
    spinach,
    pineapple,
    arugula,
    anchovies,
    capers,
    image_url,
    price,
    updated_at,
    created_at;

