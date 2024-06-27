package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"
	"music-app-backend/sqlc"

	"github.com/gin-gonic/gin"
)

func setUpSongRouter(r *gin.RouterGroup, store sqlc.SQLStore) {
	services := services.NewSongService(store)
	handler := controller.NewSongController(services)

	song_routes := r.Group("/song")
	{
		song_routes.GET("/new-song", handler.GetNewsSong)
	}
}
