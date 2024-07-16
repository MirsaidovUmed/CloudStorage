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
)

func main() {
	conf := config.NewConfig()

	db := database.NewDatabase(conf)
	repo := repositories.NewRepository(db)
	svc := services.NewService(repo, conf)
	handlers := handlers.NewHandler(svc)

	middle := middleware.NewMiddleware(conf, svc)

	router := router.InitRouter(handlers, middle)

	server := http.NewServer(conf.ServerAddress, conf.ServerPort, router)

	server.Run()
}
