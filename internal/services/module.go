package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/internal/repositories"
	"CloudStorage/pkg/config"
	"mime/multipart"

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
	GetUserList() (users []models.User, err error)
	GetUserByID(id int) (user models.User, err error)
	UpdateUser(user models.UserUpdateDto) (err error)
	DeleteUser(userId int) (err error)
	UploadFile(userID int, directoryID int, file multipart.File, header *multipart.FileHeader) (err error)
	GetFileList(userId int) (files []models.File, err error)
	GetFileById(id, userId int) (file models.File, err error)
	RemoveFile(id, userId int) error
	RenameFile(id, userId int, newFileName string) (err error)
	AdminGetUserList() (users []models.User, err error)
	CreateDirectory(directory models.Directory) (err error)
	AdminUpdateUser(user models.UserUpdateDto) (err error)
	AdminGetUserByID(id int) (user models.User, err error)
	RenameDirectory(id, userId int, newDirName string) (err error)
	GetDirectoryById(id, userId int) (models.Directory, error)
	DeleteDirectory(id, userId int) (err error)
	ShareFile(grantorId, fileId, granteeId int) error
	GetFileAccessUsers(fileId int) ([]models.FileAccess, error)
	DeleteFileAccess(grantorId, fileId, granteeId int) (err error)
}

func NewService(repo repositories.RepositoryInterface, config *config.Config, logger *logrus.Logger) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
		Logger: logger,
	}
}
