package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssociateSongArtistRequest struct {
	ArtistID int32 `json:"artist_id" binding:"required"`
	SongID   int32 `json:"song_id" binding:"required"`
	Owner    bool  `json:"owner" binding:"required"`
}

type RemoveAssociateSongArtistRequest struct {
	ArtistID int32 `json:"artist_id" binding:"required"`
	SongID   int32 `json:"song_id" binding:"required"`
}

func (s *Server) GetStatics(ctx *gin.Context) {

	data, err := s.store.Statistics(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(data, "Get statics"))
}
