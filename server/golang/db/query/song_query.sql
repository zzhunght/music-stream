
-- name: GetSongByID :one

SELECT s.*,
CASE
    WHEN COUNT(a.id) > 0 THEN jsonb_agg(jsonb_build_object('name', a.name, 'id', a.id, 'avatar_url', a.avatar_url))
    ELSE '[]'::jsonb
END AS artists 
FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
WHERE s.id = $1
GROUP BY s.id;
-- name: GetSongs :many

SELECT s.*,
CASE
    WHEN COUNT(a.id) > 0 THEN jsonb_agg(jsonb_build_object('name', a.name, 'id', a.id, 'avatar_url', a.avatar_url))
    ELSE '[]'::jsonb
END AS artists 
FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
GROUP BY s.id
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50);


-- name: GetSongOfArtist :many

SELECT s.*,
CASE
    WHEN COUNT(a.id) > 0 THEN jsonb_agg(jsonb_build_object('name', a.name, 'id', a.id, 'avatar_url', a.avatar_url))
    ELSE '[]'::jsonb
END AS artists 
FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
WHERE a.id = $1
GROUP BY s.id
OFFSET COALESCE(sqlc.arg(start)::int, 0)
LIMIT COALESCE(sqlc.arg(size)::int, 50);

-- name: GetRandomSong :many
SELECT s.*,
CASE
    WHEN COUNT(a.id) > 0 THEN jsonb_agg(jsonb_build_object('name', a.name, 'id', a.id, 'avatar_url', a.avatar_url))
    ELSE '[]'::jsonb
END AS artists 
FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
GROUP BY s.id
Order by RANDOM()
limit 15;

-- name: SearchSong :many
SELECT s.*,
CASE
    WHEN COUNT(a.id) > 0 THEN jsonb_agg(jsonb_build_object('name', a.name, 'id', a.id, 'avatar_url', a.avatar_url))
    ELSE '[]'::jsonb
END AS artists 
FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
where s.name ilike sqlc.narg(search) || '%'
GROUP BY s.id
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

-- name: UpdateSong :exec

UPDATE songs 
SET name = sqlc.arg(name), thumbnail = sqlc.arg(thumbnail), 
path = sqlc.arg(path), lyrics = sqlc.arg(lyrics), duration = sqlc.arg(duration), release_date = sqlc.arg(release_date)
WHERE id = sqlc.arg(id);

-- name: GetSongBySongCategory :many
SELECT s.*,
CASE
    WHEN COUNT(a.id) > 0 THEN jsonb_agg(jsonb_build_object('name', a.name, 'id', a.id, 'avatar_url', a.avatar_url))
    ELSE '[]'::jsonb
END AS artists 
FROM songs s
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
WHERE s.id in (
    SELECT song_id from song_categories WHERE category_id = $1
) 
GROUP BY s.id
LIMIT COALESCE(sqlc.arg(size)::int, 50)
OFFSET COALESCE(sqlc.arg(start)::int, 0);


-- name: AssociateSongArtist :exec
INSERT INTO songs_artist (song_id, artist_id, owner) VALUES ($1, $2, true);


-- name: UpdateAssociateSongArtist :exec
UPDATE  songs_artist  SET artist_id =$1
WHERE song_id = $2;

-- name: RemoveAssociateSongArtist :exec

DELETE FROM songs_artist  WHERE artist_id = $1 AND song_id = $2;


-- name: DeleteSong :exec
DELETE FROM songs  WHERE id = $1;
