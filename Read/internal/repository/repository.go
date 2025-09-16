package repository

import (
	"fmt"
	"github/SLANGERES/CQRS/Read/database"
	"github/SLANGERES/CQRS/Read/internal/models"
	"github/SLANGERES/CQRS/Read/internal/utils"
	"strings"
)

type BlogRepo struct {
	db database.Storage
}

func NewBlogRepo(db *database.Storage) BlogRepo {
	return BlogRepo{
		db: *db,
	}
}

func (h *BlogRepo) GetAllBlog() ([]models.Blog, error) {

	query := `{"query":{"match_all":{}}}`

	res, err := h.db.DB.Search(
		h.db.DB.Search.WithIndex("blogs"),
		h.db.DB.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	return utils.DecodeBlogsResponse(res)
}
func (h *BlogRepo) GetBlogByID() ([]models.Blog, error) {
	return nil, nil
}

func (h *BlogRepo) GetBlogByTag(tags []string) ([]models.Blog, error) {
	query := fmt.Sprintf(`{"query":{"term":{"tags":"%s"}}}`, tags)

	res, err := h.db.DB.Search(
		h.db.DB.Search.WithIndex("blogs"),
		h.db.DB.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	return utils.DecodeBlogsResponse(res)
}

func (h *BlogRepo) GetAllCategory() ([]string, error) {
	return nil, nil
}

func (h *BlogRepo) GetBlogByCategory() ([]models.Blog, error) {
	return nil, nil
}

func (h *BlogRepo) GetFullSearch() ([]models.Blog, error) {
	return nil, nil
}

func (h *BlogRepo) GetTopBlog() ([]models.Blog, error) {
	return nil, nil
}
