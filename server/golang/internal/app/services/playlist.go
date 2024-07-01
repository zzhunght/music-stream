package services

import (
	"context"
	db "music-app-backend/sqlc"
)

type PlaylistService struct {
	store *db.SQLStore
}

func NewPlaylistService(store *db.SQLStore) *PlaylistService {
	return &PlaylistService{store: store}
}

func (s *PlaylistService) CreateUserPlaylist(ctx context.Context, payload db.CreateUserPlaylistParams) (db.Playlist, error) {

	return s.store.CreateUserPlaylist(ctx, payload)
}

func (s *PlaylistService) AddSongToPlaylist(ctx context.Context, payload db.AddSongToPlaylistParams) (db.GetSongByIdRow, error) {

	err := s.store.AddSongToPlaylist(ctx, payload)
	if err != nil {
		return db.GetSongByIdRow{}, err
	}

	data, err := s.store.GetSongById(ctx, payload.SongID)
	return data, err
}

func (s *PlaylistService) RemoveSongFromPlaylist(ctx context.Context, payload db.RemoveSongFromPlaylistParams) error {

	return s.store.RemoveSongFromPlaylist(ctx, payload)
}

func (s *PlaylistService) CheckOwnerPlaylist(ctx context.Context, payload db.CheckOwnerPlaylistParams) (db.CheckOwnerPlaylistRow, error) {

	return s.store.CheckOwnerPlaylist(ctx, payload)
}

func (s *PlaylistService) GetPlaylistSongs(ctx context.Context, playlist_id int32) ([]db.GetSongInPlaylistRow, error) {

	return s.store.GetSongInPlaylist(ctx, playlist_id)
}
