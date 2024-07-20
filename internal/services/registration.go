package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Registration(user models.User) (err error) {
	_, err = s.Repo.GetUserByEmail(user)

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

	err = s.Repo.CreateUser(user)
	if err != nil {
		logrus.Error("Error CreateUser in Registration", err)
		return err
	}

	return
}
