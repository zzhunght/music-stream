package sqlc

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type UpateSongWithTx struct {
	UpdateSongBody UpdateSongParams
	ArtistID       int32
	CategoryID     int32
}

type CreateSongWithTxParams struct {
	CreateSongBody CreateSongParams
	ArtistID       int32
	CategoryID     int32
	AfterFunction  func([]byte) error
}

type MessageData struct {
	Artist    Artist `json:"artist"`
	SongID    int32  `json:"song_id"`
	SongName  string `json:"song_name"`
	Type      string `json:"type"`
	Thumbnail string `json:"thumbnail"`
}
type MessageBody struct {
	Event string      `json:"event"`
	Data  MessageData `json:"data"`
}

func (store *SQLStore) CreateSongWithTx(ctx context.Context, arg CreateSongWithTxParams) (GetSongByIDRow, error) {
	tx, err := store.connPool.Begin(ctx)

	if err != nil {
		log.Info().Msg("Can not begin transaction")
	}
	defer tx.Rollback(ctx)
	// tạo 1 interface với transaction
	qtx := store.WithTx(tx)
	song, err := qtx.CreateSong(ctx, arg.CreateSongBody)

	if err != nil {
		return GetSongByIDRow{}, err
	}
	err = qtx.AssociateSongArtist(ctx, AssociateSongArtistParams{
		SongID:   song.ID,
		ArtistID: arg.ArtistID,
	})
	if err != nil {
		return GetSongByIDRow{}, err
	}

	err = qtx.AddSongToCategory(ctx, AddSongToCategoryParams{
		SongID:     song.ID,
		CategoryID: arg.CategoryID,
	})
	if err != nil {
		return GetSongByIDRow{}, err
	}

	// -----------------------------
	artist, err := qtx.GetArtistById(ctx, arg.ArtistID)
	if err != nil {
		return GetSongByIDRow{}, err
	}

	body, err := json.Marshal(MessageBody{
		Event: "CREATE_NEW_SONG",
		Data: MessageData{
			SongID:    song.ID,
			Type:      "CREATE",
			SongName:  song.Name,
			Artist:    artist,
			Thumbnail: song.Thumbnail.String,
		},
	})
	if err != nil {
		return GetSongByIDRow{}, err
	}
	err = arg.AfterFunction(body)
	if err != nil {
		return GetSongByIDRow{}, err
	}
	song_return, err := qtx.GetSongByID(ctx, song.ID)
	if err != nil {
		return song_return, err
	}
	tx.Commit(ctx)
	return song_return, nil
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
	err = qtx.UpdateSongCategory(ctx, UpdateSongCategoryParams{
		SongID:     arg.UpdateSongBody.ID,
		CategoryID: arg.CategoryID,
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
