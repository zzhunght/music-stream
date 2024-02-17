package routes

import (
	"music-app-backend/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.GET("/", handlers.ListUsers)
		user.POST("/register", handlers.Register)
		user.POST("/verify-otp", handlers.VerifyOTP)
		user.POST("/send-otp", handlers.ResendOTP)
		user.POST("/login", handlers.Register)
	}

}
