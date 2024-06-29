package router

import (
	"music-app-backend/internal/app/helper"
	"music-app-backend/message"
	db "music-app-backend/sqlc"
	"music-app-backend/worker"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupRouter(
	route *gin.Engine,
	store db.SQLStore,
	task_client *worker.DeliveryTaskClient,
	messageQueue *message.RabbitMQProvider,
	rdb *redis.Client,
	tokenMaker *helper.Token,
) {
	v1 := route.Group("/api/v1")
	v1.GET("/health-check", healthCheck)
	setUpSongRouter(v1, store, messageQueue)
	setUpAlbumRouter(v1, store)
	setUpArtistRouter(v1, store)
	setUpPlaylistRouter(v1, store)
	setUpCategoriesRouter(v1, store)
}
func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
