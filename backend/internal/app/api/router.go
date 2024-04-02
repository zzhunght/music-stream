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
		admin.GET("/artists/album/:artist_id", s.GetAlbumByArtistId)
		admin.POST("/artists/", s.CreateArtist)
		admin.PUT("/artists/:artist_id", s.UpdateArtist)
		admin.DELETE("/artists/:artist_id", s.DeleteArtist)

		// admin categories
		admin.GET("/categories", s.GetCategories)
		admin.POST("/categories", s.CreateCategory)
		admin.PUT("/categories", s.UpdateCategory)
		admin.DELETE("/categories/:category_id", s.DeleteCategory)

		// admin song
		admin.GET("/song", s.GetSong)
		admin.POST("/song", s.CreateSong)
		admin.PUT("/song/:song_id", s.UpdateSong)
		admin.DELETE("/song/:song_id", s.DeleteSong)

		// album
		admin.GET("/album", s.GetAlbums)
		admin.GET("/album/:album_id", s.GetAlbumSong)
		admin.POST("/album", s.CreateAlbum)
		admin.POST("/album/add-song", s.AddSongToAlbum)
		admin.POST("/album/remove-song", s.RemoveSongFromAlbum)
		// admin.POST("/album", s.CreateAlbum)
		admin.PUT("/album/:album_id", s.UpdateAlbum)
		admin.DELETE("/album/:album_id", s.DeleteAlbum)

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
