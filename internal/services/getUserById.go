package services

import (
	"CloudStorage/internal/models"
)

func (s *Service) GetUserByID(id int) (user models.UserCreateDto, err error) {
	user, err = s.Repo.GetUserByID(id)
	if err != nil {
		return
	}
	return user, nil
}
