package repository

import "github/SLANGERES/CQRS/Read/internal/models"

type NewBlogRepository interface {

	GetAllBlog() ([]models.Blog, error)
	
	GetBlogByID() (models.Blog, error)

	GetBlogByTag() ([]models.Blog, error)

	GetAllCategory() ([]string, error)

	GetBlogByCategory() ([]models.Blog, error)

	GetFullSearch() ([]models.Blog, error)

	GetTopBlog() ([]models.Blog, error)
}
