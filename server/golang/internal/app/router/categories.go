package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"

	"github.com/gin-gonic/gin"
)

func (r *Router) setUpCategoriesRouter(route *gin.RouterGroup) {

	categoriesService := services.NewCategoriesService(r.store)
	categoriesHandler := controller.NewCategoriesController(categoriesService)

	categoriesRoute := route.Group("/categories")
	{
		categoriesRoute.GET("/", categoriesHandler.GetCategories)
		categoriesRoute.POST("", categoriesHandler.CreateCategory)
		categoriesRoute.PUT("/", categoriesHandler.UpdateCategory)
		categoriesRoute.DELETE("/:category_id", categoriesHandler.DeleteCategory)
	}
}
