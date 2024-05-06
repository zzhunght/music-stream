package sqlc

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type UpateSongWithTx struct {
	UpdateSongBody UpdateSongParams
	ArtistID       int32
}

type CreateSongWithTxParams struct {
	CreateSongBody CreateSongParams
	ArtistID       int32
	AfterFunction  func([]byte) error
}

type MessageData struct {
	ArtistID int32  `json:"artist_id"`
	SongID   int32  `json:"song_id"`
	Type     string `json:"type"`
}
type MessageBody struct {
	Event string      `json:"event"`
	Data  MessageData `json:"data"`
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
		Event: "CREATE_NEW_SONG",
		Data: MessageData{
			ArtistID: arg.ArtistID,
			SongID:   song.ID,
			Type:     "CREATE",
		},
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

func (store *SQLStore) UpateSongWithTx(ctx context.Context, arg UpateSongWithTx) (GetSongByIDRow, error) {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		log.Info().Msg("Can not begin transaction")
	}
	defer tx.Rollback(ctx)
	qtx := store.WithTx(tx)
	err = qtx.UpdateSong(ctx, arg.UpdateSongBody)
	if err != nil {
		return GetSongByIDRow{}, err
	}
	err = qtx.UpdateAssociateSongArtist(ctx, UpdateAssociateSongArtistParams{
		SongID:   arg.UpdateSongBody.ID,
		ArtistID: arg.ArtistID,
	})
	if err != nil {
		return GetSongByIDRow{}, err
	}
	song, err := qtx.GetSongByID(ctx, arg.UpdateSongBody.ID)
	if err != nil {
		return song, err
	}
	tx.Commit(ctx)
	return song, nil
}
