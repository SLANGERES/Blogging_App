package handler

import (
	"github/SLANGERES/CQRS/Read/internal/broker"
	"github/SLANGERES/CQRS/Read/internal/repository"
	"github/SLANGERES/CQRS/Read/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	repo   repository.BlogRepo
	mqConn *broker.Consumer
}

func NewBlogHandler(repo repository.BlogRepo, mqConn *broker.Consumer) BlogHandler {
	return BlogHandler{
		repo:   repo,
		mqConn: mqConn,
	}
}
func (h *BlogHandler) GetAllBlog(c *gin.Context) {
	req := utils.ConfigReq(c)
	data, err := h.repo.GetAllBlog()

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)
}

func (h *BlogHandler) GetBlogByID(c *gin.Context) {
	reqId := c.Param("id")
	req := utils.ConfigReq(c)
	data, err := h.repo.GetBlogByID(reqId)

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)
}

func (h *BlogHandler) GetBlogByTag(c *gin.Context) {
	tags := c.Query("tags")
	req := utils.ConfigReq(c)
	data, err := h.repo.GetBlogByID(tags)

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)
}

func (h *BlogHandler) GetAllCategory(c *gin.Context) {
	req := utils.ConfigReq(c)
	data, err := h.repo.GetAllCategory()

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)
}

func (h *BlogHandler) GetBlogByCategory(c *gin.Context) {
	category := c.Query("category")
	req := utils.ConfigReq(c)
	data, err := h.repo.GetBlogByCategory(category)

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)
}

func (h *BlogHandler) GetFullSearch(c *gin.Context) {
	search := c.Query("serch")
	req := utils.ConfigReq(c)
	data, err := h.repo.GetFullSearch(search)

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)
}

func (h *BlogHandler) GetTopBlog(c *gin.Context) {
	querylimit := c.Query("limit")
	req := utils.ConfigReq(c)
	limit, err := strconv.Atoi(querylimit)
	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}

	data, err := h.repo.GetTopBlogByLikes(limit)

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)

}
