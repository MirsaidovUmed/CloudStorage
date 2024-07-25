package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) AddFileAccess(grantorId, fileId, granteeId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `INSERT INTO file_access (file_id, grantor_id, grantee_id) VALUES ($1, $2, $3)`, fileId, grantorId, granteeId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, AddFileAccess")
		return err
	}
	return
}

func (repo *Repository) GetFileAccessUsers(fileId int) (users []models.FileAccess, err error) {
	rows, err := repo.Conn.Query(context.Background(), `SELECT grantee_id FROM file_access WHERE file_id = $1`, fileId)
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

func (repo *Repository) DeleteFileAccess(grantorId, fileId, granteeId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `DELETE FROM file_access where file_id = $1 AND grantor_id = $2 and grantee_id = $3`, grantorId, fileId, granteeId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, DeleteFileAccess")
		return err
	}
	return
}
