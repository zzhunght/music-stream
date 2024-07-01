package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"

	"github.com/gin-gonic/gin"
)

func (r *Router) setUpPlaylistRouter(route *gin.RouterGroup) {
	albumService := services.NewAlbumServices(r.store)
	albumHandler := controller.NewAlbumController(albumService)

	playListRoute := route.Group("/album")
	{
		playListRoute.GET("/new", albumHandler.GetNewAlbum)
		playListRoute.GET("/song/:album_id", albumHandler.GetAlbumSongs)
	}
}
