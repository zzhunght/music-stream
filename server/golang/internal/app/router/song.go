package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"
	"music-app-backend/message"
	"music-app-backend/sqlc"

	"github.com/gin-gonic/gin"
)

func setUpSongRouter(r *gin.RouterGroup, store sqlc.SQLStore, messageQueue *message.RabbitMQProvider) {
	services := services.NewSongService(store)
	handler := controller.NewSongController(services, messageQueue)

	song_routes := r.Group("/song")
	{
		song_routes.GET("/new-song", handler.GetNewsSong)
		song_routes.GET("/admin", handler.AdminGetSong)
		song_routes.POST("/", handler.CreateSong)
		song_routes.PUT("/:song_id", handler.UpdateSong)
		song_routes.DELETE("/:song_id", handler.DeleteSong)
	}
}
