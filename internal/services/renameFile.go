package services

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func (s *Service) RenameFile(id, userId int, newFileName string) (err error) {
	file, err := s.Repo.GetFileById(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in service, RenameFile - GetFileById")
		return err
	}

	oldFilePath := filepath.Join("uploads", file.FileName)
	newFilePath := filepath.Join("uploads", newFileName)

	err = os.Rename(oldFilePath, newFilePath)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in service, RenameFile - os.Rename")
		return err
	}

	err = s.Repo.RenameFile(id, userId, newFileName)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in service, RenameFile - RenameFile")
		return err
	}

	return nil
}
