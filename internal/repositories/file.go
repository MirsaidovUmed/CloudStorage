package repositories

import (
	"CloudStorage/internal/models"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) SaveFile(file models.File) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `INSERT INTO files (file_name, user_id, created_at) VALUES ($1, $2, $3)`, file.FileName, file.UserId, file.CreatedAt)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, RemoveFile")
		return err
	}
	return
}

func (repo *Repository) GetFileById(id, userId int) (file models.File, err error) {
	row := repo.Conn.QueryRow(context.Background(), `SELECT id, file_name, created_at FROM files WHERE id = $1 AND user_id = $2`, id, userId)

	err = row.Scan(&file.Id, &file.FileName, &file.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			repo.Logger.WithFields(logrus.Fields{
				"id": id,
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

func (repo *Repository) RemoveFile(id, userId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `DELETE FROM files WHERE id = $1 AND user_id = $2`, id, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, RemoveFile")
		return err
	}
	return
}

func (repo *Repository) RenameFile(id, userId int, newFileName string) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `UPDATE files SET file_name = $1 WHERE id = $2 AND user_id = $3`, newFileName, id, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, RenameFile")
		return err
	}
	return nil
}
