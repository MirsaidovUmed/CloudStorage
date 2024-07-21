package services

import (
	"CloudStorage/internal/models"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func (s *Service) UploadFile(userID int, file multipart.File, header *multipart.FileHeader) (err error) {
	filePath := filepath.Join("uploads", header.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	fileModel := models.File{
		FileName:  header.Filename,
		UserId:    userID,
		CreatedAt: time.Now(),
	}

	err = s.Repo.SaveFile(fileModel)
	if err != nil {
		return err
	}

	return nil
}
