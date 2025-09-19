package main

import (
	"github/SLANGERES/CQRS/Read/database"
	"github/SLANGERES/CQRS/Read/internal/broker"
	"github/SLANGERES/CQRS/Read/internal/handler"
	"github/SLANGERES/CQRS/Read/internal/models"
	"github/SLANGERES/CQRS/Read/internal/repository"
	"github/SLANGERES/CQRS/Read/internal/router"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Unable to load env", "error", err)
	}

	dbStorage := database.ConfigDatabase()
	if dbStorage == nil {
		slog.Error("Failed to connect to database")
		os.Exit(1)
	}

	mq_url := os.Getenv("RABBITMQ_URL")
	if mq_url == "" {
		slog.Warn("Unable to get the rabbit mq url from env")
	}

	mqConnection, err := broker.NewConsumer(mq_url, dbStorage)
	if err != nil {
		slog.Error("unable to connect with message queue server", "error", err)
		os.Exit(0)
	}

	// Start the consumer in a separate goroutine
	go func() {
		// Simple handler function that processes each blog
		handleMessage := func(blog models.Blog) error {
			slog.Info("Processing blog message", "blog_id", blog.ID, "title", blog.Title)
			return nil // The consumer already handles DB insertion
		}

		// Start consuming messages from the queue
		if err := mqConnection.Consume("blog-sync-mq", handleMessage); err != nil {
			slog.Error("Failed to start consumer", "error", err)
			os.Exit(1)
		}

		// Keep the consumer running
		select {}
	}()

	// Give the consumer a moment to start
	time.Sleep(2 * time.Second)
	slog.Info("RabbitMQ consumer started successfully")

	blogRepo := repository.NewBlogRepo(dbStorage)
	blogHander := handler.NewBlogHandler(blogRepo, mqConnection)
	ApiRouter := router.NewRouter(blogHander)

	ApiRouter.Run(":9090")



	ApiRouter.Run(":9090")

}
