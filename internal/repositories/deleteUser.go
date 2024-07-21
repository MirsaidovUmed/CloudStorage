package repositories

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) DeleteUser(userId int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, userId)
	if err != nil {
		if err != nil {
			repo.Logger.WithFields(logrus.Fields{
				"user": userId,
				"err":  err,
			}).Error("error in repo, DeleteUser")
		}
	}
	return
}
