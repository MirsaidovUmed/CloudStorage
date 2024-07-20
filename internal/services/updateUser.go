package services

import (
	"CloudStorage/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) UpdateUser(user models.UserUpdateDto) (err error) {
	existingUser, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	existingUserUpdate := existingUser.ToUserUpdateDto()

	if user.FirstName != models.UnsetValue {
		existingUser.FirstName = user.FirstName
	}

	if user.SecondName != models.UnsetValue {
		existingUser.SecondName = user.SecondName
	}

	if user.Password != models.UnsetValue {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingUser.Password = string(hashedPassword)
	}

	err = s.Repo.UpdateUser(existingUserUpdate)
	return err
}
