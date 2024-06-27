package services

import (
	"context"
	db "music-app-backend/sqlc"
)

type ArtistService struct {
	store db.SQLStore
}

func NewArtistService(store db.SQLStore) *ArtistService {
	return &ArtistService{
		store: store,
	}
}

func (s *ArtistService) RecommendedArtist(ctx context.Context) ([]db.Artist, error) {
	return s.store.GetRecommentArtist(ctx)
}

func (s *ArtistService) GetArtistSong(ctx context.Context, artist_id int) ([]db.Song, error) {
	return s.store.GetSongOfArtist(ctx, int32(artist_id))
}

func (s *ArtistService) GetArtistById(ctx context.Context, artist_id int) (db.Artist, error) {
	return s.store.GetArtistById(ctx, int32(artist_id))
}
