package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
)

func (s *Service) AdminGetUserList() (users []models.User, err error) {
	users, err = s.Repo.AdminGetUserList()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) DeleteUser(userId int) (err error) {
	_, err = s.Repo.GetUserByID(userId)
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
