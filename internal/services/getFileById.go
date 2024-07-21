package services

import "CloudStorage/internal/models"

func (s *Service) GetFileById(id, userId int) (file models.File, err error) {
	file, err = s.Repo.GetFileById(id, userId)
	if err != nil {
		return
	}
	return file, nil
}
