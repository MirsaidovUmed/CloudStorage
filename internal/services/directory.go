package services

import (
	"CloudStorage/internal/models"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (s *Service) CreateDirectory(directory models.Directory) (err error) {
	var parentDir string
	if directory.ParentId == nil {
		parentDir = filepath.Join("uploads", strconv.Itoa(directory.UserId))
	} else {
		parentDirectory, err := s.GetDirectoryById(*directory.ParentId, directory.UserId)
		if err != nil {
			return err
		}
		parentDir = filepath.Join("uploads", strconv.Itoa(directory.UserId), parentDirectory.Name)
	}

	newDirPath := filepath.Join(parentDir, directory.Name)

	err = os.MkdirAll(newDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	directory.CreatedAt = time.Now()

	err = s.Repo.CreateDirectory(directory)
	if err != nil {
		os.RemoveAll(newDirPath)
		return err
	}

	return nil
}

func (s *Service) GetDirectoryById(id, userId int) (models.Directory, error) {
	return s.Repo.GetDirectoryById(id, userId)
}
