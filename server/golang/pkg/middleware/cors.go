package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow requests from any origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow the Authorization header
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")

		// Allow GET, POST, OPTIONS methods
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// If it's an OPTIONS request, we're handling a preflight request.
		// So we don't need to execute the actual handler.
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// Call the next handler
		c.Next()
	}
}
