package utils

import (
	"encoding/json"
	"github/SLANGERES/CQRS/Read/internal/models"

	"github.com/elastic/go-elasticsearch/v9/esapi"
)

func DecodeBlogsResponse(res *esapi.Response) ([]models.Blog, error) {
	defer res.Body.Close()

	var result struct {
		Hits struct {
			Hits []struct {
				Source models.Blog `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	blogs := make([]models.Blog, 0, len(result.Hits.Hits))
	for _, hit := range result.Hits.Hits {
		blogs = append(blogs, hit.Source)
	}

	return blogs, nil
}
