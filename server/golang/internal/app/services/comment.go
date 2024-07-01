package services

import (
	"context"
	db "music-app-backend/sqlc"
)

type CommentService struct {
	store *db.SQLStore
}

func NewCommentService(store *db.SQLStore) *CommentService {
	return &CommentService{store: store}
}

func (s *CommentService) GetCommentByID(ctx context.Context, id int32) (db.Comment, error) {
	return s.store.GetCommentById(ctx, id)
}
func (s *CommentService) CreateComment(ctx context.Context, body db.CreateCommentParams) (db.Comment, error) {
	return s.store.CreateComment(ctx, body)
}

func (s *CommentService) GetCommentsBySongID(ctx context.Context, song_id int) ([]db.GetSongCommentRow, error) {
	return s.store.GetSongComment(ctx, int32(song_id))
}

func (s *CommentService) DeleteComment(ctx context.Context, id int32) (err error) {
	return s.store.DeleteComment(ctx, id)
}
