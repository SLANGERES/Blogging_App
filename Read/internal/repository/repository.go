package repository

import (
	"bytes"
	"encoding/json"
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
func (h *BlogRepo) GetBlogByID(id string) ([]models.Blog, error) {
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
	// Ask Elasticsearch for unique values of "tags"
	query := `{
		"size": 0,
		"aggs": {
			"unique_categories": {
				"terms": { "field": "tags.keyword", "size": 1000 }
			}
		}
	}`

	// Run search
	res, err := h.db.DB.Search(
		h.db.DB.Search.WithIndex("blogs"),
		h.db.DB.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Response structure we care about
	var result struct {
		Aggregations struct {
			UniqueCategories struct {
				Buckets []struct {
					Key string `json:"key"`
				} `json:"buckets"`
			} `json:"unique_categories"`
		} `json:"aggregations"`
	}

	// Decode response
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Collect all category names
	var categories []string
	for _, b := range result.Aggregations.UniqueCategories.Buckets {
		categories = append(categories, b.Key)
	}

	return categories, nil
}

func (h *BlogRepo) GetBlogByCategory(categories string) ([]models.Blog, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"terms": map[string]interface{}{
				"tags.keyword": categories,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("encode query: %w", err)
	}
	res, err := h.db.DB.Search(
		h.db.DB.Search.WithIndex("blogs"),
		h.db.DB.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, fmt.Errorf("search error: %w", err)
	}
	defer res.Body.Close()

	return utils.DecodeBlogsResponse(res)
}

func (h *BlogRepo) GetFullSearch(term string) ([]models.Blog, error) {
	// Build full-text search query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  term,
				"fields": []string{"title", "content", "tags"},
			},
		},
	}

	// Marshal query to JSON
	body, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("encode query: %w", err)
	}

	// Execute search
	res, err := h.db.DB.Search(
		h.db.DB.Search.WithIndex("blogs"),
		h.db.DB.Search.WithBody(bytes.NewReader(body)),
	)
	if err != nil {
		return nil, fmt.Errorf("search error: %w", err)
	}
	defer res.Body.Close()

	// Decode hits into []models.Blog
	return utils.DecodeBlogsResponse(res)
}

func (h *BlogRepo) GetTopBlogByLikes(limit int) ([]models.Blog, error) {
	// Build query: fetch top N blogs sorted by likes desc
	query := map[string]interface{}{
		"size": limit,
		"sort": []map[string]interface{}{
			{"likes": map[string]string{"order": "desc"}},
		},
	}

	// Marshal query
	body, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("encode query: %w", err)
	}

	// Execute search
	res, err := h.db.DB.Search(
		h.db.DB.Search.WithIndex("blogs"),
		h.db.DB.Search.WithBody(bytes.NewReader(body)),
	)
	if err != nil {
		return nil, fmt.Errorf("search error: %w", err)
	}
	defer res.Body.Close()

	// Decode into []models.Blog
	return utils.DecodeBlogsResponse(res)
}
