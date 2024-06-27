package response

import "github.com/gin-gonic/gin"

func SuccessResponse(data interface{}, message string) gin.H {
	return gin.H{
		"data":    data,
		"message": message,
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
