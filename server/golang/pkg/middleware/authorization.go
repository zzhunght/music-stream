package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Role = "x-user-role"
)

func Authorization(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		role := ctx.GetHeader(Role)

		if len(role) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Forbidden",
			})
			return
		}

		for _, r := range roles {
			if r == role {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Forbidden",
		})
	}
}
