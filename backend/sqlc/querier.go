// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddSongToAlbum(ctx context.Context, arg AddSongToAlbumParams) error
	AssociateSongArtist(ctx context.Context, arg AssociateSongArtistParams) error
	CheckEmailExists(ctx context.Context, email string) (int32, error)
	CountAlbums(ctx context.Context) (int64, error)
	CountAlbumsByArtistID(ctx context.Context, artistID int32) (int64, error)
	CountSearchAlbums(ctx context.Context, search pgtype.Text) (int64, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error)
	CreateAlbum(ctx context.Context, arg CreateAlbumParams) (Album, error)
	CreateArtist(ctx context.Context, arg CreateArtistParams) (Artist, error)
	CreateCategories(ctx context.Context, name string) (Category, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSong(ctx context.Context, arg CreateSongParams) (Song, error)
	DeleteAlbum(ctx context.Context, id int32) error
	DeleteArtist(ctx context.Context, id int32) error
	DeleteCategories(ctx context.Context, id int32) error
	DeleteManyArtist(ctx context.Context, ids []int32) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteSong(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, email string) (GetAccountRow, error)
	GetAlbumByArtistID(ctx context.Context, artistID int32) (Album, error)
	GetAlbumSong(ctx context.Context, id int32) ([]Song, error)
	GetAlbums(ctx context.Context, arg GetAlbumsParams) ([]Album, error)
	GetListArtists(ctx context.Context, arg GetListArtistsParams) ([]Artist, error)
	GetRandomSong(ctx context.Context) ([]Song, error)
	GetSecretKey(ctx context.Context, email string) (pgtype.Text, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSongBySongCategory(ctx context.Context, arg GetSongBySongCategoryParams) ([]Song, error)
	GetSongCategories(ctx context.Context) ([]Category, error)
	GetSongs(ctx context.Context, arg GetSongsParams) ([]Song, error)
	RemoveAssociateSongArtist(ctx context.Context, arg RemoveAssociateSongArtistParams) error
	RemoveSongToAlbum(ctx context.Context, arg RemoveSongToAlbumParams) error
	SearchSong(ctx context.Context, search pgtype.Text) ([]Song, error)
	UpdateAlbum(ctx context.Context, arg UpdateAlbumParams) error
	UpdateArtist(ctx context.Context, arg UpdateArtistParams) (Artist, error)
	UpdateCategories(ctx context.Context, arg UpdateCategoriesParams) (Category, error)
	UpdateSong(ctx context.Context, arg UpdateSongParams) (Song, error)
	VerifyAccount(ctx context.Context, email string) (Account, error)
}

var _ Querier = (*Queries)(nil)
