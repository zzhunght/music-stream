package middleware

import (
	"music-app-backend/internal/app/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticationPayload struct {
	UserID int32  `json:"user_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
}

const (
	authorizationKey        = "authorization"
	authorizationType       = "Bearer"
	AuthorizationPayloadKey = "AuthPayload"
	UserIDKey               = "x-user-id"
	EmailKey                = "x-user-email"
	RoleKey                 = "x-user-role"
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

// func Authentication() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		user_id := ctx.GetHeader(UserIDKey)
// 		if len(user_id) == 0 {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "Unauthorized",
// 			})
// 			return
// 		}
// 		id, err := strconv.Atoi(user_id)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"error": "Internal server error",
// 			})
// 			return
// 		}

// 		role := ctx.GetHeader(RoleKey)
// 		if len(role) == 0 {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "Unauthorized",
// 			})
// 			return
// 		}
// 		email := ctx.GetHeader(EmailKey)
// 		if len(email) == 0 {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "Unauthorized",
// 			})
// 			return
// 		}

// 		fmt.Print("x-user-id: ", user_id)
// 		fmt.Print("x-user-email: ", email)
// 		fmt.Print("x-user-role: ", role)
// 		ctx.Set(AuthorizationPayloadKey, AuthenticationPayload{
// 			Email:  email,
// 			UserID: int32(id),
// 			Role:   role,
// 		})
// 		ctx.Next()
// 	}
// }
