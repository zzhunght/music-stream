package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"

	"github.com/gin-gonic/gin"
)

func (r *Router) setUpArtistRouter(route *gin.RouterGroup) {

	artistService := services.NewArtistService(r.store)
	artistHandler := controller.NewArtistController(artistService)

	artistRoute := route.Group("/artist")
	{
		artistRoute.GET("/recommendations", artistHandler.GetRecommendArtist)
		artistRoute.GET("/song/:artist_id", artistHandler.GetArtistSong)
	}
}
