package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetUserByEmail(email string) (userFromDB models.UserCreateDto, err error) {
	row := repo.Conn.QueryRow(context.Background(), `SELECT id, first_name, second_name, email, password, role_id FROM users WHERE email = $1`, email)

	err = row.Scan(&userFromDB.Id, &userFromDB.FirstName, &userFromDB.SecondName, &userFromDB.Email, &userFromDB.Password, &userFromDB.Role.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"email": email,
			"err":   err,
		}).Error("error in repo, GetUserByEmail")
	}

	return
}
