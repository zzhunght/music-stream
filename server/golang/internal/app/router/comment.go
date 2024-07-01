package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"
	"music-app-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func (r *Router) setUpCommentRoute(route *gin.RouterGroup) {
	commentService := services.NewCommentService(r.store)
	commentHandler := controller.NewCommentController(commentService)

	commentRoute := route.Group("/comment")
	{
		commentRoute.GET("/:song_id", commentHandler.GetCommentsBySong)

		commentRoute.Use(middleware.Authentication(r.tokenMaker))
		commentRoute.POST("/", commentHandler.CreateComment)
	}
}
