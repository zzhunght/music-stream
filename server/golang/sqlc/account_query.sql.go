// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: account_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkEmailExists = `-- name: CheckEmailExists :one
SELECT id FROM accounts
WHERE email = $1 LIMIT 1
`

func (q *Queries) CheckEmailExists(ctx context.Context, email string) (int32, error) {
	row := q.db.QueryRow(ctx, checkEmailExists, email)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createAccount = `-- name: CreateAccount :one
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
) RETURNING id, name, email, created_at, updated_at
`

type CreateAccountParams struct {
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	SecretKey pgtype.Text `json:"secret_key"`
}

type CreateAccountRow struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error) {
	row := q.db.QueryRow(ctx, createAccount,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.SecretKey,
	)
	var i CreateAccountRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT a.id, a.name, a.email, a.password, a.role_id, a.is_verify, a.secret_key, a.created_at, a.updated_at,r."name" as role  FROM accounts a INNER JOIN roles r ON a.role_id = r.id WHERE email = $1
`

type GetAccountRow struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	RoleID    int32            `json:"role_id"`
	IsVerify  bool             `json:"is_verify"`
	SecretKey pgtype.Text      `json:"secret_key"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Role      string           `json:"role"`
}

func (q *Queries) GetAccount(ctx context.Context, email string) (GetAccountRow, error) {
	row := q.db.QueryRow(ctx, getAccount, email)
	var i GetAccountRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.RoleID,
		&i.IsVerify,
		&i.SecretKey,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Role,
	)
	return i, err
}

const getSecretKey = `-- name: GetSecretKey :one
SELECT secret_key FROM accounts
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetSecretKey(ctx context.Context, email string) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getSecretKey, email)
	var secret_key pgtype.Text
	err := row.Scan(&secret_key)
	return secret_key, err
}

const verifyAccount = `-- name: VerifyAccount :one
UPDATE accounts
SET is_verify = true
WHERE email = $1 RETURNING id, name, email, password, role_id, is_verify, secret_key, created_at, updated_at
`

func (q *Queries) VerifyAccount(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRow(ctx, verifyAccount, email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.RoleID,
		&i.IsVerify,
		&i.SecretKey,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}