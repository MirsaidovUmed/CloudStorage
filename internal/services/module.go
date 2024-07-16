package services

import (
	"CloudStorage/internal/repositories"
	"CloudStorage/pkg/config"
)

type Service struct {
	Repo   repositories.RepositoryInterface
	Config *config.Config
}

type ServiceInterface interface {
}

func NewService(repo repositories.RepositoryInterface, config *config.Config) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
	}
}
