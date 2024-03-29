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
		//  admin artists
		admin.GET("/artists", s.GetArtists)
		admin.POST("/artists/", s.CreateArtist)
		admin.PUT("/artists/:artist_id", s.UpdateArtist)
		admin.DELETE("/artists/:artist_id", s.DeleteArtist)

		// admin songs
		admin.GET("/categories", s.GetCategories)
		admin.POST("/categories", s.CreateCategory)
		admin.PUT("/categories", s.UpdateCategory)
		admin.DELETE("/categories/:category_id", s.DeleteCategory)

		// admin categories

		// song associated
		admin.POST("/associate-song-artist", s.AssociateSongArtist)
		admin.POST("/remove-associate-song-artist", s.RemoveAssociateSongArtist)
	}
}

func (s *Server) PublicRouter() {

	public := s.router.Group("/public")
	{
		public.GET("/artists", s.GetArtists)
		public.GET("/categories", s.GetCategories)
		public.GET("/songs", s.SearchSong)
		public.GET("/songs_by_categories/:categories_id", s.GetSongByCategories)
	}
}
