package api

import (
	"fmt"
	"music-app-backend/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateArtistRequest struct {
	Name      string `json:"name" binding:"required"`
	AvatarUrl string `json:"avatar_url" binding:"required"`
}

func (s *Server) CreateArtist(c *gin.Context) {

	var body CreateArtistRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println("erorrrrrr :>>>>>>>>>>>>>")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid body"})
		return
	}

	fmt.Println("data :", body)
	new_art, err := s.store.CreateArtist(c, sqlc.CreateArtistParams{
		Name: body.Name,
		AvatarUrl: pgtype.Text{
			String: body.AvatarUrl,
			Valid:  true,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse(new_art, "Tạo nghệ sĩ thành công"))
}
