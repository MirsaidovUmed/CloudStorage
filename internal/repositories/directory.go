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

func (repo *Repository) GetRootDirectoryByUserId(userId int) (directoryId int, err error) {
	err = repo.Conn.QueryRow(context.Background(), `
	SELECT id FROM directories WHERE user_id = $1`, userId).Scan(&directoryId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetRootDirectoryByUserId")
		return 0, err
	}
	return directoryId, nil
}

func (repo *Repository) RenameDirectory(id, userId int, newDirName string) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `UPDATE directories SET name = $1 WHERE id = $2 AND user_id = $3`, newDirName, id, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, RenameDirectory")
		return err
	}
	return nil
}

func (repo *Repository) GetFilesByDirectoryId(directoryId, userId int) (files []models.File, err error) {
	rows, err := repo.Conn.Query(context.Background(), `SELECT id, file_name, user_id, directory_id FROM files WHERE directory_id = $1 AND user_id = $2`, directoryId, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetFilesByDirectoryId")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var file models.File
		err := rows.Scan(&file.Id, &file.FileName, &file.UserId, &file.DirectoryId)
		if err != nil {
			repo.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("error scanning file in repo, GetFilesByDirectoryId")
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (repo *Repository) DeleteDirectory(id, userId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `DELETE FROM directories WHERE id = $1 AND user_id = $2`, id, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, DeleteDirectory")
		return err
	}
	return nil
}
