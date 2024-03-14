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
	CheckEmailExists(ctx context.Context, email string) (int32, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error)
	CreateArtist(ctx context.Context, arg CreateArtistParams) (Artist, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSong(ctx context.Context, arg CreateSongParams) (Song, error)
	CreateSongCategories(ctx context.Context, name string) (Category, error)
	DeleteArtist(ctx context.Context, ids []int32) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	GetAccount(ctx context.Context, email string) (GetAccountRow, error)
	GetListArtists(ctx context.Context, arg GetListArtistsParams) ([]Artist, error)
	GetRandomSong(ctx context.Context) ([]Song, error)
	GetSecretKey(ctx context.Context, email string) (pgtype.Text, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSongCategories(ctx context.Context) ([]Category, error)
	SearchSong(ctx context.Context, search pgtype.Text) ([]Song, error)
	UpdateArtist(ctx context.Context, arg UpdateArtistParams) (Artist, error)
}

var _ Querier = (*Queries)(nil)