package util

import "github/SLANGERES/CQRS/Write/internal/models"

func ConvertReqbodyMqBody(blog models.Blog) models.ElasticBlog {
	return models.ElasticBlog{
		ID:          blog.ID,
		Title:       blog.Title,
		Description: blog.Description,
		Content:     blog.Content,
		Tags:        blog.Tags,
		Category:    blog.Metadata.Category,
		Likes:       blog.Metadata.Likes,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}
}
