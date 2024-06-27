package services

import (
	"context"
	db "music-app-backend/sqlc"
)

type AlbumService struct {
	store db.SQLStore
}

func NewAlbumServices(store db.SQLStore) *AlbumService {
	return &AlbumService{
		store: store,
	}
}

func (s *AlbumService) GetNewAlbums(ctx context.Context) ([]db.Album, error) {
	return s.store.GetNewAlbum(ctx)
}
