package services

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/constants"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Service) CreateDirectory(directory models.Directory) (err error) {
	var parentDir string
	if directory.ParentId == nil {
		parentDir = filepath.Join(constants.Uploads, strconv.Itoa(directory.UserId))
	} else {
		parentDirectory, err := s.GetDirectoryById(*directory.ParentId, directory.UserId)
		if err != nil {
			return err
		}
		parentDir = filepath.Join(constants.Uploads, strconv.Itoa(directory.UserId), parentDirectory.Name)
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

func (s *Service) RenameDirectory(id, userId int, newDirName string) (err error) {
	dir, err := s.Repo.GetDirectoryById(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in service, RenameDirectory - GetDirectoryById")
		return err
	}

	oldDirPath := filepath.Join(constants.Uploads, strconv.Itoa(userId), dir.Name)
	newDirPath := filepath.Join(constants.Uploads, strconv.Itoa(userId), newDirName)

	err = os.Rename(oldDirPath, newDirPath)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in service, RenameDirectory - os.Rename")
		return err
	}

	err = s.Repo.RenameDirectory(id, userId, newDirName)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in service, RenameDirectory - RenameDirectory")
		return err
	}

	return nil
}

func (s *Service) DeleteDirectory(id, userId int) (err error) {
	dir, err := s.Repo.GetDirectoryById(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error getting dir info in service, DeleteDirectory")
		return err
	}

	files, err := s.Repo.GetFilesByDirectoryId(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error getting files in directory in service, DeleteDirectory")
		return err
	}

	for _, file := range files {
		err = s.RemoveFile(file.Id, userId)
		if err != nil {
			return err
		}
	}

	dirPath := filepath.Join(constants.Uploads, strconv.Itoa(userId), dir.Name)
	err = os.RemoveAll(dirPath)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"filePath": dirPath,
			"err":      err,
		}).Error("error removing directory from filesystem in service, DeleteDirectory")
		return err
	}

	err = s.Repo.DeleteDirectory(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error removing directory from database in service, DeleteDirectory")
		return err
	}

	return nil
}
