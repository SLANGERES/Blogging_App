package router

import (
	"github/SLANGERES/CQRS/Read/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(bloghandler handler.BlogHandler) *gin.Engine {

	router := gin.Default()

	router.GET("/api/blog", bloghandler.GetAllBlog)

	router.GET("/api/blog/:id", bloghandler.GetBlogByID)

	router.GET("/api/blog/by-tag", bloghandler.GetBlogByTag)

	router.GET("/api/blog/category", bloghandler.GetAllCategory)

	router.GET("/api/blog/by-category", bloghandler.GetBlogByCategory)

	router.GET("/api/blog/by-search", bloghandler.GetFullSearch)

	router.GET("/api/blog/top", bloghandler.GetTopBlog)

	return router
}
