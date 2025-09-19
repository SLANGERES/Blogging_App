package repository

import "github/SLANGERES/CQRS/Read/internal/models"

type NewBlogRepository interface {

	GetAllBlog() ([]models.Blog, error)
	
	GetBlogByID(string) (models.Blog, error)

	GetBlogByTag(string) ([]models.Blog, error)

	GetAllCategory() ([]string, error)

	GetBlogByCategory(string) ([]models.Blog, error)

	GetFullSearch(string) ([]models.Blog, error)

	GetTopBlog(int) ([]models.Blog, error)
}
