package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetUserByEmail(user models.User) (userFromDB models.User, err error) {
	row := repo.Conn.QueryRow(context.Background(), `SELECT * FROM users WHERE email = $1`, user.Email)

	err = row.Scan(&userFromDB.Id, &userFromDB.FirstName, &userFromDB.SecondName, &userFromDB.Email, &userFromDB.Password, &userFromDB.Role.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user": user,
			"err":  err,
		}).Error("error in repo, GetUserByEmail")
	}

	return
}
