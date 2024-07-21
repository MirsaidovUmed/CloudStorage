package repositories

import (
	"CloudStorage/internal/models"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetFileById(id, userId int) (file models.File, err error) {
	row := repo.Conn.QueryRow(context.Background(), `SELECT id, file_name, created_at FROM files WHERE id = $1 AND user_id = $2`, id, userId)

	err = row.Scan(&file.Id, &file.FileName, &file.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			repo.Logger.WithFields(logrus.Fields{
				"id":     id,
				"userId": userId,
			}).Info("File not found")
			err = nil
		} else {
			repo.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("error scanning row in GetFileById")
		}
		return
	}

	return file, err
}
