package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"CloudStorage/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(user models.User) (accessToken string, err error) {
	userFromDb, err := s.Repo.GetUserByEmail(user)
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

	accessToken, err = utils.CreateToken(s.Config.JwtSecretKey, userFromDb.Id, user.Role.Id)
	return accessToken, err
}
