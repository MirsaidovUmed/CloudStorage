package repositories

import (
	"CloudStorage/internal/models"
	"context"
)

func (repo *Repository) SaveFile(file models.File) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO files (file_name, user_id, created_at)
		VALUES ($1, $2, $3)`, file.FileName, file.UserId, file.CreatedAt)
	return
}
