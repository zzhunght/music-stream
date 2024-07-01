package router

import (
	"music-app-backend/internal/app/controller"
	"music-app-backend/internal/app/services"
	"music-app-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func (r *Router) setUpUserRouter(route *gin.RouterGroup) {
	services := services.NewUserService(r.store)
	handler := controller.NewUserController(
		services,
		r.rdb,
		r.tokenMaker,
		r.task_client,
		r.config,
	)

	user := route.Group("/user")
	{
		user.GET("/confirm-forget-password", handler.ForgetPasswordConfirm)
		user.POST("/register", handler.Register)
		user.POST("/verify-otp", handler.VerifyOTP)
		user.POST("/send-otp", handler.ResendOTP)
		user.POST("/login", handler.Login)
		user.POST("/refresh-token", handler.RenewToken)
		user.POST("/forget-password", handler.ForgetPasswordRequest)
		// user.Use(middleware.Authentication(handler.token_maker))
		user.Use(middleware.Authentication(r.tokenMaker))

		//  play list
		user.POST("/change-password", handler.ChangePassword)
		user.GET("/info", handler.GetUser)
	}
}
