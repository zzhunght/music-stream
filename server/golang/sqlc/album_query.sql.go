// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: album_query.sql

package sqlc

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const addManySongToAlbum = `-- name: AddManySongToAlbum :exec
INSERT INTO albums_songs (
    album_id,
    song_id
)  VALUES (  
  $1,  
  unnest($2::int[])  
)
`

type AddManySongToAlbumParams struct {
	AlbumID int32   `json:"album_id"`
	SongIds []int32 `json:"song_ids"`
}

func (q *Queries) AddManySongToAlbum(ctx context.Context, arg AddManySongToAlbumParams) error {
	_, err := q.db.Exec(ctx, addManySongToAlbum, arg.AlbumID, arg.SongIds)
	return err
}

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
    $4::date
) RETURNING id, artist_id, name, thumbnail, release_date, created_at
`

type CreateAlbumParams struct {
	Name        string    `json:"name"`
	ArtistID    int32     `json:"artist_id"`
	Thumbnail   string    `json:"thumbnail"`
	ReleaseDate time.Time `json:"release_date"`
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
		&i.Name,
		&i.Thumbnail,
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

const getAlbumByArtistID = `-- name: GetAlbumByArtistID :many
SELECT id, artist_id, name, thumbnail, release_date, created_at FROM albums WHERE artist_id = $1
`

func (q *Queries) GetAlbumByArtistID(ctx context.Context, artistID int32) ([]Album, error) {
	rows, err := q.db.Query(ctx, getAlbumByArtistID, artistID)
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
			&i.Name,
			&i.Thumbnail,
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

const getAlbumSong = `-- name: GetAlbumSong :many
SELECT s.id, s.name, s.thumbnail, s.path, s.lyrics, s.duration, s.release_date, s.created_at, s.updated_at from albums_songs a INNER JOIN songs s ON a.song_id = s.id WHERE a.album_id = $1
`

func (q *Queries) GetAlbumSong(ctx context.Context, albumID int32) ([]Song, error) {
	rows, err := q.db.Query(ctx, getAlbumSong, albumID)
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
SELECT id, artist_id, name, thumbnail, release_date, created_at FROM albums
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
			&i.Name,
			&i.Thumbnail,
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

const getLatestAlbum = `-- name: GetLatestAlbum :many
SELECT a.id, a.artist_id, a.name, a.thumbnail, a.release_date, a.created_at, art.name as artist_name from albums a
INNER JOIN artist art on a.artist_id = art.id
ORDER BY a.created_at DESC
OFFSET 0
LIMIT 20
`

type GetLatestAlbumRow struct {
	ID          int32            `json:"id"`
	ArtistID    int32            `json:"artist_id"`
	Name        string           `json:"name"`
	Thumbnail   string           `json:"thumbnail"`
	ReleaseDate time.Time        `json:"release_date"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	ArtistName  string           `json:"artist_name"`
}

func (q *Queries) GetLatestAlbum(ctx context.Context) ([]GetLatestAlbumRow, error) {
	rows, err := q.db.Query(ctx, getLatestAlbum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetLatestAlbumRow{}
	for rows.Next() {
		var i GetLatestAlbumRow
		if err := rows.Scan(
			&i.ID,
			&i.ArtistID,
			&i.Name,
			&i.Thumbnail,
			&i.ReleaseDate,
			&i.CreatedAt,
			&i.ArtistName,
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

const getSongNotInAlbum = `-- name: GetSongNotInAlbum :many
SELECT s.id ,s.name , s.thumbnail, s.duration, s.created_at, s.release_date from songs s
where id not in (SELECT als.song_id FROM albums_songs als WHERE als.album_id = $1) and name ilike $2 || '%'
order by s.created_at desc
`

type GetSongNotInAlbumParams struct {
	AlbumID int32       `json:"album_id"`
	Search  pgtype.Text `json:"search"`
}

type GetSongNotInAlbumRow struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Thumbnail   pgtype.Text      `json:"thumbnail"`
	Duration    pgtype.Int4      `json:"duration"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	ReleaseDate pgtype.Date      `json:"release_date"`
}

func (q *Queries) GetSongNotInAlbum(ctx context.Context, arg GetSongNotInAlbumParams) ([]GetSongNotInAlbumRow, error) {
	rows, err := q.db.Query(ctx, getSongNotInAlbum, arg.AlbumID, arg.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSongNotInAlbumRow{}
	for rows.Next() {
		var i GetSongNotInAlbumRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Thumbnail,
			&i.Duration,
			&i.CreatedAt,
			&i.ReleaseDate,
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

const removeSongFromAlbum = `-- name: RemoveSongFromAlbum :exec
DELETE FROM albums_songs 
WHERE album_id = $1 AND song_id = ANY($2::int[])
`

type RemoveSongFromAlbumParams struct {
	AlbumID int32   `json:"album_id"`
	SongIds []int32 `json:"song_ids"`
}

func (q *Queries) RemoveSongFromAlbum(ctx context.Context, arg RemoveSongFromAlbumParams) error {
	_, err := q.db.Exec(ctx, removeSongFromAlbum, arg.AlbumID, arg.SongIds)
	return err
}

const searchAlbums = `-- name: SearchAlbums :many
SELECT id, artist_id, name, thumbnail, release_date, created_at FROM albums where name ilike $1 || '%'
OFFSET COALESCE($2::int, 0)
LIMIT COALESCE($3::int, 50)
`

type SearchAlbumsParams struct {
	Search pgtype.Text `json:"search"`
	Start  int32       `json:"start"`
	Size   int32       `json:"size"`
}

func (q *Queries) SearchAlbums(ctx context.Context, arg SearchAlbumsParams) ([]Album, error) {
	rows, err := q.db.Query(ctx, searchAlbums, arg.Search, arg.Start, arg.Size)
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
			&i.Name,
			&i.Thumbnail,
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

const updateAlbum = `-- name: UpdateAlbum :one
UPDATE albums SET
    name = $2,
    artist_id = $3,
    thumbnail = $4,
    release_date = $5
WHERE id = $1 RETURNING id, artist_id, name, thumbnail, release_date, created_at
`

type UpdateAlbumParams struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	ArtistID    int32     `json:"artist_id"`
	Thumbnail   string    `json:"thumbnail"`
	ReleaseDate time.Time `json:"release_date"`
}

func (q *Queries) UpdateAlbum(ctx context.Context, arg UpdateAlbumParams) (Album, error) {
	row := q.db.QueryRow(ctx, updateAlbum,
		arg.ID,
		arg.Name,
		arg.ArtistID,
		arg.Thumbnail,
		arg.ReleaseDate,
	)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.ArtistID,
		&i.Name,
		&i.Thumbnail,
		&i.ReleaseDate,
		&i.CreatedAt,
	)
	return i, err
}
