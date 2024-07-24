package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (s *Service) ShareFile(userId, fileId, targetUserId int) error {
	file, err := s.Repo.GetFileById(fileId, userId)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	if file.UserId != userId {
		return fmt.Errorf("user is not the owner of the file")
	}

	user, err := s.Repo.GetUserByID(targetUserId)
	if err != nil {
		return fmt.Errorf("target user not found: %w", err)
	}

	err = s.Repo.AddFileAccess(fileId, user.Id)
	if err != nil {
		return fmt.Errorf("failed to share file: %w", err)
	}

	return nil
}

func (s *Service) GetFileAccessUsers(fileId int) ([]models.FileAccess, error) {
	users, err := s.Repo.GetFileAccessUsers(fileId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error getting file access users in service, GetFileAccessUsers")
		return nil, err
	}
	return users, nil
}

func (s *Service) DeleteFileAccess(fileId, userId int) (err error) {
	_, err = s.Repo.AdminGetUserByID(userId)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return err
	}

	_, err = s.Repo.GetFileById(fileId, userId)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return err
	}

	err = s.Repo.DeleteFileAccess(fileId, userId)
	if err != nil {
		return err
	}
	return
}
