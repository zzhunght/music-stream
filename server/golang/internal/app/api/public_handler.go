package api

import (
	"fmt"
	db "music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type SearchResult struct {
	Song   []db.SearchSongRow `json:"songs"`
	Artist []db.Artist        `json:"artists"`
	Album  []db.Album         `json:"albums"`
}

func (s *Server) SearchSong(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "50"))
	search := ctx.DefaultQuery("search", "")
	fmt.Println("query : >>>>>>>>>> ", search)
	songs, err := s.store.SearchSong(ctx, db.SearchSongParams{
		Size:  int32(size),
		Start: (int32(page) - 1) * int32(size),
		Search: pgtype.Text{
			String: search,
			Valid:  true,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(songs, "Tìm kiếm bài hát thành công"))
}
func (s *Server) RandomSong(ctx *gin.Context) {
	songs, err := s.store.GetRandomSong(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse(songs, "Danh sách bài hát ngẫu nhiên"))
}
func (s *Server) GetSongByCategories(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "50"))
	categories_id, _ := ctx.Params.Get("categories_id")
	id, err := strconv.Atoi(categories_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	songs, err := s.store.GetSongBySongCategory(ctx, db.GetSongBySongCategoryParams{
		CategoryID: int32(id),
		Size:       int32(size),
		Start:      (int32(page) - 1) * int32(size),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse(songs, "Tìm kiếm bài hát thành công"))
}

func (s *Server) GetSongOfArtist(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "50"))
	artist_id, _ := ctx.Params.Get("artist_id")
	id, err := strconv.Atoi(artist_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	data, err := s.store.GetSongOfArtist(ctx, db.GetSongOfArtistParams{
		ID:    int32(id),
		Size:  int32(size),
		Start: (int32(page) - 1) * int32(size),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(data, "Danh sách bài hát theo nghệ sĩ"))

}

func (s *Server) GetLatestAlbum(ctx *gin.Context) {

	data, err := s.store.GetLatestAlbum(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(data, "Danh sách album mới nhất"))
}

func (s *Server) Search(ctx *gin.Context) {

	search := ctx.DefaultQuery("search", "")
	fmt.Println("query : >>>>>>>>>> ", search)
	songs, err := s.store.SearchSong(ctx, db.SearchSongParams{
		Size:  int32(10),
		Start: 0,
		Search: pgtype.Text{
			String: search,
			Valid:  true,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	album, err := s.store.SearchAlbums(ctx, db.SearchAlbumsParams{
		Size:  int32(3),
		Start: 0,
		Search: pgtype.Text{
			String: search,
			Valid:  true,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	artist, err := s.store.GetListArtists(ctx, pgtype.Text{
		String: search,
		Valid:  true,
	})

	if err != nil {
		fmt.Println("error getting artists: ", err)
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	response := SearchResult{
		Song:   songs,
		Artist: artist,
		Album:  album,
	}
	ctx.JSON(http.StatusOK, SuccessResponse(response, "Tìm kiếm bài hát thành công"))
}
