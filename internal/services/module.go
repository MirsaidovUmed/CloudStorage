package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/internal/repositories"
	"CloudStorage/pkg/config"

	"github.com/sirupsen/logrus"
)

type Service struct {
	Repo   repositories.RepositoryInterface
	Config *config.Config
	Logger *logrus.Logger
}

type ServiceInterface interface {
	Registration(user models.User) (err error)
	CheckUserById(userID int) (err error)
}

func NewService(repo repositories.RepositoryInterface, config *config.Config, logger *logrus.Logger) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
		Logger: logger,
	}
}
