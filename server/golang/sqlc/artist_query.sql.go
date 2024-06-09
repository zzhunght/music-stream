// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: artist_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countListArtists = `-- name: CountListArtists :one
SELECT count(*) as total_rows
FROM artist 
WHERE name ILIKE $1 || '%'
`

func (q *Queries) CountListArtists(ctx context.Context, nameSearch pgtype.Text) (int64, error) {
	row := q.db.QueryRow(ctx, countListArtists, nameSearch)
	var total_rows int64
	err := row.Scan(&total_rows)
	return total_rows, err
}

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

const getArtistById = `-- name: GetArtistById :one
SELECT id, name, avatar_url, created_at FROM artist WHERE id = $1
`

func (q *Queries) GetArtistById(ctx context.Context, id int32) (Artist, error) {
	row := q.db.QueryRow(ctx, getArtistById, id)
	var i Artist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getListArtists = `-- name: GetListArtists :many
SELECT id, name, avatar_url, created_at 
FROM artist 
WHERE name ILIKE $1 || '%'

UNION

SELECT a.id, a.name, a.avatar_url, a.created_at
FROM songs s
INNER JOIN songs_artist sa on s.id = sa.song_id
INNER JOIN artist a on a.id = sa.artist_id
where s.name ilike $1 || '%'

ORDER BY created_at DESC
`

func (q *Queries) GetListArtists(ctx context.Context, nameSearch pgtype.Text) ([]Artist, error) {
	rows, err := q.db.Query(ctx, getListArtists, nameSearch)
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
