package controller

import (
	"music-app-backend/internal/app/response"
	"music-app-backend/internal/app/services"
	"net/http"

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
