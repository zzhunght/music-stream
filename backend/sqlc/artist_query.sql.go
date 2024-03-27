// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: artist_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createArtist = `-- name: CreateArtist :one
INSERT INTO artist (
    name,
    avatar_url
) VALUES ( $1, $2 ) RETURNING id, name, avatar_url, created_at
`

type CreateArtistParams struct {
	Name      string      `json:"name"`
	AvatarUrl pgtype.Text `json:"avatar_url"`
}

func (q *Queries) CreateArtist(ctx context.Context, arg CreateArtistParams) (Artist, error) {
	row := q.db.QueryRow(ctx, createArtist, arg.Name, arg.AvatarUrl)
	var i Artist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const deleteArtist = `-- name: DeleteArtist :exec

DELETE from artist WHERE id = $1
`

func (q *Queries) DeleteArtist(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteArtist, id)
	return err
}

const deleteManyArtist = `-- name: DeleteManyArtist :exec

DELETE from artist WHERE id in ($1)
`

func (q *Queries) DeleteManyArtist(ctx context.Context, ids []int32) error {
	_, err := q.db.Exec(ctx, deleteManyArtist, ids)
	return err
}

const getListArtists = `-- name: GetListArtists :many
SELECT id, name, avatar_url, created_at 
FROM artist 
WHERE name ILIKE $1 || '%'
ORDER BY $2::text 
LIMIT COALESCE($4::int, 50)
OFFSET COALESCE($3::int, 0)
`

type GetListArtistsParams struct {
	NameSearch pgtype.Text `json:"name_search"`
	OrderBy    string      `json:"order_by"`
	Start      int32       `json:"start"`
	Size       int32       `json:"size"`
}

func (q *Queries) GetListArtists(ctx context.Context, arg GetListArtistsParams) ([]Artist, error) {
	rows, err := q.db.Query(ctx, getListArtists,
		arg.NameSearch,
		arg.OrderBy,
		arg.Start,
		arg.Size,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Artist{}
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AvatarUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateArtist = `-- name: UpdateArtist :one
UPDATE artist 
SET name = $2, avatar_url = $3 
WHERE  id = $1 
RETURNING id, name, avatar_url, created_at
`

type UpdateArtistParams struct {
	ID        int32       `json:"id"`
	Name      string      `json:"name"`
	AvatarUrl pgtype.Text `json:"avatar_url"`
}

func (q *Queries) UpdateArtist(ctx context.Context, arg UpdateArtistParams) (Artist, error) {
	row := q.db.QueryRow(ctx, updateArtist, arg.ID, arg.Name, arg.AvatarUrl)
	var i Artist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}
