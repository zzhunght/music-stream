
-- name: CreateArtist :one
INSERT INTO artist (
    name,
    avatar_url
) VALUES ( $1, $2 ) RETURNING *;

-- name: UpdateArtist :one
UPDATE artist 
SET name = $2, avatar_url = $3 
WHERE  id = $1 
RETURNING *;

-- name: GetListArtists :many
SELECT * 
FROM artist 
WHERE name ILIKE sqlc.arg(name_search) || '%'
ORDER BY sqlc.arg(order_by)::text 
LIMIT COALESCE(sqlc.arg(size)::int, 50)
OFFSET COALESCE(sqlc.arg(start)::int, 0);

-- name: DeleteArtist :exec

DELETE from artist WHERE id = $1;

-- name: DeleteManyArtist :exec

DELETE from artist WHERE id in (sqlc.slice(ids));