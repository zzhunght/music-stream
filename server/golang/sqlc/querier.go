// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddManySongToAlbum(ctx context.Context, arg AddManySongToAlbumParams) error
	AddSongToAlbum(ctx context.Context, arg AddSongToAlbumParams) error
	AddSongToCategory(ctx context.Context, arg AddSongToCategoryParams) error
	AddSongToPlaylist(ctx context.Context, arg AddSongToPlaylistParams) error
	AdminGetSongs(ctx context.Context) ([]AdminGetSongsRow, error)
	AssociateSongArtist(ctx context.Context, arg AssociateSongArtistParams) error
	ChangePassword(ctx context.Context, arg ChangePasswordParams) error
	CheckEmailExists(ctx context.Context, email string) (int32, error)
	CheckOwnerPlaylist(ctx context.Context, arg CheckOwnerPlaylistParams) (CheckOwnerPlaylistRow, error)
	CheckSongInPlaylist(ctx context.Context, arg CheckSongInPlaylistParams) (int64, error)
	CountAlbums(ctx context.Context) (int64, error)
	CountAlbumsByArtistID(ctx context.Context, artistID int32) (int64, error)
	CountListArtists(ctx context.Context, nameSearch pgtype.Text) (int64, error)
	CountSearchAlbums(ctx context.Context, search pgtype.Text) (int64, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error)
	CreateAlbum(ctx context.Context, arg CreateAlbumParams) (Album, error)
	CreateArtist(ctx context.Context, arg CreateArtistParams) (Artist, error)
	CreateCategories(ctx context.Context, name string) (Category, error)
	CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (Playlist, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSong(ctx context.Context, arg CreateSongParams) (Song, error)
	DeleteAlbum(ctx context.Context, id int32) error
	DeleteArtist(ctx context.Context, id int32) error
	DeleteCategories(ctx context.Context, id int32) error
	DeleteManyArtist(ctx context.Context, ids []int32) error
	DeletePlaylist(ctx context.Context, arg DeletePlaylistParams) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteSong(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, email string) (GetAccountRow, error)
	GetAlbumByArtistID(ctx context.Context, artistID int32) ([]Album, error)
	GetAlbumInfoFromSongID(ctx context.Context, songID int32) (GetAlbumInfoFromSongIDRow, error)
	GetAlbumSong(ctx context.Context, albumID int32) ([]GetAlbumSongRow, error)
	GetAlbums(ctx context.Context, arg GetAlbumsParams) ([]Album, error)
	GetArtistById(ctx context.Context, id int32) (Artist, error)
	GetLatestAlbum(ctx context.Context) ([]GetLatestAlbumRow, error)
	GetListArtists(ctx context.Context, nameSearch pgtype.Text) ([]Artist, error)
	GetPlaylistofUser(ctx context.Context, accountID int32) ([]Playlist, error)
	GetRandomSong(ctx context.Context) ([]GetRandomSongRow, error)
	GetSecretKey(ctx context.Context, email string) (pgtype.Text, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSongByID(ctx context.Context, id int32) (GetSongByIDRow, error)
	GetSongById(ctx context.Context, id int32) (GetSongByIdRow, error)
	GetSongBySongCategory(ctx context.Context, arg GetSongBySongCategoryParams) ([]GetSongBySongCategoryRow, error)
	GetSongCategories(ctx context.Context) ([]Category, error)
	GetSongInPlaylist(ctx context.Context, playlistID int32) ([]Song, error)
	GetSongNotInAlbum(ctx context.Context, arg GetSongNotInAlbumParams) ([]GetSongNotInAlbumRow, error)
	GetSongOfArtist(ctx context.Context, arg GetSongOfArtistParams) ([]GetSongOfArtistRow, error)
	GetSongs(ctx context.Context, arg GetSongsParams) ([]GetSongsRow, error)
	RemoveAssociateSongArtist(ctx context.Context, arg RemoveAssociateSongArtistParams) error
	RemoveSongFromAlbum(ctx context.Context, arg RemoveSongFromAlbumParams) error
	RemoveSongFromPlaylist(ctx context.Context, arg RemoveSongFromPlaylistParams) error
	SearchAlbums(ctx context.Context, arg SearchAlbumsParams) ([]SearchAlbumsRow, error)
	SearchSong(ctx context.Context, arg SearchSongParams) ([]SearchSongRow, error)
	Statistics(ctx context.Context) (StatisticsRow, error)
	UpdateAlbum(ctx context.Context, arg UpdateAlbumParams) (Album, error)
	UpdateArtist(ctx context.Context, arg UpdateArtistParams) (Artist, error)
	UpdateAssociateSongArtist(ctx context.Context, arg UpdateAssociateSongArtistParams) error
	UpdateCategories(ctx context.Context, arg UpdateCategoriesParams) (Category, error)
	UpdatePlaylist(ctx context.Context, arg UpdatePlaylistParams) (Playlist, error)
	UpdateSong(ctx context.Context, arg UpdateSongParams) error
	UpdateSongCategory(ctx context.Context, arg UpdateSongCategoryParams) error
	VerifyAccount(ctx context.Context, email string) (Account, error)
}

var _ Querier = (*Queries)(nil)
