// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: follow_query.sql

package sqlc

import (
	"context"
)

const checkFollow = `-- name: CheckFollow :one
SELECT id, account_id, artist_id, created_at FROM artist_follow WHERE account_id = $1 and artist_id = $2
`

type CheckFollowParams struct {
	AccountID int32 `json:"account_id"`
	ArtistID  int32 `json:"artist_id"`
}

func (q *Queries) CheckFollow(ctx context.Context, arg CheckFollowParams) (ArtistFollow, error) {
	row := q.db.QueryRow(ctx, checkFollow, arg.AccountID, arg.ArtistID)
	var i ArtistFollow
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.ArtistID,
		&i.CreatedAt,
	)
	return i, err
}

const follow = `-- name: Follow :exec
INSERT INTO artist_follow (account_id,artist_id) VALUES ($1,$2)
`

type FollowParams struct {
	AccountID int32 `json:"account_id"`
	ArtistID  int32 `json:"artist_id"`
}

func (q *Queries) Follow(ctx context.Context, arg FollowParams) error {
	_, err := q.db.Exec(ctx, follow, arg.AccountID, arg.ArtistID)
	return err
}

const unFollow = `-- name: UnFollow :exec
DELETE FROM artist_follow WHERE account_id = $1 and artist_id = $2
`

type UnFollowParams struct {
	AccountID int32 `json:"account_id"`
	ArtistID  int32 `json:"artist_id"`
}

func (q *Queries) UnFollow(ctx context.Context, arg UnFollowParams) error {
	_, err := q.db.Exec(ctx, unFollow, arg.AccountID, arg.ArtistID)
	return err
}