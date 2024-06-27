package controller

import (
	"errors"
	"music-app-backend/internal/app/response"
	"music-app-backend/internal/app/services"
	db "music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	artistService *services.ArtistService
}

func NewArtistController(service *services.ArtistService) *ArtistController {
	return &ArtistController{
		artistService: service,
	}
}

func (c *ArtistController) GetRecommendArtist(ctx *gin.Context) {
	artists, err := c.artistService.RecommendedArtist(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(artists, "Get recommended artists"))
}

func (c *ArtistController) GetArtistSong(ctx *gin.Context) {
	id, ok := ctx.Params.Get("artist_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(errors.New("provide artist_id")))
		return
	}
	artistID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	songs, _ := c.artistService.GetArtistSong(ctx, artistID)
	artist, err := c.artistService.GetArtistById(ctx, artistID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	data := struct {
		Songs  []db.Song `json:"songs"`
		Artist db.Artist `json:"artist"`
	}{
		Songs:  songs,
		Artist: artist,
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(data, "Get artist songs"))
}
