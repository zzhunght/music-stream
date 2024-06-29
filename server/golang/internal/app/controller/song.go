package controller

import (
	"music-app-backend/internal/app/response"
	"music-app-backend/internal/app/services"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	db "music-app-backend/sqlc"
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
	Lyrics      string    `json:"lyrics"`
	Duration    int32     `json:"duration" binding:"required"`
	ReleaseDate time.Time `json:"release_date"`
	ArtistID    int32     `json:"artist_id" binding:"required"`
	CategoryID  int32     `json:"category_id" binding:"required"`
}

type UpdateSongRequest struct {
	Name        string    `json:"name" binding:"required"`
	Thumbnail   string    `json:"thumbnail" binding:"required"`
	Path        string    `json:"path" binding:"required"`
	Lyrics      string    `json:"lyrics"`
	Duration    int32     `json:"duration" binding:"required"`
	ReleaseDate time.Time `json:"release_date"`
	ArtistID    int32     `json:"artist_id" binding:"required"`
	CategoryID  int32     `json:"category_id" binding:"required"`
}

type SongController struct {
	songService  *services.SongService
	messageQueue *message.RabbitMQProvider
}

func NewSongController(services *services.SongService, messageQueue *message.RabbitMQProvider) *SongController {
	return &SongController{
		songService:  services,
		messageQueue: messageQueue,
	}
}

func (c *SongController) GetNewsSong(ctx *gin.Context) {
	songs, err := c.songService.GetNewSongs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(songs, "Get news songs successfully"))
}

func (c *SongController) AdminGetSong(ctx *gin.Context) {

	songs, err := c.songService.AdminGetSongs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(songs, "Get songs successfully"))
}

func (c *SongController) CreateSong(ctx *gin.Context) {

	var body CreateSongRequest

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	new_song, err := c.songService.CreateSong(ctx, db.CreateSongWithTxParams{
		CreateSongBody: db.CreateSongParams{
			Name:      body.Name,
			Path:      utils.ConvertStringToText(body.Path),
			Thumbnail: utils.ConvertStringToText(body.Thumbnail),
			Lyrics:    utils.ConvertStringToText(body.Lyrics),
			Duration: pgtype.Int4{
				Int32: body.Duration,
				Valid: true,
			},
			ReleaseDate: pgtype.Timestamp{
				Time:  body.ReleaseDate,
				Valid: true,
			},
			ArtistID: body.ArtistID,
		},
		CategoryID:    body.CategoryID,
		AfterFunction: c.messageQueue.Publishing,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse(new_song, "Created song successfully"))
}

func (c *SongController) UpdateSong(ctx *gin.Context) {
	song_id, err := strconv.Atoi(ctx.Param("song_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	var body UpdateSongRequest

	err = ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	new_song, err := c.songService.UpdateSong(ctx, db.UpateSongWithTx{
		UpdateSongBody: db.UpdateSongParams{
			ID:        int32(song_id),
			Name:      body.Name,
			Path:      utils.ConvertStringToText(body.Path),
			Thumbnail: utils.ConvertStringToText(body.Thumbnail),
			Lyrics:    utils.ConvertStringToText(body.Lyrics),
			Duration: pgtype.Int4{
				Int32: body.Duration,
				Valid: true,
			},
			ReleaseDate: pgtype.Timestamp{
				Time:  body.ReleaseDate,
				Valid: true,
			},
		},
		ArtistID:   body.ArtistID,
		CategoryID: body.CategoryID,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse(new_song, "Cập nhật bài hát thành công"))
}

func (c *SongController) DeleteSong(ctx *gin.Context) {
	song_id, err := strconv.Atoi(ctx.Param("song_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	err = c.songService.DeleteSong(ctx, int32(song_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse(true, "Xóa bài hát thành công"))
}
