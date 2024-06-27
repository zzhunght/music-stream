package controller

import (
	"music-app-backend/internal/app/response"
	"music-app-backend/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	songService *services.SongService
}

func NewSongController(services *services.SongService) *SongController {
	return &SongController{
		songService: services,
	}
}

// func (c *SongController) GetSong(ctx *gin.Context) {
// 	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
// 	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "999"))

// 	fmt.Print(page, size)

// 	songs, err := c.songService.GetSongs(ctx, size, page)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response.SuccessResponse(songs, "Get songs successfully"))
// }

func (c *SongController) GetNewsSong(ctx *gin.Context) {
	songs, err := c.songService.GetNewSongs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(songs, "Get news songs successfully"))
}
