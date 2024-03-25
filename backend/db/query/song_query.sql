
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


-- name: GetSongBySongCategory :many
SELECT * from  songs 
WHERE id in (
    SELECT song_id from song_categories WHERE category_id = $1
) LIMIT COALESCE(sqlc.arg(size)::int, 50)
OFFSET COALESCE(sqlc.arg(start)::int, 0);


-- name: AssociateSongArtist :exec
INSERT INTO songs_artist (song_id, artist_id, owner) VALUES ($1, $2, $3);

-- name: RemoveAssociateSongArtist :exec

DELETE FROM songs_artist  WHERE artist_id = $1 AND song_id = $2;
