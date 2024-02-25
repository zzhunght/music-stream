package api

import (
	"music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) GetArtists(c *gin.Context) {
	seach := c.DefaultQuery("search", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "50"))

	artist, err := s.store.GetListArtists(c, sqlc.GetListArtistsParams{
		Start: (int32(page) - 1) * int32(size),
		Size:  int32(size),
		NameSearch: pgtype.Text{
			String: seach,
			Valid:  true,
		},
		OrderBy: "name ASC",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(artist, "Get artists"))
}
