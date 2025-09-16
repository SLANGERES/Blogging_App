package main

import (
	"github/SLANGERES/CQRS/Write/database"
	"github/SLANGERES/CQRS/Write/internal/handler"
	"github/SLANGERES/CQRS/Write/internal/repository"
	"github/SLANGERES/CQRS/Write/internal/router"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	// Configure database
	url := os.Getenv("DATABASE_URL")
	db, err := database.ConfigStorage(url)
	if err != nil {
		//unable to config db exits
		slog.Error("unable to connect with db", "error", err)
	}
	// Create repository
	blogRepo := repository.NewBlogRepository(db)

	// Create handler (inject repo)
	blogHandler := handler.NewBlogHandler(blogRepo)
	// Pass handler into router
	r := router.New(blogHandler)

	// Start server
	http.ListenAndServe(":8000", r)
}
