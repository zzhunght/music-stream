package controller

import (
	"errors"
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/response"
	"music-app-backend/internal/app/services"
	"music-app-backend/pkg/middleware"
	db "music-app-backend/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateComment struct {
	Content string `json:"content" binding:"required"`
	SongID  int32  `json:"song_id" binding:"required"`
}

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentController(services *services.CommentService) *CommentController {
	return &CommentController{commentService: services}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {

	// authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(middleware.AuthenticationPayload)
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*helper.TokenPayload)

	var body CreateComment

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	payload := db.CreateCommentParams{
		UserID:  authPayload.UserID,
		Content: body.Content,
		SongID:  body.SongID,
	}

	data, err := c.commentService.CreateComment(ctx, payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse(data, "Create comment successfully"))
}

func (c *CommentController) GetCommentsBySong(ctx *gin.Context) {

	id, ok := ctx.Params.Get("song_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(errors.New("provide song_id")))
		return
	}
	songID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	data, err := c.commentService.GetCommentsBySongID(ctx, songID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(data, "Get comments successfully"))
}
