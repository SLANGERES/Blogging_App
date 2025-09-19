package handler

import (
	"github/SLANGERES/CQRS/Write/internal/broker"
	"github/SLANGERES/CQRS/Write/internal/models"
	"log/slog"

	"github/SLANGERES/CQRS/Write/internal/repository"
	"github/SLANGERES/CQRS/Write/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	repo   repository.BlogRepository
	mqconn *broker.MqBroker
}

func NewBlogHandler(repo repository.BlogRepository, mqconn *broker.MqBroker) *BlogHandler {
	return &BlogHandler{
		repo:   repo,
		mqconn: mqconn,
	}
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

	//Sendig to the mq
	newBlog := util.ConvertReqbodyMqBody(blog)

	if err := h.mqconn.Publish(newBlog); err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, "unable to sync maybe mq is broken")
	}
	slog.Info("Blog is sucess fully send to the rabbit mq ", "blog id", blog.ID)

	util.OkResponse(c, "Blog added successfully")
}

func (h *BlogHandler) Health(c *gin.Context) {
	util.OkResponse(c, "Health is ok")
}
