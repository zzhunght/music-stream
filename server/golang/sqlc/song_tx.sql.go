package sqlc

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type CreateSongWithTxParams struct {
	CreateSongBody CreateSongParams
	ArtistID       int32
	AfterFunction  func([]byte) error
}

type MessageBody struct {
	ArtistID int32 `json:"artist_id"`
	SongID   int32 `json:"song_id"`
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
	body, err := json.Marshal(MessageBody{
		ArtistID: arg.ArtistID,
		SongID:   song.ID,
	})
	if err != nil {
		return song, err
	}
	err = arg.AfterFunction(body)
	if err != nil {
		return song, err
	}
	tx.Commit(ctx)
	return song, nil
}
