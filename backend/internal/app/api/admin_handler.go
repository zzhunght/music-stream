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

type AssociateSongArtistRequest struct {
	ArtistID int32 `json:"artist_id" binding:"required"`
	SongID   int32 `json:"song_id" binding:"required"`
	Owner    bool  `json:"owner" binding:"required"`
}

type RemoveAssociateSongArtistRequest struct {
	ArtistID int32 `json:"artist_id" binding:"required"`
	SongID   int32 `json:"song_id" binding:"required"`
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

func (s *Server) AssociateSongArtist(c *gin.Context) {

	var body AssociateSongArtistRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid body"})
		return
	}

	fmt.Println("err :", err)
	fmt.Println("data :", body)
	err = s.store.AssociateSongArtist(c, sqlc.AssociateSongArtistParams{
		SongID:   body.SongID,
		ArtistID: body.ArtistID,
		Owner:    body.Owner,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse(true, "Cập nhật nghệ sĩ cho bài hát thành công"))
}

func (s *Server) RemoveAssociateSongArtist(c *gin.Context) {

	var body AssociateSongArtistRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid body"})
		return
	}

	fmt.Println("data :", body)
	err = s.store.RemoveAssociateSongArtist(c, sqlc.RemoveAssociateSongArtistParams{
		SongID:   body.SongID,
		ArtistID: body.ArtistID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse(true, "Cập nhật nghệ sĩ cho bài hát thành công"))
}
