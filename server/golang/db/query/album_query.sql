
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
    sqlc.arg(release_date)::date
) RETURNING *;



-- name: CountAlbumsByArtistID :one
SELECT COUNT(*) AS total_rows FROM albums WHERE artist_id = $1;
-- name: GetAlbumByArtistID :many
SELECT * FROM albums WHERE artist_id = $1;

-- name: CountAlbums :one
SELECT COUNT(*) AS total_rows FROM albums;
-- name: GetAlbums :many
SELECT * FROM albums
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50);


-- name: CountSearchAlbums :one
SELECT COUNT(*) AS total_rows FROM albums WHERE name ILIKE sqlc.arg(search) || '%';

-- name: SearchAlbums :many
SELECT * FROM albums where name ilike sqlc.narg(search) || '%'
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50);

-- name: UpdateAlbum :one
UPDATE albums SET
    name = $2,
    artist_id = $3,
    thumbnail = $4,
    release_date = $5
WHERE id = $1 RETURNING *;

-- name: DeleteAlbum :exec
DELETE FROM albums WHERE id = $1;

-- name: AddSongToAlbum :exec
INSERT INTO albums_songs (
    album_id,
    song_id
) VALUES($1, $2);

-- name: AddManySongToAlbum :exec
INSERT INTO albums_songs (
    album_id,
    song_id
)  VALUES (  
  $1,  
  unnest(@song_ids::int[])  
);

-- name: RemoveSongFromAlbum :exec
DELETE FROM albums_songs WHERE id = ANY($1::int[]);


-- name: GetAlbumSong :many
SELECT s.* from albums_songs a INNER JOIN songs s ON a.song_id = s.id WHERE a.album_id = $1;


-- name: GetLatestAlbum :many
SELECT a.*, art.name as artist_name from albums a
INNER JOIN artist art on a.artist_id = art.id
ORDER BY a.created_at DESC
OFFSET 0
LIMIT 20;