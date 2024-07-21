package services

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func (s *Service) RemoveFile(id, userId int) error {
	file, err := s.Repo.GetFileById(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error getting file info in service, RemoveFile")
		return err
	}

	filePath := filepath.Join("uploads", file.FileName)
	err = os.Remove(filePath)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"filePath": filePath,
			"err":      err,
		}).Error("error removing file from filesystem in service, RemoveFile")
		return err
	}

	err = s.Repo.RemoveFile(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error removing file record from database in service, RemoveFile")
		return err
	}

	return nil
}
