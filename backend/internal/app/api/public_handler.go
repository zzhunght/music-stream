package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) GetCategories(ctx *gin.Context) {

	categories, err := s.store.GetSongCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(categories, "Danh sách danh mục"))
}

func (s *Server) SearchSong(ctx *gin.Context) {

	search := ctx.DefaultQuery("search", "")
	fmt.Println("query : >>>>>>>>>> ", search)
	songs, err := s.store.SearchSong(ctx, pgtype.Text{
		String: search,
		Valid:  true,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(songs, "Tìm kiếm bài hát thành công"))
}
