package services

import "CloudStorage/internal/models"

func (s *Service) GetFileList(userId int) (files []models.File, err error) {
	files, err = s.Repo.GetFileList(userId)
	if err != nil {
		return nil, err
	}
	return files, nil
}
