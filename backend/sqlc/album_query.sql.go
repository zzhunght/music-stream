// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: album_query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addSongToAlbum = `-- name: AddSongToAlbum :exec
INSERT INTO albums_songs (
    album_id,
    song_id
) VALUES($1, $2)
`

type AddSongToAlbumParams struct {
	AlbumID int32 `json:"album_id"`
	SongID  int32 `json:"song_id"`
}

func (q *Queries) AddSongToAlbum(ctx context.Context, arg AddSongToAlbumParams) error {
	_, err := q.db.Exec(ctx, addSongToAlbum, arg.AlbumID, arg.SongID)
	return err
}

const countAlbums = `-- name: CountAlbums :one
SELECT COUNT(*) AS total_rows FROM albums
`

func (q *Queries) CountAlbums(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countAlbums)
	var total_rows int64
	err := row.Scan(&total_rows)
	return total_rows, err
}

const countAlbumsByArtistID = `-- name: CountAlbumsByArtistID :one
SELECT COUNT(*) AS total_rows FROM albums WHERE artist_id = $1
`

func (q *Queries) CountAlbumsByArtistID(ctx context.Context, artistID int32) (int64, error) {
	row := q.db.QueryRow(ctx, countAlbumsByArtistID, artistID)
	var total_rows int64
	err := row.Scan(&total_rows)
	return total_rows, err
}

const countSearchAlbums = `-- name: CountSearchAlbums :one
SELECT COUNT(*) AS total_rows FROM albums WHERE name ILIKE $1 || '%'
`

func (q *Queries) CountSearchAlbums(ctx context.Context, search pgtype.Text) (int64, error) {
	row := q.db.QueryRow(ctx, countSearchAlbums, search)
	var total_rows int64
	err := row.Scan(&total_rows)
	return total_rows, err
}

const createAlbum = `-- name: CreateAlbum :one
INSERT INTO albums (
    name,
    artist_id,
    thumbnail,
    release_date
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING id, artist_id, thumbnail, name, release_date, created_at
`

type CreateAlbumParams struct {
	Name        string      `json:"name"`
	ArtistID    int32       `json:"artist_id"`
	Thumbnail   string      `json:"thumbnail"`
	ReleaseDate pgtype.Date `json:"release_date"`
}

func (q *Queries) CreateAlbum(ctx context.Context, arg CreateAlbumParams) (Album, error) {
	row := q.db.QueryRow(ctx, createAlbum,
		arg.Name,
		arg.ArtistID,
		arg.Thumbnail,
		arg.ReleaseDate,
	)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.ArtistID,
		&i.Thumbnail,
		&i.Name,
		&i.ReleaseDate,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAlbum = `-- name: DeleteAlbum :exec
DELETE FROM albums WHERE id = $1
`

func (q *Queries) DeleteAlbum(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteAlbum, id)
	return err
}

const getAlbumByArtistID = `-- name: GetAlbumByArtistID :one
SELECT id, artist_id, thumbnail, name, release_date, created_at FROM albums WHERE artist_id = $1
`

func (q *Queries) GetAlbumByArtistID(ctx context.Context, artistID int32) (Album, error) {
	row := q.db.QueryRow(ctx, getAlbumByArtistID, artistID)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.ArtistID,
		&i.Thumbnail,
		&i.Name,
		&i.ReleaseDate,
		&i.CreatedAt,
	)
	return i, err
}

const getAlbumSong = `-- name: GetAlbumSong :many
SELECT s.id, s.name, s.thumbnail, s.path, s.lyrics, s.duration, s.release_date, s.created_at, s.updated_at from albums_songs a INNER JOIN songs s ON a.song_id = s.id WHERE a.id = $1
`

func (q *Queries) GetAlbumSong(ctx context.Context, id int32) ([]Song, error) {
	rows, err := q.db.Query(ctx, getAlbumSong, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Song{}
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Thumbnail,
			&i.Path,
			&i.Lyrics,
			&i.Duration,
			&i.ReleaseDate,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getAlbums = `-- name: GetAlbums :many
SELECT id, artist_id, thumbnail, name, release_date, created_at FROM albums
OFFSET COALESCE($1::int, 0)
LIMIT COALESCE($2::int, 50)
`

type GetAlbumsParams struct {
	Start int32 `json:"start"`
	Size  int32 `json:"size"`
}

func (q *Queries) GetAlbums(ctx context.Context, arg GetAlbumsParams) ([]Album, error) {
	rows, err := q.db.Query(ctx, getAlbums, arg.Start, arg.Size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Album{}
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.ArtistID,
			&i.Thumbnail,
			&i.Name,
			&i.ReleaseDate,
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

const removeSongToAlbum = `-- name: RemoveSongToAlbum :exec
INSERT INTO albums_songs (
    album_id,
    song_id
) VALUES($1, $2)
`

type RemoveSongToAlbumParams struct {
	AlbumID int32 `json:"album_id"`
	SongID  int32 `json:"song_id"`
}

func (q *Queries) RemoveSongToAlbum(ctx context.Context, arg RemoveSongToAlbumParams) error {
	_, err := q.db.Exec(ctx, removeSongToAlbum, arg.AlbumID, arg.SongID)
	return err
}

const updateAlbum = `-- name: UpdateAlbum :exec
UPDATE albums SET
    name = $2,
    artist_id = $3,
    thumbnail = $4,
    release_date = $5
WHERE id = $1
`

type UpdateAlbumParams struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	ArtistID    int32       `json:"artist_id"`
	Thumbnail   string      `json:"thumbnail"`
	ReleaseDate pgtype.Date `json:"release_date"`
}

func (q *Queries) UpdateAlbum(ctx context.Context, arg UpdateAlbumParams) error {
	_, err := q.db.Exec(ctx, updateAlbum,
		arg.ID,
		arg.Name,
		arg.ArtistID,
		arg.Thumbnail,
		arg.ReleaseDate,
	)
	return err
}
