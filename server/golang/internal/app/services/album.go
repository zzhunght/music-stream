package services

import (
	"context"
	db "music-app-backend/sqlc"
)

type AlbumService struct {
	store *db.SQLStore
}

func NewAlbumServices(store *db.SQLStore) *AlbumService {
	return &AlbumService{
		store: store,
	}
}

func (s *AlbumService) GetNewAlbums(ctx context.Context) ([]db.Album, error) {
	return s.store.GetNewAlbum(ctx)
}

func (s *AlbumService) GetAlbumSongs(ctx context.Context, album_id int32) ([]db.GetAlbumSongRow, error) {
	return s.store.GetAlbumSong(ctx, album_id)
}
