package main

import (
	"github/SLANGERES/CQRS/Read/database"
	"github/SLANGERES/CQRS/Read/internal/handler"
	"github/SLANGERES/CQRS/Read/internal/repository"
	"github/SLANGERES/CQRS/Read/internal/router"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Unable to load env", "error", err)
	}

	dbStorage := database.ConfigDatabase()

	blogRepo := repository.NewBlogRepo(dbStorage)

	blogHander := handler.NewBlogHandler(blogRepo)

	ApiRouter := router.NewRouter(blogHander)

	ApiRouter.Run(":9090")

	if dbStorage == nil {
		os.Exit(0)
	}

}
