package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, statusCode int, errorMsg string) {
	c.JSON(statusCode, gin.H{
		"error": errorMsg,
	})
}

func OkResponse(c *gin.Context, sucessMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"message": sucessMsg,
	})
}
