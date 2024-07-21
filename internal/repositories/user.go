package repositories

import (
	"CloudStorage/internal/models"
	"CloudStorage/pkg/errors"
	"context"

	"github.com/jackc/pgx/v5"
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

func (repo *Repository) DeleteUser(userId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, userId)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, DeleteUser")
	}
	return
}

func (repo *Repository) GetUserByEmail(email string) (userFromDB models.UserCreateDto, err error) {
	row := repo.Conn.QueryRow(context.Background(), `SELECT id, first_name, second_name, email, password, role_id FROM users WHERE email = $1`, email)

	err = row.Scan(&userFromDB.Id, &userFromDB.FirstName, &userFromDB.SecondName, &userFromDB.Email, &userFromDB.Password, &userFromDB.Role.Id)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetUserByEmail")
	}

	return
}

func (repo *Repository) GetUserByID(id int) (user models.UserCreateDto, err error) {
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

func (repo *Repository) UpdateUser(user models.UserUpdateDto) (err error) {
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

func (repo *Repository) AdminGetUserList() (users []models.UserCreateDto, err error) {
	rows, err := repo.Conn.Query(context.Background(), `SELECT id, first_name, second_name, email, created_at FROM users`)
	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetUserList")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserCreateDto
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
