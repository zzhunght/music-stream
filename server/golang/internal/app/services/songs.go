package services

import (
	"context"
	db "music-app-backend/sqlc"
)

type SongService struct {
	store *db.SQLStore
}

func NewSongService(store *db.SQLStore) *SongService {
	return &SongService{
		store: store,
	}
}

func (s *SongService) GetSongByID(ctx context.Context, id int32) (song db.GetSongByIDRow, err error) {
	song, err = s.store.GetSongByID(ctx, id)
	return
}

func (s *SongService) GetNewSongs(ctx context.Context) ([]db.GetNewSongsRow, error) {
	songs, err := s.store.GetNewSongs(ctx)
	return songs, err
}

func (s *SongService) CreateSong(ctx context.Context, payload db.CreateSongWithTxParams) (db.GetSongByIDRow, error) {
	song, err := s.store.CreateSongWithTx(ctx, payload)
	return song, err
}

func (s *SongService) AdminGetSongs(ctx context.Context) ([]db.AdminGetSongsRow, error) {
	songs, err := s.store.AdminGetSongs(ctx)
	return songs, err
}

func (s *SongService) UpdateSong(ctx context.Context, payload db.UpateSongWithTx) (song db.GetSongByIDRow, err error) {
	song, err = s.store.UpateSongWithTx(ctx, payload)
	return
}

func (s *SongService) DeleteSong(ctx context.Context, id int32) (err error) {
	err = s.store.DeleteSong(ctx, id)
	return
}
