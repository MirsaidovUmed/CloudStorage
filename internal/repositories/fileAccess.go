package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) AddFileAccess(fileId, userId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `INSERT INTO file_access (file_id, user_id) VALUES ($1, $2)`, fileId, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, AddFileAccess")
		return err
	}
	return
}

func (repo *Repository) GetFileAccessUsers(fileId int) (users []models.FileAccess, err error) {
	rows, err := repo.Conn.Query(context.Background(), `SELECT user_id FROM file_access WHERE file_id = $1`, fileId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, ShareFileUsersList")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fileAccess models.FileAccess
		err = rows.Scan(&fileAccess.UserId)
		if err != nil {
			repo.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("error scanning row in ShareFileUsersList")
			return nil, err
		}
		users = append(users, fileAccess)
	}

	if err = rows.Err(); err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("row iteration error in ShareFileUsersList")
		return nil, err
	}

	return users, nil
}

func (repo *Repository) DeleteFileAccess(fileId, userId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `DELETE FROM file_access where file_id = $1 and user_id = $2`, fileId, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, DeleteFileAccess")
		return err
	}
	return
}
