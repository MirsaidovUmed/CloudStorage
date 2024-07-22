package models

import "time"

type File struct {
	Id          int       `json:"id"`
	FileName    string    `json:"file_name"`
	UserId      int       `json:"user_id"`
	DirectoryId int       `json:"directory_id"`
	CreatedAt   time.Time `json:"created_at"`
}
