
-- name: GetRandomSong :many
SELECT * FROM songs
Order by RAND()
limit 12;

-- name: SearchSong :many
SELECT * FROM songs
where name ilike sqlc.narg(search) || '%';

-- name: CreateSong :one
INSERT INTO songs (
    name,
    thumbnail,
    path,
    lyrics,
    duration,
    release_date
) VALUES (
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6
) RETURNING * ;