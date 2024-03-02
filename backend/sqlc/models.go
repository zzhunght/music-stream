// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"github.com/google/uuid"
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
	ArtistID    int32            `json:"artist_id"`
	Name        string           `json:"name"`
	ReleaseDate pgtype.Date      `json:"release_date"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type AlbumsSong struct {
	ID        int32            `json:"id"`
	SongID    int32            `json:"song_id"`
	AlbumID   int32            `json:"album_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Artist struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	AvatarUrl pgtype.Text      `json:"avatar_url"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type ArtistFollow struct {
	ID        int32            `json:"id"`
	AccountID int32            `json:"account_id"`
	ArtistID  int32            `json:"artist_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Category struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type FavoriteAlbum struct {
	ID        int32            `json:"id"`
	AccountID int32            `json:"account_id"`
	AlbumID   int32            `json:"album_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type FavoriteSong struct {
	ID        int32            `json:"id"`
	AccountID int32            `json:"account_id"`
	SongID    int32            `json:"song_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Playlist struct {
	ID          int32            `json:"id"`
	Name        int32            `json:"name"`
	AccountID   int32            `json:"account_id"`
	Description pgtype.Text      `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type PlaylistSong struct {
	ID         int32            `json:"id"`
	PlaylistID int32            `json:"playlist_id"`
	SongID     int32            `json:"song_id"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type Role struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type Session struct {
	ID           uuid.UUID        `json:"id"`
	Email        string           `json:"email"`
	ClientAgent  string           `json:"client_agent"`
	RefreshToken string           `json:"refresh_token"`
	ClientIp     string           `json:"client_ip"`
	IsBlock      pgtype.Bool      `json:"is_block"`
	ExpiredAt    pgtype.Timestamp `json:"expired_at"`
}

type Song struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Thumbnail   pgtype.Text      `json:"thumbnail"`
	Path        pgtype.Text      `json:"path"`
	Lyrics      pgtype.Text      `json:"lyrics"`
	Duration    pgtype.Int4      `json:"duration"`
	ReleaseDate pgtype.Date      `json:"release_date"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type SongCategory struct {
	ID         int32            `json:"id"`
	SongID     int32            `json:"song_id"`
	CategoryID int32            `json:"category_id"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type SongsArtist struct {
	ID       int32       `json:"id"`
	SongID   int32       `json:"song_id"`
	ArtistID int32       `json:"artist_id"`
	Owner    pgtype.Bool `json:"owner"`
}
