package api

import (
	"fmt"
	"music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateArtistRequest struct {
	Name      string `json:"name" binding:"required"`
	AvatarUrl string `json:"avatar_url" binding:"required"`
}

type UpdateArtistRequest struct {
	Name      string `json:"name" binding:"required"`
	AvatarUrl string `json:"avatar_url" binding:"required"`
}

type ArtistResponse struct {
	Data  []sqlc.Artist `json:"data"`
	Count int32         `json:"count"`
}

func (s *Server) GetArtists(c *gin.Context) {
	// fmt.Println("x user id :>>>>>>>>>>>>>>>>>>>>>>", c.GetHeader("x-user-id"))
	// fmt.Println("x email id :>>>>>>>>>>>>>>>>>>>>>>", c.GetHeader("x-user-email"))
	// fmt.Println("x user role :>>>>>>>>>>>>>>>>>>>>>>", c.GetHeader("x-user-role"))
	for k, vals := range c.Request.Header {
		fmt.Printf("%s", k)
		for _, v := range vals {
			fmt.Printf("\t%s", v)
		}
	}
	seach := c.DefaultQuery("search", "")
	// page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// size, _ := strconv.Atoi(c.DefaultQuery("size", "50"))

	artist, err := s.store.GetListArtists(c,
		pgtype.Text{
			String: seach,
			Valid:  true,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	count, err := s.store.CountListArtists(c, pgtype.Text{
		String: seach,
		Valid:  true,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	data := ArtistResponse{
		Data:  artist,
		Count: int32(count),
	}
	c.JSON(http.StatusOK, SuccessResponse(data, "Get artists"))
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

func (s *Server) UpdateArtist(c *gin.Context) {
	artist_id, _ := c.Params.Get("artist_id")
	id, err := strconv.Atoi(artist_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	var body UpdateArtistRequest
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid body"})
		return
	}

	new_art, err := s.store.UpdateArtist(c, sqlc.UpdateArtistParams{
		ID:   int32(id),
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

func (s *Server) DeleteArtist(c *gin.Context) {
	artist_id, _ := c.Params.Get("artist_id")
	id, err := strconv.Atoi(artist_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	err = s.store.DeleteArtist(c, int32(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid body"})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse(1, "Xóa nghệ sĩ thành công"))
}
