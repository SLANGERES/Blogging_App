package database

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
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
	slog.Info("Connected to Elasticsearch")

	return &Storage{DB: es}
}
