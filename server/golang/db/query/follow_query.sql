

-- name: CheckFollow :one
SELECT * FROM artist_follow WHERE account_id = $1 and artist_id = $2;


-- name: Follow :exec
INSERT INTO artist_follow (account_id,artist_id) VALUES ($1,$2);

-- name: UnFollow :exec
DELETE FROM artist_follow WHERE account_id = $1 and artist_id = $2;