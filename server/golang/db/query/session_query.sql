
-- name: CreateSession :one
INSERT INTO session (
    id,
    email,
    refresh_token,
    client_agent,
    client_ip,
    expired_at
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetSession :one
SELECT * FROM session WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM session WHERE id = $1;

