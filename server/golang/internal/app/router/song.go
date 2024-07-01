package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"

	"github.com/gin-gonic/gin"
)

func (r *Router) setUpSongRouter(route *gin.RouterGroup) {
	services := services.NewSongService(r.store)
	handler := controller.NewSongController(services, r.messageQueue)

	song_routes := route.Group("/song")
	{
		song_routes.GET("/new-song", handler.GetNewsSong)
		song_routes.GET("/admin", handler.AdminGetSong)
		song_routes.POST("/", handler.CreateSong)
		song_routes.PUT("/:song_id", handler.UpdateSong)
		song_routes.DELETE("/:song_id", handler.DeleteSong)
	}
}
