package api

import (
	"fmt"
	"music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) SearchSong(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "50"))
	search := ctx.DefaultQuery("search", "")
	fmt.Println("query : >>>>>>>>>> ", search)
	songs, err := s.store.SearchSong(ctx, sqlc.SearchSongParams{
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

func (s *Server) GetSongByCategories(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "50"))
	categories_id, _ := ctx.Params.Get("categories_id")
	id, err := strconv.Atoi(categories_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	songs, err := s.store.GetSongBySongCategory(ctx, sqlc.GetSongBySongCategoryParams{
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
