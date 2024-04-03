package sqlc

import (
	"context"

	"github.com/rs/zerolog/log"
)

type CreateSongWithTxParams struct {
	CreateSongBody CreateSongParams
	ArtistID       int32
}

func (store *SQLStore) CreateSongWithTx(ctx context.Context, arg CreateSongWithTxParams) (Song, error) {
	tx, err := store.connPool.Begin(ctx)

	if err != nil {
		log.Info().Msg("Can not begin transaction")
	}
	defer tx.Rollback(ctx)
	qtx := store.WithTx(tx)
	song, err := qtx.CreateSong(ctx, arg.CreateSongBody)

	if err != nil {
		return song, err
	}
	err = qtx.AssociateSongArtist(ctx, AssociateSongArtistParams{
		SongID:   song.ID,
		ArtistID: arg.ArtistID,
	})
	if err != nil {
		return song, err
	}
	tx.Commit(ctx)
	return song, nil
}
