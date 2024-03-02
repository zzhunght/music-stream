package api

import api "music-app-backend/internal/app/api/middleware"

func (s *Server) UserRouter() {

	user := s.router.Group("/user")
	{
		user.POST("/register", s.Register)
		user.POST("/verify-otp", s.VerifyOTP)
		user.POST("/send-otp", s.ResendOTP)
		user.POST("/login", s.Login)
		user.POST("/refresh-token", s.RenewToken)
		user.Use(api.Authentication(s.token_maker))
		user.GET("/", s.GetUser)
	}

}

func (s *Server) AdminRouter() {
	admin := s.router.Group("/admin")
	{
		admin.POST("/artists", s.CreateArtist)
	}
}

func (s *Server) PublicRouter() {

	public := s.router.Group("/public")
	{
		public.GET("/artists", s.GetArtists)
		public.GET("/categories", s.GetCategories)
		public.GET("/songs", s.SearchSong)
	}
}
