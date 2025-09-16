package handler

import (
	"github/SLANGERES/CQRS/Read/internal/repository"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	repo repository.BlogRepo
}

func NewBlogHandler(repo repository.BlogRepo) BlogHandler {
	return BlogHandler{
		repo: repo,
	}
}
func (h *BlogHandler) GetAllBlog(c *gin.Context) {
	h.repo.GetAllBlog()
}

func (h *BlogHandler) GetBlogByID(c *gin.Context) {

}

func (h *BlogHandler) GetBlogByTag(c *gin.Context) {

}

func (h *BlogHandler) GetAllCategory(c *gin.Context) {

}

func (h *BlogHandler) GetBlogByCategory(c *gin.Context) {

}

func (h *BlogHandler) GetFullSearch(c *gin.Context) {

}

func (h *BlogHandler) GetTopBlog(c *gin.Context) {

}
