
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
SELECT albums.id, albums.name, albums.thumbnail, albums.release_date FROM albums where name ilike sqlc.arg(search) || '%'
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
DELETE FROM albums_songs 
WHERE album_id = $1 AND song_id = ANY(sqlc.arg(song_ids)::int[]);


-- name: GetAlbumSong :many
SELECT s.*,
CASE
    WHEN COUNT(a.id) > 0 THEN jsonb_agg(jsonb_build_object('name', a.name, 'id', a.id, 'avatar_url', a.avatar_url))
    ELSE '[]'::jsonb
END AS artists 
from albums_songs
INNER JOIN songs s ON albums_songs.song_id = s.id 
LEFT JOIN songs_artist sa on s.id = sa.song_id
LEFT JOIN artist a on a.id = sa.artist_id
WHERE albums_songs.album_id = $1
GROUP BY s.id;

-- name: GetSongNotInAlbum :many
SELECT s.id ,s.name , s.thumbnail, s.duration, s.created_at, s.release_date from songs s
where id not in (SELECT als.song_id FROM albums_songs als WHERE als.album_id = $1) and name ilike sqlc.arg(search) || '%'
order by s.created_at desc;

-- name: GetLatestAlbum :many
SELECT a.*, art.name as artist_name from albums a
INNER JOIN artist art on a.artist_id = art.id
ORDER BY a.created_at DESC
OFFSET 0
LIMIT 20;

-- name: GetAlbumInfoFromSongID :one
SELECT al.id, al.name , al.thumbnail, al.release_date from albums al
INNER JOIN albums_songs abs on al.id = abs.album_id
WHERE abs.song_id = $1 LIMIT 1;