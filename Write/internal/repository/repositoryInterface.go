// repository/blog_repository.go
package repository

import (
	"context"
	"github/SLANGERES/CQRS/Write/internal/models"
)

// Interface (abstraction)
type BlogRepository interface {
	
	CreateBlog(ctx context.Context, blog models.Blog) (string, error)
	UpdateBlog(ctx context.Context, blog models.Blog) (string, error)
	DeleteBlog(ctx context.Context, id string) (string, error)
}
