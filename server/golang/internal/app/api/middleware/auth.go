package api

import (
	"music-app-backend/internal/app/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationKey        = "authorization"
	authorizationType       = "Bearer"
	AuthorizationPayloadKey = "TokenPayload"
)

func Authentication(tokenMaker *helper.Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header",
			})
			return
		}

		authorizationData := strings.Fields(authorizationHeader)

		if len(authorizationData) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header type",
			})
			return
		}

		if authorizationData[0] != authorizationType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header type",
			})
			return
		}

		payload, err := tokenMaker.VerifyToken(authorizationData[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
