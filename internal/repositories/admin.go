package repositories

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) AdminGetUserList() (users []models.User, err error) {
	rows, err := repo.Conn.Query(context.Background(), `SELECT id, first_name, second_name, email, created_at FROM users`)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetUserList")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Email, &user.CreatedAt)
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

func (repo *Repository) DeleteUser(userId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, DeleteUser")
	}
	return
}

func (repo *Repository) AdminUpdateUser(user models.UserUpdateDto) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `UPDATE users SET 
			first_name = $2,
			second_name = $3,
			email = $4,
			password = $5,
			role_id = $6
		WHERE id = $1`, user.Id, user.FirstName, user.SecondName, user.Email, user.Password, user.Role.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, UpdateUser")
	}

	return
}

func (repo *Repository) AdminGetUserByID(id int) (user models.User, err error) {
	row := repo.Conn.QueryRow(context.Background(), `SELECT id, first_name, second_name, email, role_id, created_at FROM users WHERE id = $1`, id)

	err = row.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Email, &user.Role.Id, &user.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetUserByID")
	}
	return user, err
}
