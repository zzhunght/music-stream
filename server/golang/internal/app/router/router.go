package router

import (
	db "music-app-backend/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine, store db.SQLStore) {
	v1 := route.Group("/api/v1")
	v1.GET("/health-check", healthCheck)
	setUpSongRouter(v1, store)
	setUpAlbumRouter(v1, store)
	setUpArtistRouter(v1, store)
	setUpPlaylistRouter(v1, store)
}
func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
