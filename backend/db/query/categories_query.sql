
-- name: GetSongCategories :many
SELECT * FROM categories;

-- name: CreateCategories :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;


-- name: UpdateCategories :one

UPDATE categories 
SET name = $1 
WHERE id = $2
RETURNING *;


-- name: DeleteCategories :exec

DELETE FROM categories WHERE id = $1;

