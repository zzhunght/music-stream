-- name: CreateAccount :one
INSERT INTO accounts (
    name, 
    email, 
    password, 
    secret_key
) VALUES (
    $1, 
    $2, 
    $3, 
    $4
) RETURNING id, name, email, created_at, updated_at;

-- name: ChangePassword :exec
UPDATE accounts SET password = $1 WHERE email = $2;


-- name: CheckEmailExists :one
SELECT id FROM accounts
WHERE email = $1 LIMIT 1;

-- name: GetSecretKey :one
SELECT secret_key FROM accounts
WHERE email = $1 LIMIT 1;

-- name: GetAccount :one
SELECT a.*,r."name" as role  FROM accounts a INNER JOIN roles r ON a.role_id = r.id WHERE email = $1;

-- name: VerifyAccount :one
UPDATE accounts
SET is_verify = true
WHERE email = $1 RETURNING *;

