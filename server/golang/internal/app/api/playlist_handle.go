package api

import (
	"errors"
	api "music-app-backend/internal/app/api/middleware"
	"music-app-backend/internal/app/helper"
	db "music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePlayListRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdatePlayListRequest struct {
	Name string `json:"name" binding:"required"`
}

type HanleSongPlayListRequest struct {
	SongID int32 `json:"song_id" binding:"required"`
}

func (s *Server) CreatePlaylist(ctx *gin.Context) {
	var body CreatePlayListRequest
	authPayload := ctx.MustGet(api.AuthorizationPayloadKey).(*helper.TokenPayload)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	data, err := s.store.CreatePlaylist(ctx, db.CreatePlaylistParams{
		Name:      body.Name,
		AccountID: int32(authPayload.UserID),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, SuccessResponse(data, "Tạo playlist thành công"))
}

func (s *Server) UpdatePlaylistName(ctx *gin.Context) {
	playlist_id, err := strconv.Atoi(ctx.Param("playlist_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(api.AuthorizationPayloadKey).(*helper.TokenPayload)
	var body UpdatePlayListRequest
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	data, err := s.store.UpdatePlaylist(ctx, db.UpdatePlaylistParams{
		Name:      body.Name,
		AccountID: int32(authPayload.UserID),
		ID:        int32(playlist_id),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, SuccessResponse(data, "Cập nhật thành công"))
}

func (s *Server) AddSongToPlaylist(ctx *gin.Context) {
	playlist_id, err := strconv.Atoi(ctx.Param("playlist_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(api.AuthorizationPayloadKey).(*helper.TokenPayload)
	var body HanleSongPlayListRequest
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	check, err := s.store.CheckOwnerPlaylist(ctx, db.CheckOwnerPlaylistParams{
		ID:        int32(playlist_id),
		AccountID: authPayload.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}

	exist, err := s.store.CheckSongInPlaylist(ctx, db.CheckSongInPlaylistParams{
		PlaylistID: check.ID,
		SongID:     body.SongID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	if exist > 0 {
		ctx.JSON(http.StatusCreated, SuccessResponse(true, "Bài hát đã tồn tại trong playlist"))
		return
	}

	err = s.store.AddSongToPlaylist(ctx, db.AddSongToPlaylistParams{
		PlaylistID: check.ID,
		SongID:     body.SongID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("Có lỗi trong quá trình thêm vào")))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, "Thêm bài hát vào playlist thành công"))
}

func (s *Server) RemoveSongFromPlaylist(ctx *gin.Context) {
	playlist_id, err := strconv.Atoi(ctx.Param("playlist_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(api.AuthorizationPayloadKey).(*helper.TokenPayload)
	var body HanleSongPlayListRequest
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	check, err := s.store.CheckOwnerPlaylist(ctx, db.CheckOwnerPlaylistParams{
		ID:        int32(playlist_id),
		AccountID: authPayload.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}

	err = s.store.RemoveSongFromPlaylist(ctx, db.RemoveSongFromPlaylistParams{
		PlaylistID: check.ID,
		SongID:     body.SongID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, "Xóa bài hát vào playlist thành công"))
}

func (s *Server) GetUserPlaylists(ctx *gin.Context) {
	authPayload := ctx.MustGet(api.AuthorizationPayloadKey).(*helper.TokenPayload)

	data, err := s.store.GetPlaylistofUser(ctx, authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, SuccessResponse(data, "playlist"))
}

func (s *Server) GetPlaylistSong(ctx *gin.Context) {

	playlist_id, err := strconv.Atoi(ctx.Param("playlist_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	data, err := s.store.GetSongInPlaylist(ctx, int32(playlist_id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, SuccessResponse(data, "playlist song"))
}

func (s *Server) DeletePlaylist(ctx *gin.Context) {
	playlist_id, err := strconv.Atoi(ctx.Param("playlist_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(api.AuthorizationPayloadKey).(*helper.TokenPayload)

	err = s.store.DeletePlaylist(ctx, db.DeletePlaylistParams{
		ID:        int32(playlist_id),
		AccountID: authPayload.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse(true, "Xóa playlist thành công"))
}
