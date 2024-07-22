package models

import "time"

type Directory struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	ParentId  *int      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
}
