package services

import "CloudStorage/internal/models"

func (s *Service) AdminGetUserList() (users []models.UserCreateDto, err error) {
	users, err = s.Repo.AdminGetUserList()
	if err != nil {
		return nil, err
	}
	return users, nil
}
