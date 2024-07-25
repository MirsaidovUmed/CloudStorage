package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (s *Service) ShareFile(grantorId, fileId, granteeId int) error {
	file, err := s.Repo.GetFileById(fileId, grantorId)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	if file.UserId != grantorId {
		return fmt.Errorf("user is not the owner of the file")
	}

	granteeFromDb, err := s.Repo.GetUserByID(granteeId)
	if err != nil {
		return fmt.Errorf("target user not found: %w", err)
	}

	err = s.Repo.AddFileAccess(grantorId, fileId, granteeFromDb.Id)
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

func (s *Service) DeleteFileAccess(grantorId, fileId, granteeId int) (err error) {
	_, err = s.Repo.GetUserByID(granteeId)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return err
	}

	_, err = s.Repo.GetFileById(fileId, grantorId)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return err
	}

	err = s.Repo.DeleteFileAccess(grantorId, fileId, granteeId)
	if err != nil {
		return err
	}
	return
}
