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


-- name: CheckEmailExists :one
SELECT email FROM accounts
WHERE email = $1 LIMIT 1;

-- name: GetSecretKey :one
SELECT secret_key FROM accounts
WHERE email = $1 LIMIT 1;

-- name: GetAccount :one
SELECT * FROM accounts WHERE email = $1;


