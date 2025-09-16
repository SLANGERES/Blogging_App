package router

import (
	"github/SLANGERES/CQRS/Write/internal/handler"

	"github.com/gin-gonic/gin"
)

func New(blogHandler *handler.BlogHandler) *gin.Engine {
	r := gin.Default()

	// Route mappings
	r.GET("/health", blogHandler.Health)
	r.POST("/blogs", blogHandler.AddBlog)

	return r
}
