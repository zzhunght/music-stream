package gapi

import (
	"context"
	"music-app-backend/pb"
	"music-app-backend/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type LoginResponse struct {
	SessionID    string       `json:"session_id"`
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

func (s *Server) Login(c context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	acc, _ := s.store.GetAccount(c, req.Email)

	validate := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.Password))
	if validate != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Username or password is incorrect")
	}
	access_token, _, err := s.token_maker.CreateToken(acc.Email, acc.ID, acc.Role, s.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	refresh_token, rf_payload, err := s.token_maker.CreateToken(acc.Email, acc.ID, acc.Role, s.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "internal error")
	}

	session, err := s.store.CreateSession(c, sqlc.CreateSessionParams{
		ID:           rf_payload.ID,
		Email:        rf_payload.Email,
		RefreshToken: refresh_token,
		ExpiredAt: pgtype.Timestamp{
			Time:  rf_payload.ExpiredAt,
			Valid: true,
		},
		ClientAgent: "",
		ClientIp:    "",
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "internal error")

	}

	resp := &pb.UserLoginResponse{
		SessionId: session.ID.String(),
		User: &pb.UserResponse{
			Email:     acc.Email,
			Name:      acc.Name,
			Id:        acc.ID,
			CreatedAt: timestamppb.New(acc.CreatedAt.Time),
			UpdatedAt: timestamppb.New(acc.UpdatedAt.Time),
		},
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	return resp, nil
}
