
-- name: GetSongCategories :many
SELECT * FROM categories;

-- name: CreateCategories :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;


-- name: UpdateCategories :one

UPDATE categories 
SET name = $1 
WHERE id = $2
RETURNING *;

-- name: AddSongToCategory :exec
INSERT INTO song_categories (song_id, category_id) VALUES ($1, $2);

-- name: UpdateSongCategory :exec
UPDATE  song_categories set category_id = $1 WHERE song_id = $2;


-- name: DeleteCategories :exec

DELETE FROM categories WHERE id = $1;

