package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"
	db "music-app-backend/sqlc"

	"github.com/gin-gonic/gin"
)

func setUpPlaylistRouter(r *gin.RouterGroup, store db.SQLStore) {
	albumService := services.NewAlbumServices(store)
	albumHandler := controller.NewAlbumController(albumService)

	route := r.Group("/album")
	{
		route.GET("/new", albumHandler.GetNewAlbum)
	}
}
