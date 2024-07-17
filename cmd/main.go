package main

import (
	"CloudStorage/internal/repositories"
	"CloudStorage/internal/services"
	"CloudStorage/internal/transport/http/handlers"
	"CloudStorage/internal/transport/http/middleware"
	"CloudStorage/internal/transport/http/router"
	"CloudStorage/pkg/config"
	"CloudStorage/pkg/database"
	"CloudStorage/pkg/http"
	"CloudStorage/pkg/logger"
)

func main() {
	conf := config.NewConfig()

	logger := logger.NewLogger(conf)

	db := database.NewDatabase(conf, logger)
	repo := repositories.NewRepository(db, logger)
	svc := services.NewService(repo, conf, logger)
	handlers := handlers.NewHandler(svc, logger)

	middle := middleware.NewMiddleware(conf, svc)

	router := router.InitRouter(handlers, middle)

	server := http.NewServer(conf.ServerAddress, conf.ServerPort, router)

	server.Run()
}
