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
	Registration(user models.UserCreateDto) (err error)
	Login(user models.UserCreateDto) (accessToken string, err error)
	GetUserList() (users []models.UserCreateDto, err error)
	GetUserByID(id int) (user models.UserCreateDto, err error)
	UpdateUser(user models.UserUpdateDto) (err error)
}

func NewService(repo repositories.RepositoryInterface, config *config.Config, logger *logrus.Logger) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
		Logger: logger,
	}
}
