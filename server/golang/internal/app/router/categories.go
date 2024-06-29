package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"
	db "music-app-backend/sqlc"

	"github.com/gin-gonic/gin"
)

func setUpCategoriesRouter(r *gin.RouterGroup, store db.SQLStore) {

	categoriesService := services.NewCategoriesService(store)
	categoriesHandler := controller.NewCategoriesController(categoriesService)

	route := r.Group("/artist")
	{
		route.GET("/categories", categoriesHandler.GetCategories)
		route.POST("/categories", categoriesHandler.CreateCategory)
		route.PUT("/categories", categoriesHandler.UpdateCategory)
		route.DELETE("/categories/:category_id", categoriesHandler.DeleteCategory)
	}
}
