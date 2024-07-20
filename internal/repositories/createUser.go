package repositories

import (
	"CloudStorage/internal/models"
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateUser(user models.UserCreateDto) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `INSERT INTO users (first_name, second_name, email, password, role_id) VALUES ($1, $2, $3, $4, $5)`, user.FirstName, user.SecondName, user.Email, user.Password, user.Role.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user": user,
			"err":  err,
		}).Error("error in repo, CreateUser")
	}

	return
}
