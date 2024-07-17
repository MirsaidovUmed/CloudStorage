package repositories

import (
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Conn   *pgx.Conn
	Logger *logrus.Logger
}

type RepositoryInterface interface {
}

func NewRepository(conn *pgx.Conn, logger *logrus.Logger) RepositoryInterface {
	return &Repository{
		Conn:   conn,
		Logger: logger,
	}
}
