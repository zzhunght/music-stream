
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

-- name: CountListArtists :one
SELECT count(*) as total_rows
FROM artist 
WHERE name ILIKE sqlc.arg(name_search) || '%';

-- name: GetListArtists :many
SELECT * 
FROM artist 
WHERE name ILIKE sqlc.arg(name_search) || '%'

UNION

SELECT a.*
FROM songs s
INNER JOIN songs_artist sa on s.id = sa.song_id
INNER JOIN artist a on a.id = sa.artist_id
where s.name ilike sqlc.arg(name_search) || '%';

-- name: DeleteArtist :exec

DELETE from artist WHERE id = $1;

-- name: DeleteManyArtist :exec

DELETE from artist WHERE id in (sqlc.slice(ids));