package services

import "CloudStorage/internal/models"

func (s *Service) GetUserList() (users []models.User, err error) {
	users, err = s.Repo.GetUserList()
	if err != nil {
		return nil, err
	}
	return users, nil
}
