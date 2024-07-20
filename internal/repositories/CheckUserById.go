package repositories

import (
	"CloudStorage/pkg/errors"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) CheckUserById(userID int) (err error) {
	var idFromDB int

	row := repo.Conn.QueryRow(context.Background(), `
			SELECT id FROM users WHERE id = $1`, userID)

	err = row.Scan(&idFromDB)

	if err != nil {
		logrus.Error("Error in CheckUserById Repo ", err)
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
