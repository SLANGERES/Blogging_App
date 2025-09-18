package database

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github/SLANGERES/CQRS/Read/internal/models"
	"io"
	"log/slog"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"

)

type Storage struct {
	DB *elasticsearch.Client
}

func ConfigDatabase() *Storage {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // your ES endpoint
		},
		Username: os.Getenv("ELASTIC_USERNAME"), // optional if security enabled
		Password: os.Getenv("ELASTIC_PASSWORD"),
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		slog.Error("Unable to connect to Elasticsearch", "error", err)
		return nil
	}

	// Check if index exists first
	existsRes, err := es.Indices.Exists([]string{"blogs"})
	if err != nil {
		slog.Error("Failed to check if index exists", "error", err)
		return nil
	}
	existsRes.Body.Close()

	// Only create index if it doesn't exist
	if existsRes.StatusCode == 404 {
		mapping := `
		{
		  "mappings": {
		    "properties": {
		      "id":         { "type": "keyword" },
		      "title":      { "type": "text" },
		      "description":{ "type": "text" },
		      "content":    { "type": "text" },
		      "tags":       { "type": "keyword" },
		      "category":   { "type": "keyword" },
		      "likes":      { "type": "integer" },
		      "created_at": { "type": "date" },
		      "updated_at": { "type": "date" }
		    }
		  }
		}`

		res, err := es.Indices.Create("blogs", es.Indices.Create.WithBody(bytes.NewReader([]byte(mapping))))
		if err != nil {
			slog.Error("Failed to create index", "error", err)
			return nil
		}
		defer res.Body.Close()

		if res.IsError() {
			slog.Error("Elasticsearch index creation failed", "status", res.String())
			return nil
		}
		slog.Info("Created Elasticsearch index 'blogs'")
	} else if existsRes.StatusCode == 200 {
		slog.Info("Elasticsearch index 'blogs' already exists")
	} else {
		slog.Error("Unexpected response when checking index existence", "status", existsRes.String())
		return nil
	}

	slog.Info("Connected to Elasticsearch")

	return &Storage{DB: es}
}

// InsertInDb indexes a blog document in Elasticsearch
func (s *Storage) InsertInDb(blog models.Blog) error {
	// Convert to JSON
	data, err := json.Marshal(blog)
	if err != nil {
		return fmt.Errorf("marshal blog: %w", err)
	}

	// Index request
	req := esapi.IndexRequest{
		Index:      "blogs",
		DocumentID: blog.ID,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	// Execute
	res, err := req.Do(context.Background(), s.DB)
	if err != nil {
		return fmt.Errorf("index blog: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("elasticsearch error [%s]: %s", res.Status(), body)
	}

	slog.Info("Blog inserted", "id", blog.ID, "title", blog.Title)
	return nil
}
