package repositories

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetUserByID(id int) (user models.UserCreateDto, err error) {
	row := repo.Conn.QueryRow(context.Background(), `SELECT id, first_name, second_name, email, password, role_id, created_at FROM users WHERE id = $1`, id)

	err = row.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Email, &user.Password, &user.Role.Id, &user.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
		repo.Logger.WithFields(logrus.Fields{
			"id":  id,
			"err": err,
		}).Error("error in repo, GetUserByID")
	}
	return user, err
}
