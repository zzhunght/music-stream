package api

import (
	"fmt"
	"music-app-backend/sqlc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AlbumResponse struct {
	Data  []sqlc.Album `json:"data"`
	Count int64        `json:"count"`
}

type RemoveSongFromAlbum struct {
	Ids []int32 `json:"ids" binding:"required"`
}

type UpdateAlbumParams struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	ArtistID    int32     `json:"artist_id"`
	Thumbnail   string    `json:"thumbnail"`
	ReleaseDate time.Time `json:"release_date"`
}

func (s *Server) GetAlbums(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "50"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	album, err := s.store.GetAlbums(ctx, sqlc.GetAlbumsParams{
		Start: (int32(page) - 1) * int32(size),
		Size:  int32(size),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	count, err := s.store.CountAlbums(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	data := AlbumResponse{
		Data:  album,
		Count: count,
	}
	ctx.JSON(http.StatusOK, SuccessResponse(data, "Danh sách album"))

}
func (s *Server) GetAlbumSong(ctx *gin.Context) {
	album_id, _ := ctx.Params.Get("album_id")
	id, err := strconv.Atoi(album_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	data, err := s.store.GetAlbumSong(ctx, int32(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(data, "Danh sách bài hát trong album"))

}

func (s *Server) GetAlbumByArtistId(ctx *gin.Context) {
	artist_id, _ := ctx.Params.Get("artist_id")
	id, err := strconv.Atoi(artist_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	data, err := s.store.GetAlbumByArtistID(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	count, err := s.store.CountAlbumsByArtistID(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	resp := AlbumResponse{
		Data:  data,
		Count: count,
	}

	ctx.JSON(http.StatusOK, SuccessResponse(resp, "Danh sách album theo nghệ sĩ"))

}

func (s *Server) CreateAlbum(ctx *gin.Context) {
	var body sqlc.CreateAlbumParams

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	data, err := s.store.CreateAlbum(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(data, "Tạo Album thành công"))
}

func (s *Server) AddSongToAlbum(ctx *gin.Context) {

	var body sqlc.AddManySongToAlbumParams

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	err = s.store.AddManySongToAlbum(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, "Thêm bài hát vào album thành công"))
}

func (s *Server) RemoveSongFromAlbum(ctx *gin.Context) {

	var body RemoveSongFromAlbum
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	fmt.Print("IDS", body.Ids)
	err = s.store.RemoveSongFromAlbum(ctx, body.Ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, "Xóa bài hát khỏi album thành công"))
}

func (s *Server) UpdateAlbum(ctx *gin.Context) {
	album_id, b := ctx.Params.Get("album_id")
	id, err := strconv.Atoi(album_id)

	var body UpdateAlbumParams
	if !b {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	err = ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	data, err := s.store.UpdateAlbum(ctx, sqlc.UpdateAlbumParams{
		ID:          int32(id),
		Name:        body.Name,
		ReleaseDate: body.ReleaseDate,
		ArtistID:    body.ArtistID,
		Thumbnail:   body.Thumbnail,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(data, "Cập nhật album thành công"))

}
func (s *Server) DeleteAlbum(ctx *gin.Context) {
	album_id, b := ctx.Params.Get("album_id")
	id, err := strconv.Atoi(album_id)

	if !b {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	err = s.store.DeleteAlbum(ctx, int32(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(true, "Xóa album thành công"))

}
