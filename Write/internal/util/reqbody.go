package util

import (
	"github/SLANGERES/CQRS/Write/internal/models"
	"time"

	"github.com/google/uuid"
)

func NewReqbody(blog models.Blog) models.Blog {
	now := time.Now()
	blog.ID = uuid.New().String()
	blog.CreatedAt = now
	blog.UpdatedAt = now
	blog.Metadata.Likes = 0
	return blog
}
