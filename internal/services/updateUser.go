package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

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
