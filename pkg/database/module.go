package database

import (
	"CloudStorage/pkg/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewDatabase(config *config.Config) *pgx.Conn {
	con, err := pgx.Connect(context.Background(), config.PostgresUrl)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return con
}
