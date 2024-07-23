package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AdminGetUserList() (users []models.User, err error) {
	users, err = s.Repo.AdminGetUserList()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) DeleteUser(userId int) (err error) {
	_, err = s.Repo.AdminGetUserByID(userId)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return err
	}
	err = s.Repo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return
}

func (s *Service) AdminUpdateUser(user models.UserUpdateDto) (err error) {
	existingUser, err := s.Repo.AdminGetUserByID(user.Id)
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

	err = s.Repo.AdminUpdateUser(existingUserUpdate)
	return err
}

func (s *Service) AdminGetUserByID(id int) (user models.User, err error) {
	user, err = s.Repo.AdminGetUserByID(id)
	if err != nil {
		return
	}
	return user, nil
}
