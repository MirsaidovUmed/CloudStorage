package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetUserList() (users []models.UserCreateDto, err error) {
	rows, err := repo.Conn.Query(context.Background(), `SELECT id, first_name, second_name, email FROM users`)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetUserList")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserCreateDto
		err := rows.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Email)
		if err != nil {
			repo.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("error scanning row in GetUserList")
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error iterating rows in GetUserList")
		return nil, err
	}

	return users, nil
}
