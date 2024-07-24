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
	GetUserByEmail(email string) (userFromDB models.User, err error)
	GetUserList() (users []models.User, err error)
	GetUserByID(id int) (user models.User, err error)
	UpdateUser(user models.UserUpdateDto) (err error)
	DeleteUser(userId int) (err error)
	SaveFile(file models.File) (err error)
	GetFileList(userId int) (files []models.File, err error)
	GetFileById(id, userId int) (file models.File, err error)
	RemoveFile(id, userId int) (err error)
	RenameFile(id, userId int, newFileName string) (err error)
	AdminGetUserList() (users []models.User, err error)
	CreateDirectory(directory models.Directory) (err error)
	GetDirectoryById(id, userId int) (directory models.Directory, err error)
	GetRootDirectoryByUserId(userId int) (directory int, err error)
	AdminUpdateUser(user models.UserUpdateDto) (err error)
	AdminGetUserByID(id int) (user models.User, err error)
	RenameDirectory(id, userId int, newDirName string) (err error)
	GetFilesByDirectoryId(directoryId, userId int) (files []models.File, err error)
	DeleteDirectory(id, userId int) (err error)
	AddFileAccess(fileId, userId int) (err error)
	GetFileAccessUsers(fileId int) (users []models.FileAccess, err error)
	DeleteFileAccess(fileId, userId int) (err error)
}

func NewRepository(conn *pgx.Conn, logger *logrus.Logger) RepositoryInterface {
	return &Repository{
		Conn:   conn,
		Logger: logger,
	}
}
