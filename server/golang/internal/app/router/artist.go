package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"
	db "music-app-backend/sqlc"

	"github.com/gin-gonic/gin"
)

func setUpArtistRouter(r *gin.RouterGroup, store db.SQLStore) {

	artistService := services.NewArtistService(store)
	artistHandler := controller.NewArtistController(artistService)

	route := r.Group("/artist")
	{
		route.GET("/recommendations", artistHandler.GetRecommendArtist)
		route.GET("/song/:artist_id", artistHandler.GetArtistSong)
	}
}
