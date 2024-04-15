
-- name: CreatePlaylist :one
INSERT INTO playlist (account_id, name)
VALUES($1, $2) RETURNING *;

-- name: GetPlaylistofUser :many
SELECT * FROM playlist where account_id = $1;

-- name: GetSongInPlaylist :many
SELECT s.* from playlist_song p INNER JOIN songs s ON p.song_id = s.id WHERE p.playlist_id = $1;


-- name: UpdatePlaylist :one
UPDATE  playlist 
SET name = $1
WHERE id = $2 and account_id = $3
RETURNING *;

-- name: CheckOwnerPlaylist :one
SELECT account_id, id FROM playlist WHERE account_id = $1 and id = $2;


-- name: DeletePlaylist :exec
DELETE FROM playlist where account_id = $1 and id = $2;

-- name: AddSongToPlaylist :exec
INSERT INTO playlist_song (song_id, playlist_id)
VALUES ($1, $2);

-- name: CheckSongInPlaylist :one
SELECT count(*) 
FROM playlist_song
where song_id = $1 and playlist_id = $2;


-- name: RemoveSongFromPlaylist :exec
DELETE FROM playlist_song where playlist_id = $1 and song_id = $2;