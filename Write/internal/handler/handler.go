package handler

import (
	"github/SLANGERES/CQRS/Write/internal/models"

	"github/SLANGERES/CQRS/Write/internal/repository"
	"github/SLANGERES/CQRS/Write/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	repo repository.BlogRepository
}

func NewBlogHandler(repo repository.BlogRepository) *BlogHandler {
	return &BlogHandler{repo: repo}
}

func (h *BlogHandler) AddBlog(c *gin.Context) {
	var blog models.Blog

	// Bind + validate
	if err := c.BindJSON(&blog); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON")
		return
	}
	//addup other detail
	blog = util.NewReqbody(blog)
	if err := util.ValidateReqBody(blog); err != nil {
		util.ErrorResponse(c, http.StatusUnprocessableEntity, "Validation failed")
		return
	}

	// Call repository
	if _, err := h.repo.CreateBlog(c.Request.Context(), blog); err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, "Failed to save blog")
		return
	}

	util.OkResponse(c, "Blog added successfully")
}

func (h *BlogHandler) Health(c *gin.Context) {
	util.OkResponse(c, "Health is ok")
}
