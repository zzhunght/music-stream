package api

import api "music-app-backend/internal/app/api/middleware"

func (s *Server) UserRouter() {

	user := s.router.Group("/user")
	{
		user.POST("/register", s.Register)
		user.POST("/verify-otp", s.VerifyOTP)
		user.POST("/send-otp", s.ResendOTP)
		user.POST("/login", s.Login)
		user.Use(api.Authentication(s.token_maker))
		user.GET("/", s.GetUser)
	}

}
