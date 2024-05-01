package api

import (
	api "music-app-backend/internal/app/api/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) UserRouter(route *gin.RouterGroup) {

	user := route.Group("/user")
	{
		user.GET("/confirm-forget-password", s.ForgetPasswordConfirm)
		user.POST("/register", s.Register)
		user.POST("/verify-otp", s.VerifyOTP)
		user.POST("/send-otp", s.ResendOTP)
		user.POST("/login", s.Login)
		user.POST("/refresh-token", s.RenewToken)
		user.POST("/forget-password", s.ForgetPasswordRequest)
		user.Use(api.Authentication(s.token_maker))

		user.GET("/playlist", s.GetUserPlaylists)
		user.GET("/playlist/:playlist_id/song", s.GetPlaylistSong)
		user.POST("/playlist", s.CreatePlaylist)
		user.POST("/playlist/add-song/:playlist_id", s.AddSongToPlaylist)
		user.POST("/playlist/remove-song/:playlist_id", s.RemoveSongFromPlaylist)
		user.PUT("/playlist/:playlist_id", s.UpdatePlaylistName)
		user.DELETE("/playlist/:playlist_id", s.DeletePlaylist)
		user.POST("/change-password", s.ChangePassword)
		user.GET("/info", s.GetUser)
	}

}

func (s *Server) AdminRouter(route *gin.RouterGroup) {
	admin := route.Group("/admin")
	{
		//  admin artists
		admin.GET("/artists", s.GetArtists)
		admin.GET("/artists/album/:artist_id", s.GetAlbumByArtistId)
		admin.POST("/artists", s.CreateArtist)
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
		admin.PUT("/album/:album_id", s.UpdateAlbum)
		admin.DELETE("/album/:album_id", s.DeleteAlbum)

		// song associated
		admin.POST("/associate-song-artist", s.AssociateSongArtist)
		admin.POST("/remove-associate-song-artist", s.RemoveAssociateSongArtist)
	}
}

func (s *Server) PublicRouter(route *gin.RouterGroup) {

	public := route.Group("/public")
	{
		public.GET("/artists", s.GetArtists)
		public.GET("/album/latest", s.GetLatestAlbum)
		public.GET("/album/:album_id/songs", s.GetAlbumSong)
		public.GET("/categories", s.GetCategories)
		public.GET("/songs", s.SearchSong)
		public.GET("/songs/suggested", s.RandomSong)
		public.GET("/songs_by_categories/:categories_id", s.GetSongByCategories)
		public.GET("/search", s.Search)
		public.GET("/album", s.GetAlbums)
		public.GET("/album/:album_id", s.GetAlbumSong)
		public.GET("/artists/:artist_id/songs", s.GetSongOfArtist)
		public.GET("/artists/album/:artist_id", s.GetAlbumByArtistId)
	}
}
