-- name: GetCommentById :one
SELECT * FROM comment WHERE id = sqlc.arg(id)::int;


-- name: CreateComment :one
INSERT INTO comment (content, user_id, song_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetSongComment :many
SELECT c.id, c.content, u.name, c.created_at FROM comment c
INNER JOIN accounts u ON c.user_id = u.id
WHERE c.song_id = $1;

-- name: DeleteComment :exec
DELETE FROM comment WHERE id = sqlc.arg(id)::int;