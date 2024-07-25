package services

import (
	"CloudStorage/internal/models"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (s *Service) UploadFile(userID int, directoryID int, file multipart.File, header *multipart.FileHeader) (err error) {
	var uploadDir string
	if directoryID == 0 {
		uploadDir = filepath.Join("uploads", strconv.Itoa(userID))
		directoryID, err = s.Repo.GetRootDirectoryByUserId(userID)
	} else {
		directory, err := s.GetDirectoryById(directoryID, userID)
		if err != nil {
			return err
		}
		uploadDir = filepath.Join("uploads", strconv.Itoa(userID), directory.Name)
	}

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	filePath := filepath.Join(uploadDir, header.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	fileModel := models.File{
		FileName:    header.Filename,
		UserId:      userID,
		DirectoryId: directoryID,
	}

	err = s.Repo.SaveFile(fileModel)
	if err != nil {
		return fmt.Errorf("failed to save file record in database: %w", err)
	}

	return nil
}

func (s *Service) RenameFile(id, userId int, newFileName string) (err error) {
	file, err := s.Repo.GetFileById(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in service, RenameFile - GetFileById")
		return err
	}

	var oldFilePath, newFilePath string
	if file.DirectoryId == 0 {
		oldFilePath = filepath.Join("uploads", strconv.Itoa(userId), file.FileName)
		newFilePath = filepath.Join("uploads", strconv.Itoa(userId), newFileName)
	} else {
		dir, err := s.Repo.GetDirectoryById(file.DirectoryId, userId)
		if err != nil {
			s.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("error in service, RenameFile - GetDirectoryById")
			return err
		}
		oldFilePath = filepath.Join("uploads", strconv.Itoa(userId), dir.Name, file.FileName)
		newFilePath = filepath.Join("uploads", strconv.Itoa(userId), dir.Name, newFileName)
	}

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

func (s *Service) RemoveFile(id, userId int) error {
	file, err := s.Repo.GetFileById(id, userId)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error getting file info in service, RemoveFile")
		return err
	}

	var filePath string
	fmt.Println("id = ", file)
	if file.DirectoryId == 0 {
		filePath = filepath.Join("uploads", strconv.Itoa(userId), file.FileName)
	} else {
		dir, err := s.Repo.GetDirectoryById(file.DirectoryId, userId)
		if err != nil {
			s.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("error in service, RemoveFile - GetDirectoryById")
			return err
		}
		filePath = filepath.Join("uploads", strconv.Itoa(userId), dir.Name, file.FileName)
	}

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

func (s *Service) GetFileById(id, userId int) (file models.File, err error) {
	file, err = s.Repo.GetFileById(id, userId)
	if err != nil {
		return
	}
	return file, nil
}

func (s *Service) GetFileList(userId int) (files []models.File, err error) {
	files, err = s.Repo.GetFileList(userId)
	if err != nil {
		return nil, err
	}
	return files, nil
}
