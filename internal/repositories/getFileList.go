package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetFileList(userId int) (files []models.File, err error) {
	rows, err := repo.Conn.Query(context.Background(), `SELECT id, file_name, created_at FROM files WHERE user_id = $1`, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetFileList")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.Id, &file.FileName, &file.CreatedAt)
		if err != nil {
			repo.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("error scanning row in GetFileList")
			return nil, err
		}
		files = append(files, file)
	}

	if err = rows.Err(); err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error iterating rows in GetFileList")
		return nil, err
	}

	return files, err
}
