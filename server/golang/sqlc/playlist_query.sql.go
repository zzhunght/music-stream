// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: playlist_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addSongToPlaylist = `-- name: AddSongToPlaylist :exec
INSERT INTO playlist_song (song_id, playlist_id)
VALUES ($1, $2)
`

type AddSongToPlaylistParams struct {
	SongID     int32 `json:"song_id"`
	PlaylistID int32 `json:"playlist_id"`
}

func (q *Queries) AddSongToPlaylist(ctx context.Context, arg AddSongToPlaylistParams) error {
	_, err := q.db.Exec(ctx, addSongToPlaylist, arg.SongID, arg.PlaylistID)
	return err
}

const checkOwnerPlaylist = `-- name: CheckOwnerPlaylist :one
SELECT account_id, id FROM playlist WHERE account_id = $1 and id = $2
`

type CheckOwnerPlaylistParams struct {
	AccountID pgtype.Int4 `json:"account_id"`
	ID        int32       `json:"id"`
}

type CheckOwnerPlaylistRow struct {
	AccountID pgtype.Int4 `json:"account_id"`
	ID        int32       `json:"id"`
}

func (q *Queries) CheckOwnerPlaylist(ctx context.Context, arg CheckOwnerPlaylistParams) (CheckOwnerPlaylistRow, error) {
	row := q.db.QueryRow(ctx, checkOwnerPlaylist, arg.AccountID, arg.ID)
	var i CheckOwnerPlaylistRow
	err := row.Scan(&i.AccountID, &i.ID)
	return i, err
}

const checkSongInPlaylist = `-- name: CheckSongInPlaylist :one
SELECT count(*) 
FROM playlist_song
where song_id = $1 and playlist_id = $2
`

type CheckSongInPlaylistParams struct {
	SongID     int32 `json:"song_id"`
	PlaylistID int32 `json:"playlist_id"`
}

func (q *Queries) CheckSongInPlaylist(ctx context.Context, arg CheckSongInPlaylistParams) (int64, error) {
	row := q.db.QueryRow(ctx, checkSongInPlaylist, arg.SongID, arg.PlaylistID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUserPlaylist = `-- name: CreateUserPlaylist :one
INSERT INTO playlist (account_id, name)
VALUES($1, $2) RETURNING id, name, account_id, artist_id, description, created_at
`

type CreateUserPlaylistParams struct {
	AccountID pgtype.Int4 `json:"account_id"`
	Name      string      `json:"name"`
}

func (q *Queries) CreateUserPlaylist(ctx context.Context, arg CreateUserPlaylistParams) (Playlist, error) {
	row := q.db.QueryRow(ctx, createUserPlaylist, arg.AccountID, arg.Name)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccountID,
		&i.ArtistID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deletePlaylist = `-- name: DeletePlaylist :exec
DELETE FROM playlist where account_id = $1 and id = $2
`

type DeletePlaylistParams struct {
	AccountID pgtype.Int4 `json:"account_id"`
	ID        int32       `json:"id"`
}

func (q *Queries) DeletePlaylist(ctx context.Context, arg DeletePlaylistParams) error {
	_, err := q.db.Exec(ctx, deletePlaylist, arg.AccountID, arg.ID)
	return err
}

const getPlaylistofUser = `-- name: GetPlaylistofUser :many
SELECT id, name, account_id, artist_id, description, created_at FROM playlist where account_id = $1
`

func (q *Queries) GetPlaylistofUser(ctx context.Context, accountID pgtype.Int4) ([]Playlist, error) {
	rows, err := q.db.Query(ctx, getPlaylistofUser, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AccountID,
			&i.ArtistID,
			&i.Description,
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

const getSongInPlaylist = `-- name: GetSongInPlaylist :many
SELECT s.id, s.name, s.thumbnail, s.artist_id, s.path, s.lyrics, s.duration, s.release_date, s.created_at, s.updated_at , a.name as artist_name, a.avatar_url 
from playlist_song p 
INNER JOIN songs s ON p.song_id = s.id 
LEFT JOIN artist a on s.artist_id = a.id
WHERE p.playlist_id = $1
`

type GetSongInPlaylistRow struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Thumbnail   pgtype.Text      `json:"thumbnail"`
	ArtistID    int32            `json:"artist_id"`
	Path        pgtype.Text      `json:"path"`
	Lyrics      pgtype.Text      `json:"lyrics"`
	Duration    pgtype.Int4      `json:"duration"`
	ReleaseDate pgtype.Timestamp `json:"release_date"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	ArtistName  pgtype.Text      `json:"artist_name"`
	AvatarUrl   pgtype.Text      `json:"avatar_url"`
}

func (q *Queries) GetSongInPlaylist(ctx context.Context, playlistID int32) ([]GetSongInPlaylistRow, error) {
	rows, err := q.db.Query(ctx, getSongInPlaylist, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSongInPlaylistRow{}
	for rows.Next() {
		var i GetSongInPlaylistRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Thumbnail,
			&i.ArtistID,
			&i.Path,
			&i.Lyrics,
			&i.Duration,
			&i.ReleaseDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ArtistName,
			&i.AvatarUrl,
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

const removeSongFromPlaylist = `-- name: RemoveSongFromPlaylist :exec
DELETE FROM playlist_song where playlist_id = $1 and song_id = $2
`

type RemoveSongFromPlaylistParams struct {
	PlaylistID int32 `json:"playlist_id"`
	SongID     int32 `json:"song_id"`
}

func (q *Queries) RemoveSongFromPlaylist(ctx context.Context, arg RemoveSongFromPlaylistParams) error {
	_, err := q.db.Exec(ctx, removeSongFromPlaylist, arg.PlaylistID, arg.SongID)
	return err
}

const updatePlaylist = `-- name: UpdatePlaylist :one
UPDATE  playlist 
SET name = $1
WHERE id = $2 and account_id = $3
RETURNING id, name, account_id, artist_id, description, created_at
`

type UpdatePlaylistParams struct {
	Name      string      `json:"name"`
	ID        int32       `json:"id"`
	AccountID pgtype.Int4 `json:"account_id"`
}

func (q *Queries) UpdatePlaylist(ctx context.Context, arg UpdatePlaylistParams) (Playlist, error) {
	row := q.db.QueryRow(ctx, updatePlaylist, arg.Name, arg.ID, arg.AccountID)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccountID,
		&i.ArtistID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}
