// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	RoleID    int32            `json:"role_id"`
	SecretKey pgtype.Text      `json:"secret_key"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type Album struct {
	ID          int32            `json:"id"`
	ArtistID    pgtype.Int4      `json:"artist_id"`
	Name        pgtype.Text      `json:"name"`
	ReleaseDate interface{}      `json:"release_date"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type AlbumsSong struct {
	ID        int32            `json:"id"`
	SongID    pgtype.Int4      `json:"song_id"`
	AlbumID   pgtype.Int4      `json:"album_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Artist struct {
	ID        int32            `json:"id"`
	Name      pgtype.Text      `json:"name"`
	AvatarUrl pgtype.Text      `json:"avatar_url"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type ArtistFollow struct {
	ID        int32            `json:"id"`
	AccountID pgtype.Int4      `json:"account_id"`
	ArtistID  pgtype.Int4      `json:"artist_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Category struct {
	ID        int32            `json:"id"`
	Name      pgtype.Text      `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type FavoriteAlbum struct {
	ID        int32            `json:"id"`
	AccountID pgtype.Int4      `json:"account_id"`
	AlbumID   pgtype.Int4      `json:"album_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type FavoriteSong struct {
	ID        int32            `json:"id"`
	AccountID pgtype.Int4      `json:"account_id"`
	SongID    pgtype.Int4      `json:"song_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Playlist struct {
	ID          int32            `json:"id"`
	Name        pgtype.Int4      `json:"name"`
	AccountID   pgtype.Int4      `json:"account_id"`
	Description pgtype.Text      `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type PlaylistSong struct {
	ID         int32            `json:"id"`
	PlaylistID pgtype.Int4      `json:"playlist_id"`
	SongID     pgtype.Int4      `json:"song_id"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type Role struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type Song struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Thumbnail   pgtype.Text      `json:"thumbnail"`
	Path        pgtype.Text      `json:"path"`
	Lyrics      pgtype.Text      `json:"lyrics"`
	Duration    pgtype.Int4      `json:"duration"`
	ReleaseDate interface{}      `json:"release_date"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type SongCategory struct {
	ID         int32            `json:"id"`
	SongID     pgtype.Int4      `json:"song_id"`
	CategoryID pgtype.Int4      `json:"category_id"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type SongsArtist struct {
	ID       int32       `json:"id"`
	SongID   pgtype.Int4 `json:"song_id"`
	ArtistID pgtype.Int4 `json:"artist_id"`
	Owner    pgtype.Bool `json:"owner"`
}
