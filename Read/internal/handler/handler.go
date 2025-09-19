package handler

import (
	"fmt"
	"github/SLANGERES/CQRS/Read/internal/broker"
	"github/SLANGERES/CQRS/Read/internal/repository"
	"github/SLANGERES/CQRS/Read/internal/utils"
	"log/slog"
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

	slog.Info("get blog by ID api/blog/id its ", "blog id ", reqId)

	req := utils.ConfigReq(c)

	data, err := h.repo.GetBlogByID(reqId)

	slog.Info("get blog by ID api/blog/id its ", "blog->", data)

	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
	}
	req.OkRespose(data)
}

func (h *BlogHandler) GetBlogByTag(c *gin.Context) {
	tags := c.Query("tags")

	slog.Info("get blog by ID api/blog/?tags its ", "tag-> ", tags)

	req := utils.ConfigReq(c)
	data, err := h.repo.GetBlogByTag(tags)
	slog.Info("get blog by ID api/blog/?tags its ", "blog with tags ", data)
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

	// Log the requested category
	slog.Info("Finding blogs with category", "category-name", category)

	// Validate input
	if category == "" {
		req.ErrorResponse(http.StatusBadRequest, fmt.Errorf("category query param is required"))
		return
	}

	// Fetch blogs from repo
	data, err := h.repo.GetBlogByCategory(category)
	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	// Log the result
	slog.Info("Blogs found for category", "blogs", data)

	// Send successful response
	req.OkRespose(data)
}

func (h *BlogHandler) GetFullSearch(c *gin.Context) {
	search := c.Query("search")
	req := utils.ConfigReq(c)

	// Log the incoming search query
	slog.Info("Search requested", "query", search)

	// Validate input
	if search == "" {
		req.ErrorResponse(http.StatusBadRequest, fmt.Errorf("search query param is required"))
		return
	}

	// Perform search via repository
	data, err := h.repo.GetFullSearch(search)
	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	// Log results
	slog.Info("Search results", "data", data)

	// Return successful response
	req.OkRespose(data)
}


func (h *BlogHandler) GetTopBlog(c *gin.Context) {
	queryLimit := c.Query("limit")
	req := utils.ConfigReq(c)

	// Set default limit if not provided
	limit := 10
	var err error
	if queryLimit != "" {
		limit, err = strconv.Atoi(queryLimit)
		if err != nil || limit <= 0 {
			req.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid limit value"))
			return
		}
	}

	slog.Info("Fetching top blogs", "limit", limit)

	data, err := h.repo.GetTopBlogByLikes(limit)
	if err != nil {
		req.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	req.OkRespose(data)
}
