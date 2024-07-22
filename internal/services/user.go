package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"CloudStorage/pkg/utils"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Registration(user models.UserCreateDto) (err error) {
	_, err = s.Repo.GetUserByEmail(user.Email)
	if err != errors.ErrDataNotFound {
		if err == nil {
			return errors.ErrAlreadyHasUser
		}
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("Error in Hashing Password", err)
		return err
	}
	user.Password = string(hashedPassword)

	if user.Role.Id == 0 {
		user.Role.Id = 2
	}

	err = s.Repo.CreateUser(user)
	if err != nil {
		logrus.Error("Error CreateUser in Registration", err)
		return err
	}

	createdUser, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		logrus.Error("Error retrieving created user", err)
		return err
	}

	userDir := filepath.Join("uploads", strconv.Itoa(createdUser.Id))
	err = os.MkdirAll(userDir, os.ModePerm)
	if err != nil {
		logrus.Error("Error creating user directory", err)
		return err
	}

	rootDirectory := models.Directory{
		Name:      strconv.Itoa(createdUser.Id),
		UserId:    createdUser.Id,
		ParentId:  nil,
		CreatedAt: time.Now(),
	}

	err = s.Repo.CreateDirectory(rootDirectory)
	if err != nil {
		logrus.Error("Error creating root directory in database", err)
		os.RemoveAll(userDir)
		return err
	}

	return nil
}
func (s *Service) Login(user models.UserCreateDto) (accessToken string, err error) {
	userFromDb, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.ErrWrongPassword
		return "", err
	} else if err != nil {
		return "", err
	}

	accessToken, err = utils.CreateToken(s.Config.JwtSecretKey, userFromDb.Id, userFromDb.Role.Id)
	return accessToken, err
}

func (s *Service) GetUserByID(id int) (user models.UserCreateDto, err error) {
	user, err = s.Repo.GetUserByID(id)
	if err != nil {
		return
	}
	return user, nil
}

func (s *Service) GetUserList() (users []models.UserCreateDto, err error) {
	users, err = s.Repo.GetUserList()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) UpdateUser(user models.UserUpdateDto) (err error) {
	existingUser, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	existingUserUpdate := existingUser.ToUserUpdateDto()

	if user.FirstName != models.UnsetValue {
		existingUserUpdate.FirstName = user.FirstName
	}

	if user.SecondName != models.UnsetValue {
		existingUserUpdate.SecondName = user.SecondName
	}

	if user.Password != models.UnsetValue {
		err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
		if err == nil {
			return errors.ErrWrongPassword
		} else if err != bcrypt.ErrMismatchedHashAndPassword {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingUserUpdate.Password = string(hashedPassword)
	}

	err = s.Repo.UpdateUser(existingUserUpdate)
	return err
}
