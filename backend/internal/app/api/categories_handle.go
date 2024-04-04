package api

import (
	"music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateCategory struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategory struct {
	Name string `json:"name" binding:"required"`
	ID   int    `json:"id" binding:"required"`
}

func (s *Server) GetCategories(ctx *gin.Context) {

	categories, err := s.store.GetSongCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(categories, "Danh sách danh mục"))
}

func (s *Server) CreateCategory(ctx *gin.Context) {
	var body CreateCategory

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	data, err := s.store.CreateCategories(ctx, body.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(data, "Create category successfully"))

}

func (s *Server) UpdateCategory(ctx *gin.Context) {
	var body UpdateCategory

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	data, err := s.store.UpdateCategories(ctx, sqlc.UpdateCategoriesParams{
		ID:   int32(body.ID),
		Name: body.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(data, "Update category successfully"))

}

func (s *Server) DeleteCategory(ctx *gin.Context) {
	categories_id, _ := ctx.Params.Get("category_id")
	id, err := strconv.Atoi(categories_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	err = s.store.DeleteCategories(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, "Delete category successfully"))

}
