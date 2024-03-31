
-- name: CreateAlbum :one
INSERT INTO albums (
    name,
    artist_id,
    thumbnail,
    release_date
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;


-- name: CountAlbumsByArtistID :one
SELECT COUNT(*) AS total_rows FROM albums WHERE artist_id = $1;
-- name: GetAlbumByArtistID :one
SELECT * FROM albums WHERE artist_id = $1;

-- name: CountAlbums :one
SELECT COUNT(*) AS total_rows FROM albums;
-- name: GetAlbums :many
SELECT * FROM albums
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50);


-- name: CountSearchAlbums :one
SELECT COUNT(*) AS total_rows FROM albums WHERE name ILIKE sqlc.arg(search) || '%';

-- search: SearchAlbums :many
SELECT * FROM albums where name ilike sqlc.narg(search) || '%'
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50);

-- name: UpdateAlbum :exec
UPDATE albums SET
    name = $2,
    artist_id = $3,
    thumbnail = $4,
    release_date = $5
WHERE id = $1;

-- name: DeleteAlbum :exec
DELETE FROM albums WHERE id = $1;

-- name: AddSongToAlbum :exec
INSERT INTO albums_songs (
    album_id,
    song_id
) VALUES($1, $2);

-- name: RemoveSongToAlbum :exec
INSERT INTO albums_songs (
    album_id,
    song_id
) VALUES($1, $2);


-- name: GetAlbumSong :many
SELECT s.* from albums_songs a INNER JOIN songs s ON a.song_id = s.id WHERE a.id = $1;