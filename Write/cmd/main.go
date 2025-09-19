package main

import (
	"github/SLANGERES/CQRS/Write/database"
	"github/SLANGERES/CQRS/Write/internal/broker"
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
	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		slog.Warn("unable to get db url")
	}
	db, err := database.ConfigStorage(db_url)
	if err != nil {
		//unable to config db exits
		slog.Error("unable to connect with db", "error", err)
	}
	slog.Info("DB connected sucessfully")

	mq_url := os.Getenv("RABBITMQ_URL")
	if mq_url == "" {
		slog.Warn("unable to get db url")
	}
	// 2️⃣ Create a new connection
	mq, err := broker.NewConnection(mq_url)
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ:", "error", err)
	}
	slog.Info("Message Queue connected sucessfully")
	defer mq.Close()

	// Create repository
	blogRepo := repository.NewBlogRepository(db)

	// Create handler (inject repo)
	blogHandler := handler.NewBlogHandler(blogRepo, mq)

	// Pass handler into router
	r := router.New(blogHandler)
	slog.Info("Server connected sucessfully 8000")

	// Start server
	http.ListenAndServe(":8000", r)
}
