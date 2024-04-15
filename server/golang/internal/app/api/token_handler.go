package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RenewTokenBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (s *Server) RenewToken(ctx *gin.Context) {

	var body RenewTokenBody
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("token not found")))
		return
	}

	token_payload, err := s.token_maker.VerifyToken(body.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}
	session, err := s.store.GetSession(ctx, token_payload.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("Unauthorized")))
		return
	}
	if session.ID != token_payload.ID {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("Unauthorized")))
		return
	}
	if session.Email != token_payload.Email {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("Unauthorized")))
		return
	}

	if session.RefreshToken != body.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("Unauthorized")))
		return
	}

	acc, err := s.store.GetAccount(ctx, session.Email)

	if session.RefreshToken != body.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("Unauthorized")))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("Unauthorized")))
		return
	}

	new_access_token, _, err := s.token_maker.CreateToken(acc.Email, acc.ID, acc.Role, s.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("Unauthorized")))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse(new_access_token, "Access token created"))
}
