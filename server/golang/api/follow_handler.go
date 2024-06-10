package api

import (
	"errors"
	"music-app-backend/pkg/middleware"
	db "music-app-backend/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Follow struct {
	ArtistID int32 `json:"artist_id" binding:"required"`
}

func (s *Server) UnfollowArtist(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(middleware.AuthenticationPayload)

	var body Follow

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("invalid body")))
		return
	}

	_, err = s.store.CheckFollow(ctx, db.CheckFollowParams{
		ArtistID:  int32(body.ArtistID),
		AccountID: int32(authPayload.UserID),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	err = s.store.UnFollow(ctx, db.UnFollowParams{
		ArtistID:  int32(body.ArtistID),
		AccountID: int32(authPayload.UserID),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(true, "Un follow success"))
}

func (s *Server) FollowArtist(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(middleware.AuthenticationPayload)

	var body Follow

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("invalid body")))
		return
	}
	check, _ := s.store.CheckFollow(ctx, db.CheckFollowParams{
		ArtistID:  int32(body.ArtistID),
		AccountID: int32(authPayload.UserID),
	})
	if check != (db.ArtistFollow{}) {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("already follow")))
		return
	}

	err = s.store.Follow(ctx, db.FollowParams{
		ArtistID:  int32(body.ArtistID),
		AccountID: int32(authPayload.UserID),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(true, "Follow success"))
}
