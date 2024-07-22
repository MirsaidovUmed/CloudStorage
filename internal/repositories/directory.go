package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateDirectory(directory models.Directory) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO directories (name, user_id, parent_id)
		VALUES ($1, $2, $3)`, directory.Name, directory.UserId, directory.ParentId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, CreateDirectory")
		return err
	}
	return nil
}

func (repo *Repository) GetDirectoryById(id, userId int) (directory models.Directory, err error) {
	err = repo.Conn.QueryRow(context.Background(), `
		SELECT id, name, user_id, parent_id, created_at
		FROM directories WHERE id = $1 AND user_id = $2`, id, userId).Scan(&directory.Id, &directory.Name, &directory.UserId, &directory.ParentId, &directory.CreatedAt)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetDirectoryById")
		return directory, err
	}
	return directory, nil
}
