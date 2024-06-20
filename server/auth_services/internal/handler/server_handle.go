package handler

import (
	"fmt"
	"music-app/authentication-services/internal/config"
	"music-app/authentication-services/internal/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	config      *config.Config
	token_maker *helper.Token
}

func NewHandler(config *config.Config,
	token_maker *helper.Token) *Handler {
	return &Handler{
		config:      config,
		token_maker: token_maker,
	}
}

func (h *Handler) Authentication(ctx *gin.Context) {

	header := ctx.GetHeader("authorization")
	fmt.Println("header :>>>>>>>>>>>>>>>>: ", header)

	if len(header) == 0 {
		ctx.Header("X-Error-Message", "Invalid authorization header")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization header",
		})
		return
	}
	split := strings.Split(header, " ")
	if len(split) != 2 {
		ctx.Header("X-Error-Message", "Invalid authorization header")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization header",
		})
		return
	}
	if split[0] != "Bearer" {
		ctx.Header("X-Error-Message", "Invalid authorization header")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization header",
		})
		return
	}
	token := split[1]
	payload, err := h.token_maker.VerifyToken(token)
	fmt.Println("error ,", err)

	if err != nil {
		ctx.Header("X-Error-Message", err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("x-user-id", strconv.Itoa(int(payload.UserID)))
	ctx.Header("x-user-email", payload.Email)
	ctx.Header("x-user-role", payload.Role)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
