
-- name: Statistics :one
SELECT
    (SELECT COUNT(*) FROM artist) AS total_artists,
    (SELECT COUNT(*) FROM accounts) AS total_users,
    (SELECT COUNT(*) FROM songs) AS total_songs;
