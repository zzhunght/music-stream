
-- name: GetSongCategories :many
SELECT * FROM categories;

-- name: CreateSongCategories :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;

