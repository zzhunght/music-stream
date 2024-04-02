
-- name: GetSongs :many

SELECT s.*, a.name as artist_name FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50);


-- name: GetRandomSong :many
SELECT * FROM songs
Order by RAND()
limit 12;

-- name: SearchSong :many
SELECT s.*, a.name as artist_name FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
where name ilike sqlc.narg(search) || '%'
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50)
;

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

-- name: UpdateSong :one

UPDATE songs 
SET name = sqlc.arg(name), thumbnail = sqlc.arg(thumbnail), 
path = sqlc.arg(path), lyrics = sqlc.arg(lyrics), duration = sqlc.arg(duration), release_date = sqlc.arg(release_date)
WHERE id = sqlc.arg(id)
RETURNING * ;

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


-- name: DeleteSong :exec
DELETE FROM songs  WHERE id = $1;
