// repository/blog_repo.go
package repository

import (
	"context"
	"github/SLANGERES/CQRS/Write/database"
	"github/SLANGERES/CQRS/Write/internal/models"
)

// concrete implementation
type blogRepo struct {
	db database.Database // <- depend on abstraction, not concrete pgx
}

// constructor
func NewBlogRepository(db database.Database) BlogRepository {
	return &blogRepo{db: db}
}

func (r *blogRepo) CreateBlog(ctx context.Context, blog models.Blog) (string, error) {

	_, err := r.db.Exec(ctx,
		"INSERT INTO blogs (id, title, description, content, tags, category) VALUES ($1, $2, $3, $4, $5, $6)",
		blog.ID, blog.Title, blog.Description, blog.Content, blog.Tags, blog.Metadata.Category,
	)
	if err != nil {
		return "", err
	}
	return blog.ID, nil
}

func (r *blogRepo) UpdateBlog(ctx context.Context, blog models.Blog) (string, error) {
	_, err := r.db.Exec(ctx,
		"UPDATE blogs SET title=$1, content=$2, description=$3, tags=$4, category=$5, updated_at = now() WHERE id=$6",
		blog.Title, blog.Content, blog.Description, blog.Tags, blog.Metadata.Category, blog.ID,
	)
	if err != nil {
		return "", err
	}
	return blog.ID, nil
}

func (r *blogRepo) DeleteBlog(ctx context.Context, id string) (string, error) {
	_, err := r.db.Exec(ctx,
		"DELETE FROM blogs WHERE id=$1",
		id,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}
