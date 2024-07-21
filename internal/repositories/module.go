package repositories

import (
	"CloudStorage/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Conn   *pgx.Conn
	Logger *logrus.Logger
}

type RepositoryInterface interface {
	CreateUser(user models.UserCreateDto) (err error)
	GetUserByEmail(email string) (userFromDB models.UserCreateDto, err error)
	GetUserList() (users []models.UserCreateDto, err error)
	GetUserByID(id int) (user models.UserCreateDto, err error)
	UpdateUser(user models.UserUpdateDto) (err error)
	DeleteUser(userId int) (err error)
	SaveFile(file models.File) (err error)
	GetFileList(userId int) (files []models.File, err error)
}

func NewRepository(conn *pgx.Conn, logger *logrus.Logger) RepositoryInterface {
	return &Repository{
		Conn:   conn,
		Logger: logger,
	}
}
