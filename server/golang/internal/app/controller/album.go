package controller

import (
	"errors"
	"music-app-backend/internal/app/response"
	"music-app-backend/internal/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumService *services.AlbumService
}

func NewAlbumController(services *services.AlbumService) *AlbumController {
	return &AlbumController{
		albumService: services,
	}
}

func (c *AlbumController) GetNewAlbum(ctx *gin.Context) {

	albums, err := c.albumService.GetNewAlbums(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(albums, "Danh sách album mới nhất"))
	return
}

func (c *AlbumController) GetAlbumSongs(ctx *gin.Context) {

	id, ok := ctx.Params.Get("album_id")

	if !ok {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(errors.New("provide album_id")))
		return
	}
	album_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	albums, err := c.albumService.GetAlbumSongs(ctx, int32(album_id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(albums, "Danh sách album mới nhất"))
	return
}
