package services

import (
	"context"
	db "music-app-backend/sqlc"
)

type SongService struct {
	store db.SQLStore
}

func NewSongService(store db.SQLStore) *SongService {
	return &SongService{
		store: store,
	}
}

func (s *SongService) GetNewSongs(ctx context.Context) ([]db.GetNewSongsRow, error) {
	songs, err := s.store.GetNewSongs(ctx)
	return songs, err
}

func (s *SongService) CreateSong(ctx context.Context) ([]db.GetNewSongsRow, error) {
	songs, err := s.store.GetNewSongs(ctx)
	return songs, err
}
