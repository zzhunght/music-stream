package api

import (
	"fmt"
	"music-app-backend/internal/app/utils"
	"music-app-backend/sqlc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateSongRequest struct {
	Name        string    `json:"name" binding:"required"`
	Thumbnail   string    `json:"thumbnail" binding:"required"`
	Path        string    `json:"path" binding:"required"`
	Lyrics      string    `json:"lyrics" binding:"required"`
	Duration    int32     `json:"duration" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
}

func (s *Server) GetSong(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "50"))

	fmt.Print(page, size)

	songs, err := s.store.GetSongs(ctx, sqlc.GetSongsParams{
		Size:  int32(size),
		Start: int32(size) * int32(page-1),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(songs, "Get songs successfully"))
}

func (s *Server) CreateSong(ctx *gin.Context) {

	var body CreateSongRequest

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	new_song, err := s.store.CreateSong(ctx, sqlc.CreateSongParams{
		Name:      body.Name,
		Path:      utils.ConvertStringToText(body.Path),
		Thumbnail: utils.ConvertStringToText(body.Thumbnail),
		Lyrics:    utils.ConvertStringToText(body.Lyrics),
		Duration: pgtype.Int4{
			Int32: body.Duration,
			Valid: true,
		},
		ReleaseDate: pgtype.Date{
			Time:  body.ReleaseDate,
			Valid: true,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(new_song, "Created song successfully"))
}
