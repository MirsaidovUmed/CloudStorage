package repositories

import "github.com/jackc/pgx/v5"

type Repository struct {
	Conn *pgx.Conn
}

type RepositoryInterface interface {
}

func NewRepository(conn *pgx.Conn) RepositoryInterface {
	return &Repository{
		Conn: conn,
	}
}
